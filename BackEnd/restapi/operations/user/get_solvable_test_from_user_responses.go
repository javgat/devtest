// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetSolvableTestFromUserOKCode is the HTTP code returned for type GetSolvableTestFromUserOK
const GetSolvableTestFromUserOKCode int = 200

/*GetSolvableTestFromUserOK publishedTest found

swagger:response getSolvableTestFromUserOK
*/
type GetSolvableTestFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Test `json:"body,omitempty"`
}

// NewGetSolvableTestFromUserOK creates GetSolvableTestFromUserOK with default headers values
func NewGetSolvableTestFromUserOK() *GetSolvableTestFromUserOK {

	return &GetSolvableTestFromUserOK{}
}

// WithPayload adds the payload to the get solvable test from user o k response
func (o *GetSolvableTestFromUserOK) WithPayload(payload *models.Test) *GetSolvableTestFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solvable test from user o k response
func (o *GetSolvableTestFromUserOK) SetPayload(payload *models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolvableTestFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolvableTestFromUserBadRequestCode is the HTTP code returned for type GetSolvableTestFromUserBadRequest
const GetSolvableTestFromUserBadRequestCode int = 400

/*GetSolvableTestFromUserBadRequest Incorrect Request, or invalida data

swagger:response getSolvableTestFromUserBadRequest
*/
type GetSolvableTestFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolvableTestFromUserBadRequest creates GetSolvableTestFromUserBadRequest with default headers values
func NewGetSolvableTestFromUserBadRequest() *GetSolvableTestFromUserBadRequest {

	return &GetSolvableTestFromUserBadRequest{}
}

// WithPayload adds the payload to the get solvable test from user bad request response
func (o *GetSolvableTestFromUserBadRequest) WithPayload(payload *models.Error) *GetSolvableTestFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solvable test from user bad request response
func (o *GetSolvableTestFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolvableTestFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolvableTestFromUserForbiddenCode is the HTTP code returned for type GetSolvableTestFromUserForbidden
const GetSolvableTestFromUserForbiddenCode int = 403

/*GetSolvableTestFromUserForbidden Not authorized to this content

swagger:response getSolvableTestFromUserForbidden
*/
type GetSolvableTestFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolvableTestFromUserForbidden creates GetSolvableTestFromUserForbidden with default headers values
func NewGetSolvableTestFromUserForbidden() *GetSolvableTestFromUserForbidden {

	return &GetSolvableTestFromUserForbidden{}
}

// WithPayload adds the payload to the get solvable test from user forbidden response
func (o *GetSolvableTestFromUserForbidden) WithPayload(payload *models.Error) *GetSolvableTestFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solvable test from user forbidden response
func (o *GetSolvableTestFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolvableTestFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolvableTestFromUserGoneCode is the HTTP code returned for type GetSolvableTestFromUserGone
const GetSolvableTestFromUserGoneCode int = 410

/*GetSolvableTestFromUserGone That resource does not exist

swagger:response getSolvableTestFromUserGone
*/
type GetSolvableTestFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSolvableTestFromUserGone creates GetSolvableTestFromUserGone with default headers values
func NewGetSolvableTestFromUserGone() *GetSolvableTestFromUserGone {

	return &GetSolvableTestFromUserGone{}
}

// WithPayload adds the payload to the get solvable test from user gone response
func (o *GetSolvableTestFromUserGone) WithPayload(payload *models.Error) *GetSolvableTestFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get solvable test from user gone response
func (o *GetSolvableTestFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSolvableTestFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSolvableTestFromUserInternalServerErrorCode is the HTTP code returned for type GetSolvableTestFromUserInternalServerError
const GetSolvableTestFromUserInternalServerErrorCode int = 500

/*GetSolvableTestFromUserInternalServerError Internal error

swagger:response getSolvableTestFromUserInternalServerError
*/
type GetSolvableTestFromUserInternalServerError struct {
}

// NewGetSolvableTestFromUserInternalServerError creates GetSolvableTestFromUserInternalServerError with default headers values
func NewGetSolvableTestFromUserInternalServerError() *GetSolvableTestFromUserInternalServerError {

	return &GetSolvableTestFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetSolvableTestFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}