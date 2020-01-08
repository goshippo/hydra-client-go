// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IPNet IPNet IPNet IPNet IPNet IPNet An IPNet represents an IP network.
// swagger:model IPNet
type IPNet struct {

	// IP
	IP IP `json:"IP,omitempty"`

	// mask
	Mask IPMask `json:"Mask,omitempty"`
}

// Validate validates this IP net
func (m *IPNet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPNet) validateIP(formats strfmt.Registry) error {

	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := m.IP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("IP")
		}
		return err
	}

	return nil
}

func (m *IPNet) validateMask(formats strfmt.Registry) error {

	if swag.IsZero(m.Mask) { // not required
		return nil
	}

	if err := m.Mask.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Mask")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPNet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPNet) UnmarshalBinary(b []byte) error {
	var res IPNet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
