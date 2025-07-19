package products

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// List list Product
func List(ctx context.Context, clt *creem.Client, req *ListRequest, ret *ListResult) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrListProduct, err)
	}
	return nil
}
