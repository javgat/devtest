// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetPublicPublishedTestsFromUserOKCode is the HTTP code returned for type GetPublicPublishedTestsFromUserOK
const GetPublicPublishedTestsFromUserOKCode int = 200

/*GetPublicPublishedTestsFromUserOK publishedTests found

swagger:response getPublicPublishedTestsFromUserOK
*/
type GetPublicPublishedTestsFromUserOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Test `json:"body,omitempty"`
}

// NewGetPublicPublishedTestsFromUserOK creates GetPublicPublishedTestsFromUserOK with default headers values
func NewGetPublicPublishedTestsFromUserOK() *GetPublicPublishedTestsFromUserOK {

	return &GetPublicPublishedTestsFromUserOK{}
}

// WithPayload adds the payload to the get public published tests from user o k response
func (o *GetPublicPublishedTestsFromUserOK) WithPayload(payload []*models.Test) *GetPublicPublishedTestsFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public published tests from user o k response
func (o *GetPublicPublishedTestsFromUserOK) SetPayload(payload []*models.Test) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicPublishedTestsFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetPublicPublishedTestsFromUserBadRequestCode is the HTTP code returned for type GetPublicPublishedTestsFromUserBadRequest
const GetPublicPublishedTestsFromUserBadRequestCode int = 400

/*GetPublicPublishedTestsFromUserBadRequest Incorrect Request, or invalida data

swagger:response getPublicPublishedTestsFromUserBadRequest
*/
type GetPublicPublishedTestsFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicPublishedTestsFromUserBadRequest creates GetPublicPublishedTestsFromUserBadRequest with default headers values
func NewGetPublicPublishedTestsFromUserBadRequest() *GetPublicPublishedTestsFromUserBadRequest {

	return &GetPublicPublishedTestsFromUserBadRequest{}
}

// WithPayload adds the payload to the get public published tests from user bad request response
func (o *GetPublicPublishedTestsFromUserBadRequest) WithPayload(payload *models.Error) *GetPublicPublishedTestsFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public published tests from user bad request response
func (o *GetPublicPublishedTestsFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicPublishedTestsFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicPublishedTestsFromUserForbiddenCode is the HTTP code returned for type GetPublicPublishedTestsFromUserForbidden
const GetPublicPublishedTestsFromUserForbiddenCode int = 403

/*GetPublicPublishedTestsFromUserForbidden Not authorized to this content

swagger:response getPublicPublishedTestsFromUserForbidden
*/
type GetPublicPublishedTestsFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicPublishedTestsFromUserForbidden creates GetPublicPublishedTestsFromUserForbidden with default headers values
func NewGetPublicPublishedTestsFromUserForbidden() *GetPublicPublishedTestsFromUserForbidden {

	return &GetPublicPublishedTestsFromUserForbidden{}
}

// WithPayload adds the payload to the get public published tests from user forbidden response
func (o *GetPublicPublishedTestsFromUserForbidden) WithPayload(payload *models.Error) *GetPublicPublishedTestsFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public published tests from user forbidden response
func (o *GetPublicPublishedTestsFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicPublishedTestsFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicPublishedTestsFromUserGoneCode is the HTTP code returned for type GetPublicPublishedTestsFromUserGone
const GetPublicPublishedTestsFromUserGoneCode int = 410

/*GetPublicPublishedTestsFromUserGone That resource does not exist

swagger:response getPublicPublishedTestsFromUserGone
*/
type GetPublicPublishedTestsFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPublicPublishedTestsFromUserGone creates GetPublicPublishedTestsFromUserGone with default headers values
func NewGetPublicPublishedTestsFromUserGone() *GetPublicPublishedTestsFromUserGone {

	return &GetPublicPublishedTestsFromUserGone{}
}

// WithPayload adds the payload to the get public published tests from user gone response
func (o *GetPublicPublishedTestsFromUserGone) WithPayload(payload *models.Error) *GetPublicPublishedTestsFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get public published tests from user gone response
func (o *GetPublicPublishedTestsFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPublicPublishedTestsFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPublicPublishedTestsFromUserInternalServerErrorCode is the HTTP code returned for type GetPublicPublishedTestsFromUserInternalServerError
const GetPublicPublishedTestsFromUserInternalServerErrorCode int = 500

/*GetPublicPublishedTestsFromUserInternalServerError Internal error

swagger:response getPublicPublishedTestsFromUserInternalServerError
*/
type GetPublicPublishedTestsFromUserInternalServerError struct {
}

// NewGetPublicPublishedTestsFromUserInternalServerError creates GetPublicPublishedTestsFromUserInternalServerError with default headers values
func NewGetPublicPublishedTestsFromUserInternalServerError() *GetPublicPublishedTestsFromUserInternalServerError {

	return &GetPublicPublishedTestsFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetPublicPublishedTestsFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}