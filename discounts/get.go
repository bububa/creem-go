package discounts

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Get Discount Code
func Get(ctx context.Context, clt *creem.Client, req *GetRequest, ret *Discount) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrGetDiscount, err)
	}
	return nil
}
