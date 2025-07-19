package products

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/bububa/creem-go"
)

type CreateRequest struct {
	creem.PostRequest
	// Name of the product
	Name string `json:"name,omitempty"`
	// Price The price of the product in cents
	// Required range: x >= 100
	Price int64 `json:"price,omitempty"`
	// Currency Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency string `json:"currency,omitempty"`
	// BillingType Indicates the billing method for the customer. It can either be a recurring billing cycle or a onetime payment.
	BillingType BillingType `json:"billing_type,omitempty"`
	// Description of the product
	Description string `json:"description,omitempty"`
	// ImageURL URL of the product image
	ImageURL string `json:"image_url,omitempty"`
	// BillingPeriod Billing period, required if billing_type is recurring
	BillingPeriod string `json:"billing_period,omitempty"`
	// TaxMode Specifies the tax calculation mode for the transaction. If set to "inclusive," the tax is included in the price. If set to "exclusive," the tax is added on top of the price.
	TaxMode TaxMode `json:"tax_mode,omitempty"`
	// TaxCategory Categorizes the type of product or service for tax purposes. This helps determine the applicable tax rules based on the nature of the item or service.
	TaxCategory string `json:"tax_category,omitempty"`
	// DefaultSuccessURL The URL to which the user will be redirected after successfull payment.
	DefaultSuccessURL string `json:"default_success_url,omitempty"`
	// CustomField Collect additional information from your customer using custom fields during checkout. Up to 3 fields are supported.
	CustomField []creem.CustomField `json:"custom_field,omitempty"`
}

func (r CreateRequest) Gateway() string {
	return "v1/products"
}

type GetRequest struct {
	creem.GetRequest
	ID string `json:"id,omitempty"`
}

func (r GetRequest) Gateway() string {
	return fmt.Sprintf("v1/products?product_id=%s", url.QueryEscape(r.ID))
}

type ListRequest struct {
	creem.GetRequest
	// PageNumber The page number
	PageNumber int64 `json:"page_number,omitempty"`
	// PageSize The the page size
	PageSize int64 `json:"page_size,omitempty"`
}

func (r ListRequest) Gateway() string {
	values := url.Values{}
	values.Set("page_number", strconv.FormatInt(r.PageNumber, 10))
	values.Set("page_size", strconv.FormatInt(r.PageSize, 10))
	return fmt.Sprintf("v1/products?%s", values.Encode())
}

// TextField configuration for type of text field.
type TextField struct {
	// MaxLength Maximum character length constraint for the input.
	MaxLength int `json:"max_length,omitempty"`
	// MinLength Minimum character length requirement for the input.
	MinLength int `json:"min_length,omitempty"`
}

type Product struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environmentt.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Name The name of the product
	Name string `json:"name,omitempty"`
	// Description of the product
	Description string `json:"description,omitempty"`
	// Price The price of the product in cents
	Price int64 `json:"price,omitempty"`
	// Currency Three-letter ISO currency code, in uppercase. Must be a supported currency.
	Currency string `json:"currency,omitempty"`
	// BillingType Indicates the billing method for the customer. It can either be a recurring billing cycle or a onetime payment.
	BillingType BillingType `json:"billing_type,omitempty"`
	// Status of the product
	Status string `json:"status,omitempty"`
	// TaxMode Specifies the tax calculation mode for the transaction. If set to "inclusive," the tax is included in the price. If set to "exclusive," the tax is added on top of the price.
	TaxMode TaxMode `json:"tax_mode,omitempty"`
	// TaxCategory Categorizes the type of product or service for tax purposes. This helps determine the applicable tax rules based on the nature of the item or service.
	TaxCategory string `json:"tax_category,omitempty"`
	// CreatedAt Creation date of the product
	CreatedAt time.Time `json:"created_at,omitzero"`
	// UpdatedAt Last updated date of the product
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	// ImageURL URL of the product image. Only png as jpg are supported
	ImageURL string `json:"image_url,omitempty"`
	// ProductURL The product page you can redirect your customers to for express checkout.
	ProductURL string `json:"product_url,omitempty"`
	// DefaultSuccessURL The URL to which the user will be redirected after successfull payment.
	DefaultSuccessURL string `json:"default_success_url,omitempty"`
	// Features of the product.
	Features []Feature `json:"features,omitempty"`
}

// Feature of the product.
type Feature struct {
	// ID Unique identifier for the feature.
	ID string `json:"id,omitempty"`
	// Type The feature type.
	Type string `json:"type,omitempty"`
	// Description A brief description of the feature
	Description string `json:"description,omitempty"`
}

type ListResult struct {
	// Pagination details for the list
	Pagination *creem.Pagination `json:"pagination,omitempty"`
	// Items List of product items
	Items []Product `json:"items,omitempty"`
}
