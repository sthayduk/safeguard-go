package pkg

// Preference represents a user-specific application setting or preference.
//
// Preferences are key-value pairs that can be used to store user-specific
// settings like UI preferences, default views, or custom configurations.
//
// Example:
//
//	pref := Preference{
//	    Name: "DefaultView",
//	    Value: "grid"
//	}
type Preference struct {
	Name  string `json:"Name,omitempty"`  // The unique identifier/key of the preference
	Value string `json:"Value,omitempty"` // The value/setting of the preference
}
