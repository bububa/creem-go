package products

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Get get Product
func Get(ctx context.Context, clt *creem.Client, productID string, ret *Product) error {
	req := GetRequest{
		ID: productID,
	}
	if err := clt.Do(ctx, &req, ret); err != nil {
		return errors.Join(creem.ErrGetProduct, err)
	}
	return nil
}
