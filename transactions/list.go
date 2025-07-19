package transactions

import (
	"context"
	"errors"

	"github.com/bububa/creem-go"
)

// List list Transactions
func List(ctx context.Context, clt *creem.Client, req *ListRequest, ret *ListResult) error {
	if err := clt.Do(ctx, req, ret); err != nil {
		return errors.Join(creem.ErrListTransaction, err)
	}
	return nil
}
