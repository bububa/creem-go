package licenses

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Deactivate License Key
func Deactivate(ctx context.Context, clt *creem.Client, req *DeactivateRequest, ret *License) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrDeactivateLicense, err)
	}
	return nil
}
