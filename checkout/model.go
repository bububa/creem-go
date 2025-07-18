package checkout

import (
	"time"

	"github.com/bububa/creem-go"
	"github.com/bububa/creem-go/licenses"
)

// CreateRequest create checkout request payload
type CreateRequest struct {
	// ProductID the ID of the product associated with the checkout session.
	ProductID string `json:"product_id,omitempty"`
	// RequestID identify and track each checkout request.
	RequestID string `json:"request_id,omitempty"`
	// Units the number of units for the order.
	Units int64 `json:"units,omitempty"`
	// DiscountCode prefill the checkout session with a discount code.
	DiscountCode string `json:"discount_code,omitempty"`
	// customer data for checkout session. This will prefill the customer info on the checkout page
	Customer *Customer `json:"customer,omitempty"`
	// CustomField collect additional information from your customer using custom fields. Up to 3 fields are supported.
	CustomField []CustomField `json:"custom_field,omitempty"`
	// SuccessURL the URL to which the user will be redirected after the checkout process is completed.
	SuccessURL string `json:"success_url,omitempty"`
	// Metadata Metadata for the checkout in the form of key-value pairs
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Customer customer data for checkout session
type Customer struct {
	// ID unique identifier of the customer. You may specify only one of these parameters: id or email.
	ID string `json:"id,omitempty"`
	// Email customer email address. You may only specify one of these parameters: id, email.
	Email string `json:"email,omitempty"`
}

// CustomFieldType the type of the field.
type CustomFieldType string

const (
	TextFieldType CustomFieldType = "text"
)

// TextField configuration for type of text field.
type TextField struct {
	// MaxLength Maximum character length constraint for the input.
	MaxLength int `json:"max_length,omitempty"`
	// MinLength Minimum character length requirement for the input.
	MinLength int `json:"min_length,omitempty"`
}

// CustomField collect additional information from your customer using custom fields. Up to 3 fields are supported.
type CustomField struct {
	// Type the type of the field.
	Type CustomFieldType `json:"type,omitempty"`
	// Key Unique key for custom field. Must be unique to this field, alphanumeric, and up to 200 characters.
	// Maximum length: 200
	Key string `json:"key,omitempty"`
	// Label The label for the field, displayed to the customer, up to 50 characters
	Label string `json:"label,omitempty"`
	// Optional Whether the customer is required to complete the field. Defaults to false
	Optional bool `json:"optional,omitempty"`
	// Text configuration for type of text field.
	Text *TextField `json:"text,omitempty"`
}

type Session struct {
	// ID unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode string representing the environmentt.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Status of the checkout.
	Status string `json:"status,omitempty"`
	// Product The product associated with the checkout sessiont.
	Product string `json:"product,omitempty"`
	// RequestID Identify and track each checkout request.
	RequestID string `json:"request_id,omitempty"`
	// Units The number of units for the of the product.
	Units int64 `json:"units,omitempty"`
	// Order the order associated with the checkout session.
	Order *Order `json:"order,omitempty"`
	// Subscription The subscription associated with the checkout session.
	Subscription string `json:"subscription,omitempty"`
	// Customer The customer associated with the checkout session.
	Customer string `json:"customer,omitempty"`
	// CustomFields Additional information collected from your customer during the checkout process.
	CustomFields []CustomField `json:"custom_fields,omitempty"`
	// CheckoutURL The URL to which the customer will be redirected to complete the payment.
	CheckoutURL string `json:"checkout_url,omitempty"`
	// SuccessURL The URL to which the user will be redirected after the checkout process is completed.
	SuccessURL string `json:"success_url,omitempty"`
	// Feature features issued for the order.
	Feature []Feature `json:"feature,omitempty"`
	// Metadata for the checkout in the form of key-value pairs
	Metadata map[string]any `json:"metadata,omitempty"`
}

// Feature features issued for the order.
type Feature struct {
	// License key issued for the order.
	License *licenses.License `json:"license,omitempty"`
}

type OrderStatus string

const (
	PendingOrder OrderStatus = "pending"
	PaidOrder    OrderStatus = "paid"
)

type OrderType string

const (
	RecurringOrder OrderType = "recurring"
	OnetimeOrder   OrderType = "onetime"
)

// Order the order associated with the checkout session.
type Order struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	// t Available options: test, prod, sandbox
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Product The product associated with the ordert.
	Product string `json:"product,omitempty"`
	// Amount the total amount of the order in cents. 1000 = $10.00
	Amount int64 `json:"amount,omitempty"`
	// Currency Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency string `json:"currencye,omitempty"`
	// Status Current status of the order.
	Status OrderStatus `json:"stats,omitempty"`
	// Type The type of order. This can specify whether it's a regular purchase, subscription, etc.
	Typte OrderType `json:"type,omitempty"`
	// CreateTime Creation date of the order
	CreateTime time.Time `json:"create_time,omitzero"`
	// UpdateTime Last updated date of the order
	UpdateTime time.Time `json:"update_time,omitzero"`
	// Customer The customer who placed the order.
	Customer string `json:"custotmer,omitempty"`
	// Transaction The transaction ID of the order
	Transaction string `json:"transaction,omitempty"`
	// Discount The discount ID of the order
	Discount string `json:"discount,omitempty"`
	// SubTotal The subtotal of the order in cents. 1000 = $10.00
	SubTotal int64 `json:"sub_total,omitempty"`
	// TaxAmount The tax amount of the order in cents. 1000 = $10.00
	TaxAmount int64 `json:"tax_amount,omitempty"`
	// DiscountAmount The discount amount of the order in cents. 1000 = $10.00
	DiscountAmount int64 `json:"discount_amount,omitempty"`
	// AmountDue The amount due for the order in cents. 1000 = $10.00
	AmountDue int64 `json:"amount_due,omitempty"`
	// AmountPaid The amount paid for the order in cents. 1000 = $10.00
	AmountPaid int64 `json:"amount_paid,omitempty"`
	// FxAmount The amount in the foreign currency, if applicable.
	FxAmount int64 `json:"fx_amount,omitempty"`
	// FxCurrency Three-letter ISO code of the foreign currency, if applicable.
	FxCurrency int64 `json:"fx_currency,omitempty"`
	// FxRate The exchange rate used for converting between currencies, if applicable.
	FxRate float64 `json:"fx_rate,omitempty"`
	// Affiliate The affiliate associated with the order, if applicable.
	Affiliate string `json:"affiliate,omitempty"`
}
