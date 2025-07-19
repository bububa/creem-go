package discounts

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Delete Discount Code
func Delete(ctx context.Context, clt *creem.Client, req *DeleteRequest, ret *Discount) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrDeleteDiscount, err)
	}
	return nil
}
