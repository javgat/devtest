// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetQuestionFromUserOKCode is the HTTP code returned for type GetQuestionFromUserOK
const GetQuestionFromUserOKCode int = 200

/*GetQuestionFromUserOK question found

swagger:response getQuestionFromUserOK
*/
type GetQuestionFromUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.Question `json:"body,omitempty"`
}

// NewGetQuestionFromUserOK creates GetQuestionFromUserOK with default headers values
func NewGetQuestionFromUserOK() *GetQuestionFromUserOK {

	return &GetQuestionFromUserOK{}
}

// WithPayload adds the payload to the get question from user o k response
func (o *GetQuestionFromUserOK) WithPayload(payload *models.Question) *GetQuestionFromUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from user o k response
func (o *GetQuestionFromUserOK) SetPayload(payload *models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromUserBadRequestCode is the HTTP code returned for type GetQuestionFromUserBadRequest
const GetQuestionFromUserBadRequestCode int = 400

/*GetQuestionFromUserBadRequest Incorrect Request, or invalida data

swagger:response getQuestionFromUserBadRequest
*/
type GetQuestionFromUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromUserBadRequest creates GetQuestionFromUserBadRequest with default headers values
func NewGetQuestionFromUserBadRequest() *GetQuestionFromUserBadRequest {

	return &GetQuestionFromUserBadRequest{}
}

// WithPayload adds the payload to the get question from user bad request response
func (o *GetQuestionFromUserBadRequest) WithPayload(payload *models.Error) *GetQuestionFromUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from user bad request response
func (o *GetQuestionFromUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromUserForbiddenCode is the HTTP code returned for type GetQuestionFromUserForbidden
const GetQuestionFromUserForbiddenCode int = 403

/*GetQuestionFromUserForbidden Not authorized to this content

swagger:response getQuestionFromUserForbidden
*/
type GetQuestionFromUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromUserForbidden creates GetQuestionFromUserForbidden with default headers values
func NewGetQuestionFromUserForbidden() *GetQuestionFromUserForbidden {

	return &GetQuestionFromUserForbidden{}
}

// WithPayload adds the payload to the get question from user forbidden response
func (o *GetQuestionFromUserForbidden) WithPayload(payload *models.Error) *GetQuestionFromUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from user forbidden response
func (o *GetQuestionFromUserForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromUserGoneCode is the HTTP code returned for type GetQuestionFromUserGone
const GetQuestionFromUserGoneCode int = 410

/*GetQuestionFromUserGone That user (password and name) does not exist

swagger:response getQuestionFromUserGone
*/
type GetQuestionFromUserGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetQuestionFromUserGone creates GetQuestionFromUserGone with default headers values
func NewGetQuestionFromUserGone() *GetQuestionFromUserGone {

	return &GetQuestionFromUserGone{}
}

// WithPayload adds the payload to the get question from user gone response
func (o *GetQuestionFromUserGone) WithPayload(payload *models.Error) *GetQuestionFromUserGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get question from user gone response
func (o *GetQuestionFromUserGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionFromUserGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetQuestionFromUserInternalServerErrorCode is the HTTP code returned for type GetQuestionFromUserInternalServerError
const GetQuestionFromUserInternalServerErrorCode int = 500

/*GetQuestionFromUserInternalServerError Internal error

swagger:response getQuestionFromUserInternalServerError
*/
type GetQuestionFromUserInternalServerError struct {
}

// NewGetQuestionFromUserInternalServerError creates GetQuestionFromUserInternalServerError with default headers values
func NewGetQuestionFromUserInternalServerError() *GetQuestionFromUserInternalServerError {

	return &GetQuestionFromUserInternalServerError{}
}

// WriteResponse to the client
func (o *GetQuestionFromUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
