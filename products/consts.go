package products

// BillingType Indicates the billing method for the customer. It can either be a recurring billing cycle or a onetime payment.
type BillingType string

const (
	RecurringBilling BillingType = "recurring"
	OnetimeBilling   BillingType = "onetime"
)

// TaxMode Specifies the tax calculation mode for the transaction. If set to "inclusive," the tax is included in the price. If set to "exclusive," the tax is added on top of the price.
type TaxMode string

const (
	TaxInclusive TaxMode = "inclusive"
	TaxExclusive TaxMode = "exclusive"
)

// CustomFieldType the type of the field.
type CustomFieldType string

const (
	TextFieldType CustomFieldType = "text"
)
