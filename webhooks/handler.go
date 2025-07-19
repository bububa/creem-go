package webhooks

import (
	"bytes"
	"context"
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
	w := new(bytes.Buffer)
	tee := io.TeeReader(req.Body, w)
	signature := []byte(req.Header.Get("creem-signature"))
	if ok, err := verifySignature(tee, h.secret, signature); err != nil {
		return errors.Join(creem.ErrInvalidSignature, err)
	} else if !ok {
		return creem.ErrInvalidSignature
	}
	var ev Event
	if err := json.NewDecoder(w).Decode(&ev); err != nil {
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
