// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// RemoveTeamToPublishedTestOKCode is the HTTP code returned for type RemoveTeamToPublishedTestOK
const RemoveTeamToPublishedTestOKCode int = 200

/*RemoveTeamToPublishedTestOK team removed

swagger:response removeTeamToPublishedTestOK
*/
type RemoveTeamToPublishedTestOK struct {
}

// NewRemoveTeamToPublishedTestOK creates RemoveTeamToPublishedTestOK with default headers values
func NewRemoveTeamToPublishedTestOK() *RemoveTeamToPublishedTestOK {

	return &RemoveTeamToPublishedTestOK{}
}

// WriteResponse to the client
func (o *RemoveTeamToPublishedTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RemoveTeamToPublishedTestBadRequestCode is the HTTP code returned for type RemoveTeamToPublishedTestBadRequest
const RemoveTeamToPublishedTestBadRequestCode int = 400

/*RemoveTeamToPublishedTestBadRequest Incorrect Request, or invalida data

swagger:response removeTeamToPublishedTestBadRequest
*/
type RemoveTeamToPublishedTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToPublishedTestBadRequest creates RemoveTeamToPublishedTestBadRequest with default headers values
func NewRemoveTeamToPublishedTestBadRequest() *RemoveTeamToPublishedTestBadRequest {

	return &RemoveTeamToPublishedTestBadRequest{}
}

// WithPayload adds the payload to the remove team to published test bad request response
func (o *RemoveTeamToPublishedTestBadRequest) WithPayload(payload *models.Error) *RemoveTeamToPublishedTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to published test bad request response
func (o *RemoveTeamToPublishedTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToPublishedTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToPublishedTestForbiddenCode is the HTTP code returned for type RemoveTeamToPublishedTestForbidden
const RemoveTeamToPublishedTestForbiddenCode int = 403

/*RemoveTeamToPublishedTestForbidden Not authorized to this content

swagger:response removeTeamToPublishedTestForbidden
*/
type RemoveTeamToPublishedTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToPublishedTestForbidden creates RemoveTeamToPublishedTestForbidden with default headers values
func NewRemoveTeamToPublishedTestForbidden() *RemoveTeamToPublishedTestForbidden {

	return &RemoveTeamToPublishedTestForbidden{}
}

// WithPayload adds the payload to the remove team to published test forbidden response
func (o *RemoveTeamToPublishedTestForbidden) WithPayload(payload *models.Error) *RemoveTeamToPublishedTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to published test forbidden response
func (o *RemoveTeamToPublishedTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToPublishedTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToPublishedTestGoneCode is the HTTP code returned for type RemoveTeamToPublishedTestGone
const RemoveTeamToPublishedTestGoneCode int = 410

/*RemoveTeamToPublishedTestGone That user (password and name) does not exist

swagger:response removeTeamToPublishedTestGone
*/
type RemoveTeamToPublishedTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoveTeamToPublishedTestGone creates RemoveTeamToPublishedTestGone with default headers values
func NewRemoveTeamToPublishedTestGone() *RemoveTeamToPublishedTestGone {

	return &RemoveTeamToPublishedTestGone{}
}

// WithPayload adds the payload to the remove team to published test gone response
func (o *RemoveTeamToPublishedTestGone) WithPayload(payload *models.Error) *RemoveTeamToPublishedTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove team to published test gone response
func (o *RemoveTeamToPublishedTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveTeamToPublishedTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveTeamToPublishedTestInternalServerErrorCode is the HTTP code returned for type RemoveTeamToPublishedTestInternalServerError
const RemoveTeamToPublishedTestInternalServerErrorCode int = 500

/*RemoveTeamToPublishedTestInternalServerError Internal error

swagger:response removeTeamToPublishedTestInternalServerError
*/
type RemoveTeamToPublishedTestInternalServerError struct {
}

// NewRemoveTeamToPublishedTestInternalServerError creates RemoveTeamToPublishedTestInternalServerError with default headers values
func NewRemoveTeamToPublishedTestInternalServerError() *RemoveTeamToPublishedTestInternalServerError {

	return &RemoveTeamToPublishedTestInternalServerError{}
}

// WriteResponse to the client
func (o *RemoveTeamToPublishedTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}