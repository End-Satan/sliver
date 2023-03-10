package webauthn

// Code generated by cdproto-gen. DO NOT EDIT.

// EventCredentialAdded triggered when a credential is added to an
// authenticator.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#event-credentialAdded
type EventCredentialAdded struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}

// EventCredentialAsserted triggered when a credential is used in a webauthn
// assertion.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAuthn#event-credentialAsserted
type EventCredentialAsserted struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"`
	Credential      *Credential     `json:"credential"`
}
