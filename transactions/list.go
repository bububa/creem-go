package transactions

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bububa/creem-go"
)

// List list Transactions
func List(ctx context.Context, clt *creem.Client, req *ListRequest, ret *ListResult) error {
	values := make(url.Values)
	if req.CustomerID != "" {
		values.Set("customer_id", req.CustomerID)
	}
	if req.OrderID != "" {
		values.Set("order_id", req.OrderID)
	}
	if req.ProductID != "" {
		values.Set("product_id", req.ProductID)
	}
	values.Set("page_number", strconv.FormatInt(req.PageNumber, 10))
	values.Set("page_size", strconv.FormatInt(req.PageSize, 10))
	gw := fmt.Sprintf("%s/v1/transactions/search?%s", clt.BaseURL(), values.Encode())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gw, nil)
	if err != nil {
		return errors.Join(creem.ErrListTransaction, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrListTransaction, err)
	}
	return nil
}
