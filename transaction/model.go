package transaction

import (
	"time"

	"github.com/bububa/creem-go"
)

type ListRequest struct {
	// CustomerID The customer id
	CustomerID string `json:"customer_id,omitempty"`
	// OrderID The order id
	OrderID string `json:"order_id,omitempty"`
	// ProductID The product id
	ProductID string `json:"product_id,omitempty"`
	// PageNumber The page number
	PageNumber int64 `json:"page_number,omitempty"`
	// PageSize The the page size
	PageSize int64 `json:"page_size,omitempty"`
}

type ListResult struct {
	// Pagination details for the list
	Pagination *creem.Pagination `json:"pagination,omitempty"`
	// Items List of transactions items
	Items []Transaction `json:"items,omitempty"`
}

type Transaction struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Amount The transaction amount in cents. 1000 = $10.00
	Amount int64 `json:"amount,omitempty"`
	// Currency Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency string `json:"currency,omitempty"`
	// Type The type of transaction. payment(one time payments) and invoice(subscription)
	Type string `json:"type,omitempty"`
	// Status Status of the transaction.
	Status string `json:"status,omitempty"`
	// CreatedAt Creation date of the order as timestamp
	CreatedAt time.Time `json:"created_at,omitzero"`
	// AmountPaid The amount the customer paid in cents. 1000 = $10.00
	AmountPaid int64 `json:"amount_paid,omitempty"`
	// DiscountAmount The discount amount in cents. 1000 = $10.00
	DiscountAmount int64 `json:"discount_amount"`
	// TaxCountry The ISO alpha-2 country code where tax is collected.
	TaxCountry string `json:"tax_country,omitempty"`
	// TaxAmount The sale tax amount in cents. 1000 = $10.00
	TaxAmount string `json:"tax_amount,omitempty"`
	// RefundedAmount The amount that has been refunded in cents. 1000 = $10.00
	RefundedAmount string `json:"refunded_amount,omitempty"`
	// Order The order associated with the transaction.
	Order string `json:"order,omitempty"`
	// Subscription The subscription associated with the transaction.
	Subscription string `json:"subscription,omitempty"`
	// Customer The customer associated with the transaction.
	Customer string `json:"customer,omitempty"`
	// Description The description of the transaction.
	Description string `json:"description,omitempty"`
	// PeriodStart Start period for the invoice as timestamp
	PeriodStart int64 `json:"period_start,omitempty"`
	// PeriodEnd End period for the invoice as timestamp
	PeriodEnd int64 `json:"period_end,omitempty"`
}
