// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewAddMemberParams creates a new AddMemberParams object
//
// There are no default values defined in the spec.
func NewAddMemberParams() AddMemberParams {

	return AddMemberParams{}
}

// AddMemberParams contains all the bound params for the add member operation
// typically these are obtained from a http.Request
//
// swagger:parameters AddMember
type AddMemberParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Teamname of the team to modify
	  Required: true
	  In: path
	*/
	Teamname string
	/*Username of the user to add
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddMemberParams() beforehand.
func (o *AddMemberParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTeamname, rhkTeamname, _ := route.Params.GetOK("teamname")
	if err := o.bindTeamname(rTeamname, rhkTeamname, route.Formats); err != nil {
		res = append(res, err)
	}

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTeamname binds and validates parameter Teamname from path.
func (o *AddMemberParams) bindTeamname(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Teamname = raw

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *AddMemberParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
