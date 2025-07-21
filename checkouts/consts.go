package checkouts

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

type CheckoutStatus string

const (
	PendingCheckout   CheckoutStatus = "pending"
	CompletedCheckout CheckoutStatus = "completed"
)
