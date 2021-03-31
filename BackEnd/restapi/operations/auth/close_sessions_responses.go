// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// CloseSessionsOKCode is the HTTP code returned for type CloseSessionsOK
const CloseSessionsOKCode int = 200

/*CloseSessionsOK Sessions deleted successfully

swagger:response closeSessionsOK
*/
type CloseSessionsOK struct {
}

// NewCloseSessionsOK creates CloseSessionsOK with default headers values
func NewCloseSessionsOK() *CloseSessionsOK {

	return &CloseSessionsOK{}
}

// WriteResponse to the client
func (o *CloseSessionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// CloseSessionsBadRequestCode is the HTTP code returned for type CloseSessionsBadRequest
const CloseSessionsBadRequestCode int = 400

/*CloseSessionsBadRequest Incorrect Request, or invalida data

swagger:response closeSessionsBadRequest
*/
type CloseSessionsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCloseSessionsBadRequest creates CloseSessionsBadRequest with default headers values
func NewCloseSessionsBadRequest() *CloseSessionsBadRequest {

	return &CloseSessionsBadRequest{}
}

// WithPayload adds the payload to the close sessions bad request response
func (o *CloseSessionsBadRequest) WithPayload(payload *models.Error) *CloseSessionsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the close sessions bad request response
func (o *CloseSessionsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloseSessionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloseSessionsForbiddenCode is the HTTP code returned for type CloseSessionsForbidden
const CloseSessionsForbiddenCode int = 403

/*CloseSessionsForbidden Not authorized to this content

swagger:response closeSessionsForbidden
*/
type CloseSessionsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCloseSessionsForbidden creates CloseSessionsForbidden with default headers values
func NewCloseSessionsForbidden() *CloseSessionsForbidden {

	return &CloseSessionsForbidden{}
}

// WithPayload adds the payload to the close sessions forbidden response
func (o *CloseSessionsForbidden) WithPayload(payload *models.Error) *CloseSessionsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the close sessions forbidden response
func (o *CloseSessionsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloseSessionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloseSessionsGoneCode is the HTTP code returned for type CloseSessionsGone
const CloseSessionsGoneCode int = 410

/*CloseSessionsGone That user (password and name) does not exist

swagger:response closeSessionsGone
*/
type CloseSessionsGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCloseSessionsGone creates CloseSessionsGone with default headers values
func NewCloseSessionsGone() *CloseSessionsGone {

	return &CloseSessionsGone{}
}

// WithPayload adds the payload to the close sessions gone response
func (o *CloseSessionsGone) WithPayload(payload *models.Error) *CloseSessionsGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the close sessions gone response
func (o *CloseSessionsGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloseSessionsGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CloseSessionsInternalServerErrorCode is the HTTP code returned for type CloseSessionsInternalServerError
const CloseSessionsInternalServerErrorCode int = 500

/*CloseSessionsInternalServerError Internal error

swagger:response closeSessionsInternalServerError
*/
type CloseSessionsInternalServerError struct {
}

// NewCloseSessionsInternalServerError creates CloseSessionsInternalServerError with default headers values
func NewCloseSessionsInternalServerError() *CloseSessionsInternalServerError {

	return &CloseSessionsInternalServerError{}
}

// WriteResponse to the client
func (o *CloseSessionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}