package checkout

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bububa/creem-go"
)

// Get get Checkout Session
func Get(ctx context.Context, clt *creem.Client, checkoutID string, ret *Session) error {
	gw := fmt.Sprintf("%s/v1/checkouts?checkout_id=%s", clt.BaseURL(), url.QueryEscape(checkoutID))
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gw, nil)
	if err != nil {
		return errors.Join(creem.ErrGetCheckoutSession, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrGetCheckoutSession, err)
	}
	return nil
}
