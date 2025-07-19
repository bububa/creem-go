package products

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bububa/creem-go"
)

// Get get Product
func Get(ctx context.Context, clt *creem.Client, productID string, ret *Product) error {
	gw := fmt.Sprintf("%s/v1/products?product_id=%s", clt.BaseURL(), url.QueryEscape(productID))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gw, nil)
	if err != nil {
		return errors.Join(creem.ErrGetProduct, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrGetProduct, err)
	}
	return nil
}
