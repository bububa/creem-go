package checkouts

// CustomFieldType the type of the field.
type CustomFieldType string

const (
	TextFieldType CustomFieldType = "text"
)

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
