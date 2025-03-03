package safeguard

import "encoding/json"

// DirectoryServiceEntry represents a Generic Directory Service object
// containing information about an entry in a directory service such as
// Active Directory or LDAP.
type DirectoryServiceEntry struct {
	Name                string                        `json:"Name"`
	DirectoryProperties DirectoryServiceEntryProperty `json:"DirectoryProperties"`
}

// DirectoryServiceEntryProperty represents the directory-specific properties
// of a directory service entry, including identifiers and names used by
// the directory service.
type DirectoryServiceEntryProperty struct {
	DirectoryId       int    `json:"DirectoryId"`
	DirectoryName     string `json:"DirectoryName"`
	DomainName        string `json:"DomainName"`
	NetbiosName       string `json:"NetbiosName"`
	DistinguishedName string `json:"DistinguishedName"`
	ObjectGuid        string `json:"ObjectGuid"`
	ObjectSid         string `json:"ObjectSid"`
}

// ToJson serializes a DirectoryServiceEntry instance to its JSON representation.
//
// This method converts the DirectoryServiceEntry and all its nested structures
// into a JSON string that can be used for transmission or storage.
//
// Returns:
//   - string: A JSON-encoded string representation of the directory service entry
//   - error: An error if JSON marshaling encounters any issues with the data structures
func (d DirectoryServiceEntry) ToJson() (string, error) {
	directoryServiceEntryJSON, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(directoryServiceEntryJSON), nil
}
