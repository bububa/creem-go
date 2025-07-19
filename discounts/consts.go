package discounts

// DiscountType The type of the discount, either "percentage" or "fixed".
type DiscountType string

const (
	PercentageDiscount DiscountType = "percentage"
	FixedDiscount      DiscountType = "fixed"
)

// Duration The duration type for the discount.
type Duration string

const (
	Forever   Duration = "forever"
	Once      Duration = "once"
	Repeating Duration = "repeating"
)

// Status The status of the discount (e.g., active, inactive).
// Available options: active, draft, expired, scheduled
type Status string

const (
	Active    Status = "active"
	Draft     Status = "draft"
	Expired   Status = "expired"
	Scheduled Status = "scheduled"
)
