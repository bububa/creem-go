package subscriptions

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Cancel Subscription
func Cancel(ctx context.Context, clt *creem.Client, req *CancelRequest, ret *Subscription) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrCancelSubscription, err)
	}
	return nil
}
