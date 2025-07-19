package checkouts

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Get get Checkout Session
func Get(ctx context.Context, clt *creem.Client, checkoutID string, ret *Session) error {
	req := GetRequest{
		ID: checkoutID,
	}
	if err := clt.Do(ctx, &req, ret); err != nil {
		return errors.Join(creem.ErrGetCheckoutSession, err)
	}
	return nil
}
