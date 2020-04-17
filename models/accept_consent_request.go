// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AcceptConsentRequest AcceptConsentRequest AcceptConsentRequest AcceptConsentRequest AcceptConsentRequest AcceptConsentRequest The request payload used to accept a consent request.
// swagger:model acceptConsentRequest
type AcceptConsentRequest struct {

	// grant access token audience
	GrantAccessTokenAudience []string `json:"grant_access_token_audience,omitempty"`

	// grant scope
	GrantScope []string `json:"grant_scope,omitempty"`

	// handled at
	// Format: date-time
	HandledAt NullTime `json:"handled_at,omitempty"`

	// Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same
	// client asks the same user for the same, or a subset of, scope.
	Remember bool `json:"remember,omitempty"`

	// RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the
	// authorization will be remembered indefinitely.
	RememberFor int64 `json:"remember_for,omitempty"`

	// session
	Session *ConsentRequestSession `json:"session,omitempty"`
}

// Validate validates this accept consent request
func (m *AcceptConsentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHandledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcceptConsentRequest) validateHandledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.HandledAt) { // not required
		return nil
	}

	if err := m.HandledAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("handled_at")
		}
		return err
	}

	return nil
}

func (m *AcceptConsentRequest) validateSession(formats strfmt.Registry) error {

	if swag.IsZero(m.Session) { // not required
		return nil
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcceptConsentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcceptConsentRequest) UnmarshalBinary(b []byte) error {
	var res AcceptConsentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
