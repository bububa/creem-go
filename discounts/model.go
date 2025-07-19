package discounts

import (
	"fmt"
	"net/url"
	"time"

	"github.com/bububa/creem-go"
)

type CreateRequest struct {
	creem.PostRequest
	// Name The name of the discount.
	Name string `json:"name,omitempty"`
	// Type The type of the discount, either "percentage" or "fixed".
	Type DiscountType `json:"type,omitempty"`
	// Duration The duration type for the discount.
	Duration Duration `json:"duration,omitempty"`
	// AppliesToProduct The list of product IDs to which this discount applies.
	AppliesToProduct string `json:"applies_to_product,omitempty"`
	// Code Optional discount code. If left empty, a code will be generated.
	Code string `json:"code,omitempty"`
	// Amount The fixed value for the discount. Only applicable if the type is "fixed".
	Amount int64 `json:"amount,omitempty"`
	// Currency The currency of the discount. Only required if type is "fixed".
	Currency string `json:"currency,omitempty"`
	// Percentage The percentage value for the discount. Only applicable if the type is "percentage".
	Percentage int `json:"percentage,omitempty"`
	// ExpireDate The expiry date of the discount.
	ExpireDate time.Time `json:"expire_date,omitzero"`
	// MaxRedemptions The maximum number of redemptions for the discount.
	MaxRedemptions int64 `json:"max_redemptions,omitempty"`
	// DurationInMonth The number of months the discount is valid for. Only applicable if the duration is "repeating" and the product is a subscription.
	DurationInMonth int `json:"duration_in_month,omitempty"`
}

func (r CreateRequest) Gateway() string {
	return "v1/discounts"
}

type DeleteRequest struct {
	creem.DeleteRequest
	ID string `json:"id,omitempty"`
}

func (r DeleteRequest) Gateway() string {
	return fmt.Sprintf("v1/discounts/%s/delete", r.ID)
}

type GetRequest struct {
	creem.GetRequest
	DiscountID   string `json:"discount_id,omitempty"`
	DiscountCode string `json:"discount_code,omitempty"`
}

func (r GetRequest) Gateway() string {
	values := url.Values{}
	if r.DiscountID != "" {
		values.Set("discount_id", r.DiscountID)
	}
	if r.DiscountCode != "" {
		values.Set("discount_code", r.DiscountCode)
	}
	return fmt.Sprintf("v1/discounts?%s", values.Encode())
}

type Discount struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object A string representing the object’s type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Status The status of the discount (e.g., active, inactive).
	// Available options: active, draft, expired, scheduled
	Status Status `json:"status,omitempty"`
	// Name The name of the discount.
	Name string `json:"name,omitempty"`
	// Code The discount code. A unique identifier for the discount.
	Code string `json:"code,omitempty"`
	// Type The type of the discount, either "percentage" or "fixed".
	Type DiscountType `json:"type,omitempty"`
	// Amount The amount of the discount. Can be a percentage or a fixed amount.
	Amount int64 `json:"discount,omitempty"`
	// Curerency The currency of the discount. Only required if type is "fixed".
	Currency string `json:"currency,omitempty"`
	// Percentage The percentage value for the discount. Only applicable if the type is "percentage".
	Percentage int `json:"percentage,omitempty"`
	// ExpireDate The expiry date of the discount.
	ExpireDate time.Time `json:"expire_date,omitzero"`
	// MaxRedemptions The maximum number of redemptions for the discount.
	MaxRedemptions int64 `json:"max_redemptions,omitempty"`
	// DurationInMonth The number of months the discount is valid for. Only applicable if the duration is "repeating" and the product is a subscription.
	DurationInMonth int `json:"duration_in_month,omitempty"`
	// AppliesToProduct The list of product IDs to which this discount applies.
	AppliesToProduct string `json:"applies_to_product,omitempty"`
}
