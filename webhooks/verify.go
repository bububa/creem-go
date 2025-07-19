package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
)

func generateSignature(r io.Reader, key []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, key)
	_, err := io.Copy(mac, r)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

func verifySignature(r io.Reader, key, signature []byte) (bool, error) {
	computedMAC, err := generateSignature(r, key)
	if err != nil {
		return false, err
	}
	return hmac.Equal(computedMAC, signature), nil
}
