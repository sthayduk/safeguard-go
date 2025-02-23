package models

import "encoding/json"

// DirectoryServiceEntry represents a Generic Directory Service object
type DirectoryServiceEntry struct {
	Name                string                        `json:"Name"`
	DirectoryProperties DirectoryServiceEntryProperty `json:"DirectoryProperties"`
}

// DirectoryServiceEntryProperty represents the directory properties of a directory service entry
type DirectoryServiceEntryProperty struct {
	DirectoryId       int    `json:"DirectoryId"`
	DirectoryName     string `json:"DirectoryName"`
	DomainName        string `json:"DomainName"`
	NetbiosName       string `json:"NetbiosName"`
	DistinguishedName string `json:"DistinguishedName"`
	ObjectGuid        string `json:"ObjectGuid"`
	ObjectSid         string `json:"ObjectSid"`
}

// ToJson converts a DirectoryServiceEntry to its JSON string representation
func (d DirectoryServiceEntry) ToJson() (string, error) {
	directoryServiceEntryJSON, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(directoryServiceEntryJSON), nil
}
