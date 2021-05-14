// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetPublicEditTestsFromUserParams creates a new GetPublicEditTestsFromUserParams object
//
// There are no default values defined in the spec.
func NewGetPublicEditTestsFromUserParams() GetPublicEditTestsFromUserParams {

	return GetPublicEditTestsFromUserParams{}
}

// GetPublicEditTestsFromUserParams contains all the bound params for the get public edit tests from user operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPublicEditTestsFromUser
type GetPublicEditTestsFromUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	LikeTitle *string
	/*Indicates which element is first returned. In case of tie it unties with newdate first
	  In: query
	*/
	Orderby *string
	/*
	  In: query
	  Collection Format: pipes
	*/
	Tags [][]string
	/*Username of the teacher who owns the tests
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPublicEditTestsFromUserParams() beforehand.
func (o *GetPublicEditTestsFromUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLikeTitle, qhkLikeTitle, _ := qs.GetOK("likeTitle")
	if err := o.bindLikeTitle(qLikeTitle, qhkLikeTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrderby, qhkOrderby, _ := qs.GetOK("orderby")
	if err := o.bindOrderby(qOrderby, qhkOrderby, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
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

// bindLikeTitle binds and validates parameter LikeTitle from query.
func (o *GetPublicEditTestsFromUserParams) bindLikeTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.LikeTitle = &raw

	return nil
}

// bindOrderby binds and validates parameter Orderby from query.
func (o *GetPublicEditTestsFromUserParams) bindOrderby(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Orderby = &raw

	if err := o.validateOrderby(formats); err != nil {
		return err
	}

	return nil
}

// validateOrderby carries on validations for parameter Orderby
func (o *GetPublicEditTestsFromUserParams) validateOrderby(formats strfmt.Registry) error {

	if err := validate.EnumCase("orderby", "query", *o.Orderby, []interface{}{"newDate", "oldDate", "moreFav", "lessFav", "moreTime", "lessTime"}, true); err != nil {
		return err
	}

	return nil
}

// bindTags binds and validates array parameter Tags from query.
//
// Arrays are parsed according to CollectionFormat: "pipes" (defaults to "csv" when empty).
func (o *GetPublicEditTestsFromUserParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvTags string
	if len(rawData) > 0 {
		qvTags = rawData[len(rawData)-1]
	}

	// CollectionFormat: pipes
	tagsIC := swag.SplitByFormat(qvTags, "pipes")
	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR [][]string
	for _, tagsIV := range tagsIC {
		// items.CollectionFormat: csv
		tagsIIC := swag.SplitByFormat(tagsIV, "csv")
		if len(tagsIIC) > 0 {

			var tagsIIR []string
			for _, tagsIIV := range tagsIIC {
				tagsII := tagsIIV

				tagsIIR = append(tagsIIR, tagsII)
			}

			tagsIR = append(tagsIR, tagsIIR)
		}
	}

	o.Tags = tagsIR

	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *GetPublicEditTestsFromUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	return nil
}
