package product

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bububa/creem-go"
)

// List list Product
func List(ctx context.Context, clt *creem.Client, req *creem.PageInfo, ret *ListResult) error {
	values := make(url.Values)
	values.Set("page_number", strconv.FormatInt(req.PageNumber, 10))
	values.Set("page_size", strconv.FormatInt(req.PageSize, 10))
	gw := fmt.Sprintf("%s/v1/products?%s", clt.BaseURL(), values.Encode())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gw, nil)
	if err != nil {
		return errors.Join(creem.ErrListProduct, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrListProduct, err)
	}
	return nil
}
