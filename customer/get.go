package customer

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bububa/creem-go"
)

// Get get Customer
func Get(ctx context.Context, clt *creem.Client, req *GetRequest, ret *Customer) error {
	values := make(url.Values)
	if req.CustomerID != "" {
		values.Set("customer_id", req.CustomerID)
	}
	if req.Email != "" {
		values.Set("email", req.Email)
	}
	gw := fmt.Sprintf("%s/v1/customers?%s", clt.BaseURL(), values.Encode())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gw, nil)
	if err != nil {
		return errors.Join(creem.ErrGetCustomer, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrGetCustomer, err)
	}
	return nil
}
