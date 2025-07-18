package licenses

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bububa/creem-go"
)

// Deactivate License Key
func Deactivate(ctx context.Context, clt *creem.Client, req *DeactivateRequest, ret *License) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return errors.Join(creem.ErrDeactivateLicense, err)
	}
	gw := fmt.Sprintf("%s/v1/licenses/deaactivate", clt.BaseURL())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, gw, &buf)
	if err != nil {
		return errors.Join(creem.ErrDeactivateLicense, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrDeactivateLicense, err)
	}
	return nil
}
