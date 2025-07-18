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

// Activate License Key
func Activate(ctx context.Context, clt *creem.Client, req *ActivateRequest, ret *License) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return errors.Join(creem.ErrActivateLicense, err)
	}
	gw := fmt.Sprintf("%s/v1/licenses/activate", clt.BaseURL())
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, gw, &buf)
	if err != nil {
		return errors.Join(creem.ErrActivateLicense, err)
	}
	if err := clt.Do(httpReq, ret); err != nil {
		return errors.Join(creem.ErrActivateLicense, err)
	}
	return nil
}
