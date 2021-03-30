// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuestionAnswer question answer
//
// swagger:model QuestionAnswer
type QuestionAnswer struct {

	// id pregunta
	// Example: 1
	IDPregunta int64 `json:"idPregunta,omitempty"`

	// id respuesta
	// Example: 1
	IDRespuesta int64 `json:"idRespuesta,omitempty"`

	// indice opcion
	// Example: 1
	IndiceOpcion int64 `json:"indiceOpcion,omitempty"`

	// respuesta
	// Example: Javadoc
	Respuesta string `json:"respuesta,omitempty"`
}

// Validate validates this question answer
func (m *QuestionAnswer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this question answer based on context it is used
func (m *QuestionAnswer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuestionAnswer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionAnswer) UnmarshalBinary(b []byte) error {
	var res QuestionAnswer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
