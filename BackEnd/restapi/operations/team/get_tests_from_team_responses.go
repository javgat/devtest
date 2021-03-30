// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTestsFromTeamOKCode is the HTTP code returned for type GetTestsFromTeamOK
const GetTestsFromTeamOKCode int = 200

/*GetTestsFromTeamOK tests found

swagger:response getTestsFromTeamOK
*/
type GetTestsFromTeamOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetTestsFromTeamOK creates GetTestsFromTeamOK with default headers values
func NewGetTestsFromTeamOK() *GetTestsFromTeamOK {

	return &GetTestsFromTeamOK{}
}

// WithPayload adds the payload to the get tests from team o k response
func (o *GetTestsFromTeamOK) WithPayload(payload []*models.Test) *GetTestsFromTeamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from team o k response
func (o *GetTestsFromTeamOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromTeamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Test, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTestsFromTeamBadRequestCode is the HTTP code returned for type GetTestsFromTeamBadRequest
const GetTestsFromTeamBadRequestCode int = 400

/*GetTestsFromTeamBadRequest Incorrect Request, or invalida data

swagger:response getTestsFromTeamBadRequest
*/
type GetTestsFromTeamBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromTeamBadRequest creates GetTestsFromTeamBadRequest with default headers values
func NewGetTestsFromTeamBadRequest() *GetTestsFromTeamBadRequest {

	return &GetTestsFromTeamBadRequest{}
}

// WithPayload adds the payload to the get tests from team bad request response
func (o *GetTestsFromTeamBadRequest) WithPayload(payload *models.Error) *GetTestsFromTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from team bad request response
func (o *GetTestsFromTeamBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromTeamForbiddenCode is the HTTP code returned for type GetTestsFromTeamForbidden
const GetTestsFromTeamForbiddenCode int = 403

/*GetTestsFromTeamForbidden Not authorized to this content

swagger:response getTestsFromTeamForbidden
*/
type GetTestsFromTeamForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromTeamForbidden creates GetTestsFromTeamForbidden with default headers values
func NewGetTestsFromTeamForbidden() *GetTestsFromTeamForbidden {

	return &GetTestsFromTeamForbidden{}
}

// WithPayload adds the payload to the get tests from team forbidden response
func (o *GetTestsFromTeamForbidden) WithPayload(payload *models.Error) *GetTestsFromTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from team forbidden response
func (o *GetTestsFromTeamForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromTeamGoneCode is the HTTP code returned for type GetTestsFromTeamGone
const GetTestsFromTeamGoneCode int = 410

/*GetTestsFromTeamGone That user (password and name) does not exist

swagger:response getTestsFromTeamGone
*/
type GetTestsFromTeamGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestsFromTeamGone creates GetTestsFromTeamGone with default headers values
func NewGetTestsFromTeamGone() *GetTestsFromTeamGone {

	return &GetTestsFromTeamGone{}
}

// WithPayload adds the payload to the get tests from team gone response
func (o *GetTestsFromTeamGone) WithPayload(payload *models.Error) *GetTestsFromTeamGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tests from team gone response
func (o *GetTestsFromTeamGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestsFromTeamGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestsFromTeamInternalServerErrorCode is the HTTP code returned for type GetTestsFromTeamInternalServerError
const GetTestsFromTeamInternalServerErrorCode int = 500

/*GetTestsFromTeamInternalServerError Internal error

swagger:response getTestsFromTeamInternalServerError
*/
type GetTestsFromTeamInternalServerError struct {
}

// NewGetTestsFromTeamInternalServerError creates GetTestsFromTeamInternalServerError with default headers values
func NewGetTestsFromTeamInternalServerError() *GetTestsFromTeamInternalServerError {

	return &GetTestsFromTeamInternalServerError{}
}

// WriteResponse to the client
func (o *GetTestsFromTeamInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
