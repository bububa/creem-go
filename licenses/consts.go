package licenses

// LicenseStatus The current status of the license key.
type LicenseStatus string

const (
	InactiveLicense LicenseStatus = "inactive"
	ActiveLicense   LicenseStatus = "active"
	ExpiredLicense  LicenseStatus = "expired"
	DisabledLicese  LicenseStatus = "disabled"
)
