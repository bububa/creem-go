package checkouts

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Create create Checkout Session
func Create(ctx context.Context, clt *creem.Client, req *CreateRequest, ret *Session) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrCreateCheckoutSession, err)
	}
	return nil
}
