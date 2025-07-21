package webhooks

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/bububa/creem-go"
)

type EventHandler func(context.Context, *Event) error

type Handler struct {
	secret  []byte
	handler EventHandler
}

func NewHandler(secret []byte, handler EventHandler) *Handler {
	return &Handler{secret: secret, handler: handler}
}

func (h *Handler) Handle(req *http.Request) error {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, req.Body); err != nil {
		return errors.Join(creem.ErrInvalidSignature, err)
	}
	sigHex := req.Header.Get("creem-signature")
	signature, err := hex.DecodeString(sigHex)
	if err != nil {
		return errors.Join(creem.ErrInvalidSignature, err)
	}
	if ok, err := verifySignature(bytes.NewReader(buf.Bytes()), h.secret, signature); err != nil {
		return errors.Join(creem.ErrInvalidSignature, err)
	} else if !ok {
		return creem.ErrInvalidSignature
	}
	var ev Event
	if err := json.NewDecoder(buf).Decode(&ev); err != nil {
		return err
	}
	if h.handler != nil {
		return h.handler(req.Context(), &ev)
	}
	return nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.Handle(r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
