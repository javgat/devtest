// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Role role
//
// swagger:model Role
type Role struct {

	// rol
	// Required: true
	// Enum: [estudiante profesor administrador]
	Rol *string `json:"rol"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var roleTypeRolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["estudiante","profesor","administrador"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		roleTypeRolPropEnum = append(roleTypeRolPropEnum, v)
	}
}

const (

	// RoleRolEstudiante captures enum value "estudiante"
	RoleRolEstudiante string = "estudiante"

	// RoleRolProfesor captures enum value "profesor"
	RoleRolProfesor string = "profesor"

	// RoleRolAdministrador captures enum value "administrador"
	RoleRolAdministrador string = "administrador"
)

// prop value enum
func (m *Role) validateRolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, roleTypeRolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Role) validateRol(formats strfmt.Registry) error {

	if err := validate.Required("rol", "body", m.Rol); err != nil {
		return err
	}

	// value enum
	if err := m.validateRolEnum("rol", "body", *m.Rol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role based on context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
