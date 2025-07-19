package customers

import (
	"fmt"
	"net/url"
	"time"

	"github.com/bububa/creem-go"
)

type GetRequest struct {
	creem.GetRequest
	// CustomerID The unique identifier of the customer
	CustomerID string `json:"customer_id,omitempty"`
	// Email The unique email of the customer
	Email string `json:"email,omitempty"`
}

func (r GetRequest) Gateway() string {
	values := url.Values{}
	if r.CustomerID != "" {
		values.Set("customer_id", r.CustomerID)
	}
	if r.Email != "" {
		values.Set("email", r.Email)
	}
	return fmt.Sprintf("v1/customers?%s", values.Encode())
}

type Customer struct {
	// CustomerID The unique identifier of the customer
	CustomerID string `json:"customer_id,omitempty"`
	// Email The unique email of the customer
	Email string `json:"email,omitempty"`
	// Name Customer name.
	Name string `json:"name,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object’s type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Country The ISO alpha-2 country code for the customer.
	Country string `json:"country,omitempty"`
	// CreatedAt Creation date of the product
	CreatedAt time.Time `json:"created_at,omitzero"`
	// UpdatedAt Last updated date of the product
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}
