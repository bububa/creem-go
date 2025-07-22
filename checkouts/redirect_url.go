package checkouts

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
)

type RedirectParams struct {
	// CheckoutID The ID of the checkout session created for this payment.
	CheckoutID string `json:"checkout_id,omitempty"`
	// OrderID The ID of the order created after successful payment.
	OrderID string `json:"order_id,omitempty"`
	// CustomerID The customer ID, based on the email that executed the successful payment.
	CustomerID string `json:"customer_id,omitempty"`
	// SubscriptionID The subscription ID of the product.
	SubscriptionID string `json:"subscription_id,omitempty"`
	// ProductID The product ID that the payment is related to.
	ProductID string `json:"product_id,omitempty"`
	// RequestID Optional The request ID you provided when creating this checkout session.
	RequestID string `json:"request_id,omitempty"`
	// Signature All previous parameters signed by creem using your API-key, verifiable by you.
	Signature string `json:"signature,omitempty"`
}

func VerifyRedirect(key string, req *RedirectParams) bool {
	var sb strings.Builder
	if req.RequestID != "" {
		sb.WriteString("request_id=")
		sb.WriteString(req.RequestID)
		sb.WriteString("|")
	}
	if req.CheckoutID != "" {
		sb.WriteString("checkout_id=")
		sb.WriteString(req.CheckoutID)
		sb.WriteString("|")
	}
	if req.OrderID != "" {
		sb.WriteString("order_id=")
		sb.WriteString(req.OrderID)
		sb.WriteString("|")
	}
	if req.CustomerID != "" {
		sb.WriteString("customer_id=")
		sb.WriteString(req.CustomerID)
		sb.WriteString("|")
	}
	if req.SubscriptionID != "" {
		sb.WriteString("subscription_id=")
		sb.WriteString(req.SubscriptionID)
		sb.WriteString("|")
	}
	if req.ProductID != "" {
		sb.WriteString("product_id=")
		sb.WriteString(req.ProductID)
		sb.WriteString("|")
	}
	sb.WriteString("salt=")
	sb.WriteString(key)
	h := sha256.New()
	io.WriteString(h, sb.String())
	computed := hex.EncodeToString(h.Sum(nil))
	return computed == req.Signature
}
