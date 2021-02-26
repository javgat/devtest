// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SigninUser signin user
//
// swagger:model SigninUser
type SigninUser struct {

	// email
	// Example: carlos@mail.com
	// Required: true
	// Pattern: ^[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+$
	// Format: email
	Email *strfmt.Email `json:"email"`

	// pass
	// Example: password
	// Required: true
	// Pattern: ^.{6,}$
	// Format: password
	Pass *strfmt.Password `json:"pass"`

	// username
	// Example: carlosg72
	// Required: true
	// Pattern: ^[^@]+$
	Username *string `json:"username"`
}

// Validate validates this signin user
func (m *SigninUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SigninUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.Pattern("email", "body", m.Email.String(), `^[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+$`); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SigninUser) validatePass(formats strfmt.Registry) error {

	if err := validate.Required("pass", "body", m.Pass); err != nil {
		return err
	}

	if err := validate.Pattern("pass", "body", m.Pass.String(), `^.{6,}$`); err != nil {
		return err
	}

	if err := validate.FormatOf("pass", "body", "password", m.Pass.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SigninUser) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", *m.Username, `^[^@]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signin user based on context it is used
func (m *SigninUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SigninUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SigninUser) UnmarshalBinary(b []byte) error {
	var res SigninUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
