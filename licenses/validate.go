package licenses

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Validate License Key
func Validate(ctx context.Context, clt *creem.Client, req *ValidateRequest, ret *License) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrValidateLicense, err)
	}
	return nil
}
