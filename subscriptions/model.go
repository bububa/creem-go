package subscriptions

import (
	"fmt"
	"net/url"
	"time"

	"github.com/bububa/creem-go"
	"github.com/bububa/creem-go/customers"
	"github.com/bububa/creem-go/products"
	"github.com/bububa/creem-go/transactions"
)

type GetRequest struct {
	creem.GetRequest
	ID string `json:"id,omitempty"`
}

func (r GetRequest) Gateway() string {
	return fmt.Sprintf("v1/subscriptions?subscription_id=%s", url.QueryEscape(r.ID))
}

type UpdateRequest struct {
	creem.PostRequest
	ID string `json:"-"`
	// Items List of subscription items to update/create. If no item ID is provided, the item will be created.
	Items []SubscriptionItem `json:"items,omitempty"`
	// UpdateBehavior The update behavior for the subscription (defaults to proration)
	// Available options: proration-charge-immediately, proration-charge, proration-none
	UpdateBehavior UpdateBehavior `json:"update_behavior,omitempty"`
}

func (r UpdateRequest) Gateway() string {
	return fmt.Sprintf("v1/subscription/%s", r.ID)
}

type UpgradeRequest struct {
	creem.PostRequest
	ID string `json:"-"`
	// ProductID The ID of the product to upgrade to
	ProductID string `json:"product_id,omitempty"`
	// UpdateBehavior The update behavior for the subscription (defaults to proration-charge-immediately)
	UpdateBehavior UpdateBehavior `json:"update_behavior,omitempty"`
}

func (r UpgradeRequest) Gateway() string {
	return fmt.Sprintf("v1/subscription/%s/upgrade", r.ID)
}

type CancelRequest struct {
	creem.GetRequest
	ID string `json:"id,omitempty"`
}

func (r CancelRequest) Gateway() string {
	return fmt.Sprintf("v1/subscriptions/%s/cancel", r.ID)
}

type Subscription struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Product The product associated with the subscription.
	Product *products.Product `json:"product,omitempty"`
	// Customer The customer who owns the subscription.
	Customer *customers.Customer `json:"customer,omitempty"`
	// CollectionMethod The method used for collecting payments for the subscription.
	CollectionMethod string `json:"collection_method,omitempty"`
	// Status The current status of the subscription.
	// Available options: active, canceled, unpaid, paused, trialing
	Status Status `json:"status,omitempty"`
	// CreatedAt The date and time when the subscription was created.
	CreatedAt time.Time `json:"created_at,omitzero"`
	// UpdatedAt The date and time when the subscription was last updated.
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	// Items Subscription items.
	Items []SubscriptionItem `json:"items,omitempty"`
	// LastTransactionID The ID of the last paid transaction.
	LastTransactionID string `json:"last_transaction_id,omitempty"`
	// LastTransaction The last paid transaction.
	LastTransaction *transactions.Transaction `json:"last_transaction,omitempty"`
	// LastTransactionDate The date of the last paid transaction.
	LastTransactionDate time.Time `json:"last_transaction_date,omitzero"`
	// NextTransactionDate The date when the next subscription transaction will be charged.
	NextTransactionDate time.Time `json:"next_transaction_date,omitzero"`
	// CurrentPeriodStartDate The start date of the current subscription period.
	CurrentPeriodStartDate time.Time `json:"current_period_start_date,omitzero"`
	// CurrentPeriodEndDate The end date of the current subscription period.
	CurrentPeriodEndDate time.Time `json:"current_period_end_date,omitzero"`
	// CanceledAt The date and time when the subscription was canceled, if applicable.
	CanceledAt time.Time `json:"canceled_at,omitzero"`
}

// SubscriptionItem Subscription item.
type SubscriptionItem struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object’s type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// ProductID The ID of the product associated with the subscription item.
	ProductID string `json:"product_id,omitempty"`
	// PriceID The ID of the price associated with the subscription item.
	PriceID string `json:"price_id,omitempty"`
	// Units The number of units for the subscription item.
	Uints int64 `json:"units,omitempty"`
}
