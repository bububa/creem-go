package creem

// TextField configuration for type of text field.
type TextField struct {
	// MaxLength Maximum character length constraint for the input.
	MaxLength int `json:"max_length,omitempty"`
	// MinLength Minimum character length requirement for the input.
	MinLength int `json:"min_length,omitempty"`
}

// CustomField collect additional information from your customer using custom fields. Up to 3 fields are supported.
type CustomField struct {
	// Type the type of the field.
	Type CustomFieldType `json:"type,omitempty"`
	// Key Unique key for custom field. Must be unique to this field, alphanumeric, and up to 200 characters.
	// Maximum length: 200
	Key string `json:"key,omitempty"`
	// Label The label for the field, displayed to the customer, up to 50 characters
	Label string `json:"label,omitempty"`
	// Optional Whether the customer is required to complete the field. Defaults to false
	Optional bool `json:"optional,omitempty"`
	// Text configuration for type of text field.
	Text *TextField `json:"text,omitempty"`
}
