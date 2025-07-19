package webhooks

type EventType string

const (
	// CheckoutCompleted A checkout session was completed, returning all the information about the payment and the order created.
	CheckoutCompleted EventType = "checkout.completed"
	// SubscriptionActive Received when a new subscription is created, the payment was successful and Creem collected the payment creating a new subscription object in your account. Use only for synchronization, we encourage using subscription.paid for activating access.
	SubscriptionActive EventType = "subscription.active"
	// SubscriptionPaid A subscription transaction was paid by the customer
	SubscriptionPaid EventType = "subscription.paid"
	// SubscriptionCanceled A subscription transaction was canceled by the customer
	SubscriptionCanceled EventType = "subscription.canceled"
	// SubscriptionExpired The subscription was expired, given that the current_end_period has been reached without a new payment. Payment retries can happen at this stage, and the subscription status will be terminal only when status is changed to canceled.
	SubscriptionExpired EventType = "subscription.expired"
	// SubscriptionUpdate A subscription object was updated
	SubscriptionUpdate EventType = "subscription.update"
	// SubscriptionTrialing A subscription started a trial period
	SubscriptionTrialing EventType = "subscription.trialing"
	// RefundCreated A refund was created by the merchant
	RefundCreated EventType = "refund.created"
	// DisputedCreated A dispute was created by the customer
	DisputedCreated EventType = "dispute.created"
)
