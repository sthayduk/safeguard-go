package models

import (
	"encoding/json"

	"github.com/sthayduk/safeguard-go/src/client"
)

type TypeReferenceName string

const (
	Unknown          TypeReferenceName = "Unknown"
	LocalMachine     TypeReferenceName = "LocalMachine"
	Certificate      TypeReferenceName = "Certificate"
	DirectoryAccount TypeReferenceName = "DirectoryAccount"
	ExternalFed      TypeReferenceName = "ExternalFederation"
	Radius           TypeReferenceName = "Radius"
	OneLoginMfa      TypeReferenceName = "OneLoginMfa"
	Fido2            TypeReferenceName = "Fido2"
	Starling         TypeReferenceName = "Starling"
)

type AuthenticationProvider struct {
	client *client.SafeguardClient

	Id                 int    `json:"Id,omitempty"`
	Name               string `json:"Name,omitempty"`
	TypeReferenceName  string `json:"TypeReferenceName,omitempty"`
	IdentityProviderId int    `json:"IdentityProviderId,omitempty"`
	Identity           string `json:"Identity"`
	RstsProviderId     string `json:"RstsProviderId,omitempty"`
	RstsProviderScope  string `json:"RstsProviderScope,omitempty"`
	ForceAsDefault     bool   `json:"ForceAsDefault,omitempty"`
}

func (a AuthenticationProvider) ToJson() (string, error) {
	userJSON, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(userJSON), nil
}
