package license

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bububa/creem-go"
)

// Validate License Key
func Validate(ctx context.Context, clt *creem.Client, req *ValidateRequest, ret *License) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return errors.Join(creem.ErrValidateLicense, err)
	}
	gw := fmt.Sprintf("%s/v1/licenses/validate", clt.BaseURL())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, gw, &buf)
	if err != nil {
		return errors.Join(creem.ErrValidateLicense, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrValidateLicense, err)
	}
	return nil
}
