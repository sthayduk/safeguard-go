package models

type AuthenticationProvider struct {
	Id                 int    `json:"Id,omitempty"`
	Name               string `json:"Name,omitempty"`
	TypeReferenceName  string `json:"TypeReferenceName,omitempty"`
	IdentityProviderId int    `json:"IdentityProviderId,omitempty"`
	RstsProviderId     string `json:"RstsProviderId,omitempty"`
	RstsProviderScope  string `json:"RstsProviderScope,omitempty"`
	ForceAsDefault     bool   `json:"ForceAsDefault,omitempty"`
}
