package subscriptions

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Update Subscription
func Update(ctx context.Context, clt *creem.Client, req *UpdateRequest, ret *Subscription) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrUpdateSubscription, err)
	}
	return nil
}
