package licenses

import (
	"time"

	"github.com/bububa/creem-go"
)

type ValidateRequest struct {
	// Key The license key to validate.
	Key string `json:"key,omitempty"`
	// InstanceID Id of the instance to validate.
	InstanceID string `json:"instance_id,omitempty"`
}

type ActivateRequest struct {
	// Key The license key to activate.
	Key string `json:"key,omitempty"`
	// InstanceName A label for the new instance to identify it in Creem.
	InstanceName string `json:"instance_name,omitempty"`
}

type DeactivateRequest struct {
	// Key The license key to deactivate.
	Key string `json:"key,omitempty"`
	// InstanceID Id of the instance to deactivate.
	InstanceID string `json:"instance_id,omitempty"`
}

// License key issued
type License struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object A string representing the object’s type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Status The current status of the license key.
	Status LicenseStatus `json:"status,omitempty"`
	// Key The license key.
	Key string `json:"key,omitempty"`
	// Activation The number of instances that this license key was activated.
	Activation int64 `json:"activation,omitempty"`
	// CreatedAt The creation date of the license key.
	CreatedAt time.Time `json:"created_at,omitzero"`
	// ActivationLimit The activation limit. Null if activations are unlimited.
	ActivationLimit int64 `json:"activation_limit,omitzero"`
	// ExpiresAt The date the license key expires. Null if it does not have an expiration date.
	ExipresAt time.Time `json:"expires_at,omitzero"`
	// Instance Associated license instances.
	Instance *Instance `json:"instance,omitempty"`
}

// LicenseStatus The current status of the license key.
type LicenseStatus string

const (
	InactiveLicense LicenseStatus = "inactive"
	ActiveLicense   LicenseStatus = "active"
	ExpiredLicense  LicenseStatus = "expired"
	DisabledLicese  LicenseStatus = "disabled"
)

// Instance Associated license instances.
type Instance struct {
	// ID Unique identifier for the object.
	ID string `json:"id,omitempty"`
	// Mode String representing the environment.
	Mode creem.Mode `json:"mode,omitempty"`
	// Object A string representing the object’s type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// Name The name of the license instance.
	Name string `json:"name,omitempty"`
	// Status The status of the license instance.
	Status LicenseStatus `json:"status,omitempty"`
	// CreatedAt The creation date of the license instance.
	CreatedAt time.Time `json:"created_at,omitzero"`
}
