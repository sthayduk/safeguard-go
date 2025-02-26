package pkg

import "time"

type Fido2Authenticator struct {
	CredentialId          string    `json:"CredentialId,omitempty"`
	DateRegistered        time.Time `json:"DateRegistered,omitempty"`
	DateLastAuthenticated time.Time `json:"DateLastAuthenticated,omitempty"`
	Name                  string    `json:"Name,omitempty"`
}
