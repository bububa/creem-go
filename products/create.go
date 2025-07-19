package products

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bububa/creem-go"
)

// Create create Product
func Create(ctx context.Context, clt *creem.Client, req *CreateRequest, ret *Product) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return errors.Join(creem.ErrCreateCheckoutSession, err)
	}
	gw := fmt.Sprintf("%s/v1/products", clt.BaseURL())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, gw, &buf)
	if err != nil {
		return errors.Join(creem.ErrCreateProduct, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrCreateProduct, err)
	}
	return nil
}
