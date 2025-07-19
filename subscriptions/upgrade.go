package subscriptions

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Upgrade Subscription
func Upgrade(ctx context.Context, clt *creem.Client, req *UpgradeRequest, ret *Subscription) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrUpgradeSubscription, err)
	}
	return nil
}
