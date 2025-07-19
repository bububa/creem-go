package subscriptions

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Get Subscription
func Get(ctx context.Context, clt *creem.Client, req *GetRequest, ret *Subscription) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrGetSubscription, err)
	}
	return nil
}
