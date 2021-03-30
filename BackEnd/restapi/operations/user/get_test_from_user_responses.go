// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetTestFromUserOKCode is the HTTP code returned for type GetTestFromUserOK
const GetTestFromUserOKCode int = 200

/*GetTestFromUserOK test found

swagger:response getTestFromUserOK
*/
type GetTestFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Test `json:"body,omitempty"`
}

// NewGetTestFromUserOK creates GetTestFromUserOK with default headers values
func NewGetTestFromUserOK() *GetTestFromUserOK {

	return &GetTestFromUserOK{}
}

// WithPayload adds the payload to the get test from user o k response
func (o *GetTestFromUserOK) WithPayload(payload *models.Test) *GetTestFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test from user o k response
func (o *GetTestFromUserOK) SetPayload(payload *models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestFromUserBadRequestCode is the HTTP code returned for type GetTestFromUserBadRequest
const GetTestFromUserBadRequestCode int = 400

/*GetTestFromUserBadRequest Incorrect Request, or invalida data

swagger:response getTestFromUserBadRequest
*/
type GetTestFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestFromUserBadRequest creates GetTestFromUserBadRequest with default headers values
func NewGetTestFromUserBadRequest() *GetTestFromUserBadRequest {

	return &GetTestFromUserBadRequest{}
}

// WithPayload adds the payload to the get test from user bad request response
func (o *GetTestFromUserBadRequest) WithPayload(payload *models.Error) *GetTestFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test from user bad request response
func (o *GetTestFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestFromUserForbiddenCode is the HTTP code returned for type GetTestFromUserForbidden
const GetTestFromUserForbiddenCode int = 403

/*GetTestFromUserForbidden Not authorized to this content

swagger:response getTestFromUserForbidden
*/
type GetTestFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestFromUserForbidden creates GetTestFromUserForbidden with default headers values
func NewGetTestFromUserForbidden() *GetTestFromUserForbidden {

	return &GetTestFromUserForbidden{}
}

// WithPayload adds the payload to the get test from user forbidden response
func (o *GetTestFromUserForbidden) WithPayload(payload *models.Error) *GetTestFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test from user forbidden response
func (o *GetTestFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestFromUserGoneCode is the HTTP code returned for type GetTestFromUserGone
const GetTestFromUserGoneCode int = 410

/*GetTestFromUserGone That user (password and name) does not exist

swagger:response getTestFromUserGone
*/
type GetTestFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTestFromUserGone creates GetTestFromUserGone with default headers values
func NewGetTestFromUserGone() *GetTestFromUserGone {

	return &GetTestFromUserGone{}
}

// WithPayload adds the payload to the get test from user gone response
func (o *GetTestFromUserGone) WithPayload(payload *models.Error) *GetTestFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test from user gone response
func (o *GetTestFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestFromUserInternalServerErrorCode is the HTTP code returned for type GetTestFromUserInternalServerError
const GetTestFromUserInternalServerErrorCode int = 500

/*GetTestFromUserInternalServerError Internal error

swagger:response getTestFromUserInternalServerError
*/
type GetTestFromUserInternalServerError struct {
}

// NewGetTestFromUserInternalServerError creates GetTestFromUserInternalServerError with default headers values
func NewGetTestFromUserInternalServerError() *GetTestFromUserInternalServerError {

	return &GetTestFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetTestFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
