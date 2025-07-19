package products

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Create create Product
func Create(ctx context.Context, clt *creem.Client, req *CreateRequest, ret *Product) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrCreateProduct, err)
	}
	return nil
}
