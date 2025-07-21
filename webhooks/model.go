package webhooks

import (
	"encoding/json"
	"time"

	"github.com/bububa/creem-go"
	"github.com/bububa/creem-go/checkouts"
	"github.com/bububa/creem-go/customers"
	"github.com/bububa/creem-go/products"
	"github.com/bububa/creem-go/subscriptions"
	"github.com/bububa/creem-go/transactions"
)

type Event struct {
	// ID event id
	ID string `json:"id,omitempty"`
	// EventType type of the event
	EventType EventType `json:"eventType,omitempty"`
	// CreatedAt event created timestamp
	CreatedAt int64 `json:"created_at,omitempty"`
	// Object event payload
	Object json.RawMessage `json:"object,omitempty"`
}

// CheckoutCompletedPayload A checkout session was completed, returning all the information about the payment and the order created.
type CheckoutCompletedPayload struct {
	ID           string                      `json:"id,omitempty"`
	Object       string                      `json:"object,omitempty"`
	RequestID    string                      `json:"request_id,omitempty"`
	Order        *checkouts.Order            `json:"order,omitempty"`
	Product      *products.Product           `json:"product,omitempty"`
	Customer     *customers.Customer         `json:"customer,omitempty"`
	Subscription *subscriptions.Subscription `json:"subscription,omitempty"`
	CustomFields []creem.CustomField         `json:"custom_fields,omitempty"`
	Units        int64                       `json:"units,omitempty"`
	SuccessURL   string                      `json:"success_url,omitempty"`
	Status       string                      `json:"status,omitempty"`
	Metadata     map[string]any              `json:"metadata,omitempty"`
	Mode         creem.Mode                  `json:"mode,omitempty"`
}

// SubscriptionPayload common payload in subscription events
type SubscriptionPayload struct {
	ID               string              `json:"id,omitempty"`
	Object           string              `json:"object,omitempty"`
	Product          *products.Product   `json:"product,omitempty"`
	Customer         *customers.Customer `json:"customer,omitempty"`
	CollectionMethod string              `json:"collection_method,omitempty"`
	Status           string              `json:"status,omitempty"`
	CanceledAt       time.Time           `json:"canceled_at,omitzero"`
	CreatedAt        time.Time           `json:"created_at,omitzero"`
	UpdatedAt        time.Time           `json:"updated_at,omitzero"`
	Metadata         map[string]any      `json:"metadata,omitempty"`
	Mode             creem.Mode          `json:"mode,omitempty"`
}

// SubscriptionActivePayload Received when a new subscription is created, the payment was successful and Creem collected the payment creating a new subscription object in your account. Use only for synchronization, we encourage using subscription.paid for activating access.
type SubscriptionActivePayload SubscriptionPayload

// SubscriptionPaidPayload A subscription transaction was paid by the customer
type SubscriptionPaidPayload struct {
	SubscriptionPayload
	LastTransactionID      string    `json:"last_transaction_id,omitempty"`
	LastTransactionDate    time.Time `json:"last_transaction_date,omitzero"`
	NextTransactionDate    time.Time `json:"next_transaction_date,omitzero"`
	CurrentPeriodStartDate time.Time `json:"current_period_start_date,omitzero"`
	CurrentPeriodEndDate   time.Time `json:"current_period_end_date,omitzero"`
}

// SubscriptionCanceledPayload The subscription was canceled by the merchant or by the customer.
type SubscriptionCanceledPayload struct {
	SubscriptionPayload
	LastTransactionID      string    `json:"last_transaction_id,omitempty"`
	LastTransactionDate    time.Time `json:"last_transaction_date,omitzero"`
	CurrentPeriodStartDate time.Time `json:"current_period_start_date,omitzero"`
	CurrentPeriodEndDate   time.Time `json:"current_period_end_date,omitzero"`
}

// SubscriptionExpiredPayload The subscription was expired, given that the current_end_period has been reached without a new payment. Payment retries can happen at this stage, and the subscription status will be terminal only when status is changed to canceled.
type SubscriptionExpiredPayload SubscriptionPaidPayload

// SubscriptionUpdatePayload A subscription object was updated
type SubscriptionUpdatePayload struct {
	SubscriptionPayload
	Items                  []subscriptions.SubscriptionItem `json:"items,omitempty"`
	CurrentPeriodStartDate time.Time                        `json:"current_period_start_date,omitzero"`
	CurrentPeriodEndDate   time.Time                        `json:"current_period_end_date,omitzero"`
}

// SubscriptionTrialingPayload A subscription started a trial period
type SubscriptionTrialingPayload SubscriptionUpdatePayload

// RefundCreatedPayload A refund was created by the merchant
type RefundCreatedPayload struct {
	ID             string                      `json:"id,omitempty"`
	Object         string                      `json:"object,omitempty"`
	Status         string                      `json:"status,omitempty"`
	RefundAmount   int64                       `json:"refund_amount,omitempty"`
	RefundCurrency string                      `json:"refund_currency,omitempty"`
	Reason         string                      `json:"reason,omitempty"`
	Transaction    *transactions.Transaction   `json:"transaction,omitempty"`
	Subscription   *subscriptions.Subscription `json:"subscriptions,omitempty"`
	Checkout       *checkouts.Checkout         `json:"checkout,omitempty"`
	Order          *checkouts.Order            `json:"order,omitempty"`
	Customer       *customers.Customer         `json:"customer,omitempty"`
	CreatedAt      time.Time                   `json:"created_at,omitzero"`
	Mode           creem.Mode                  `json:"mode,omitempty"`
}

// DisputeCreate A dispute was created by the customer
type DisputeCreate struct {
	ID           string                      `json:"id,omitempty"`
	Object       string                      `json:"object,omitempty"`
	Amount       int64                       `json:"amount,omitempty"`
	Currency     string                      `json:"currency,omitempty"`
	Transaction  *transactions.Transaction   `json:"transaction,omitempty"`
	Subscription *subscriptions.Subscription `json:"subscription,omitempty"`
	Checkout     *checkouts.Checkout         `json:"checkout,omitempty"`
	Order        *checkouts.Order            `json:"order,omitempty"`
	Customer     *customers.Customer         `json:"customer,omitempty"`
	CreatedAt    time.Time                   `json:"created_at,omitzero"`
	Mode         creem.Mode                  `json:"mode,omitempty"`
}
