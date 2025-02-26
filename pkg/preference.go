package pkg

// Represents an application user's preference.
type Preference struct {
	Name  string `json:"Name,omitempty"`
	Value string `json:"Value,omitempty"`
}
