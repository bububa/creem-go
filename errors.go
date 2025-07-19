package creem

import "errors"

var (
	ErrUnknown           = errors.New("unknown error")
	ErrInvalidParameters = errors.New("check that the parameters were correct")
	ErrMissingAPIKey     = errors.New("the API key used was missing")
	ErrInvalidAPIKey     = errors.New("the API key used was invalid")
	ErrNoResourceFound   = errors.New("the resource was not found")
	ErrExceededRateLimit = errors.New("the rate limit was exceeded")
	ErrServerError       = errors.New("indicates an error with Creem servers")

	ErrCreateCheckoutSession = errors.New("create checkout session failed")
	ErrGetCheckoutSession    = errors.New("get checkout session failed")

	ErrCreateProduct = errors.New("create product failed")
	ErrGetProduct    = errors.New("get product failed")
	ErrListProduct   = errors.New("list products failed")

	ErrGetCustomer = errors.New("get customer failed")

	ErrListTransaction = errors.New("list transactions failed")

	ErrValidateLicense   = errors.New("validate license failed")
	ErrActivateLicense   = errors.New("activate license failed")
	ErrDeactivateLicense = errors.New("deactivate license failed")

	ErrCreateDiscount = errors.New("create discount failed")
	ErrDeleteDiscount = errors.New("delete discount failed")
	ErrGetDiscount    = errors.New("get discount failed")
)
