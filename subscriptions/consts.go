package subscriptions

// Status The current status of the subscription.
// Available options: active, canceled, unpaid, paused, trialing
type Status string

const (
	Active   Status = "active"
	Canceled Status = "canceled"
	Unpaid   Status = "unpaid"
	Paused   Status = "paused"
	Trialing Status = "trialing"
)

// UpdateBehavior The update behavior for the subscription (defaults to proration)
// Available options: proration-charge-immediately, proration-charge, proration-none
type UpdateBehavior string

const (
	ProrationChargeImmediately UpdateBehavior = "proration-charge-immediately"
	ProprationCharge           UpdateBehavior = "proration-charge"
	ProrationNone              UpdateBehavior = "proration-none"
)
