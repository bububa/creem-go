package licenses

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Activate License Key
func Activate(ctx context.Context, clt *creem.Client, req *ActivateRequest, ret *License) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrActivateLicense, err)
	}
	return nil
}
