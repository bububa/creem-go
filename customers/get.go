package customers

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// Get get Customer
func Get(ctx context.Context, clt *creem.Client, req *GetRequest, ret *Customer) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrGetCustomer, err)
	}
	return nil
}
