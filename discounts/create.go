package discounts

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Create create Discount Code
func Create(ctx context.Context, clt *creem.Client, req *CreateRequest, ret *Discount) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrCreateDiscount, err)
	}
	return nil
}
