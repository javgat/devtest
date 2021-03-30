// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetAnswerOKCode is the HTTP code returned for type GetAnswerOK
const GetAnswerOKCode int = 200

/*GetAnswerOK Answers found

swagger:response getAnswerOK
*/
type GetAnswerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Answer `json:"body,omitempty"`
}

// NewGetAnswerOK creates GetAnswerOK with default headers values
func NewGetAnswerOK() *GetAnswerOK {

	return &GetAnswerOK{}
}

// WithPayload adds the payload to the get answer o k response
func (o *GetAnswerOK) WithPayload(payload *models.Answer) *GetAnswerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get answer o k response
func (o *GetAnswerOK) SetPayload(payload *models.Answer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnswerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnswerBadRequestCode is the HTTP code returned for type GetAnswerBadRequest
const GetAnswerBadRequestCode int = 400

/*GetAnswerBadRequest Incorrect Request, or invalida data

swagger:response getAnswerBadRequest
*/
type GetAnswerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAnswerBadRequest creates GetAnswerBadRequest with default headers values
func NewGetAnswerBadRequest() *GetAnswerBadRequest {

	return &GetAnswerBadRequest{}
}

// WithPayload adds the payload to the get answer bad request response
func (o *GetAnswerBadRequest) WithPayload(payload *models.Error) *GetAnswerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get answer bad request response
func (o *GetAnswerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnswerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnswerForbiddenCode is the HTTP code returned for type GetAnswerForbidden
const GetAnswerForbiddenCode int = 403

/*GetAnswerForbidden Not authorized to this content

swagger:response getAnswerForbidden
*/
type GetAnswerForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAnswerForbidden creates GetAnswerForbidden with default headers values
func NewGetAnswerForbidden() *GetAnswerForbidden {

	return &GetAnswerForbidden{}
}

// WithPayload adds the payload to the get answer forbidden response
func (o *GetAnswerForbidden) WithPayload(payload *models.Error) *GetAnswerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get answer forbidden response
func (o *GetAnswerForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnswerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnswerGoneCode is the HTTP code returned for type GetAnswerGone
const GetAnswerGoneCode int = 410

/*GetAnswerGone That user (password and name) does not exist

swagger:response getAnswerGone
*/
type GetAnswerGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAnswerGone creates GetAnswerGone with default headers values
func NewGetAnswerGone() *GetAnswerGone {

	return &GetAnswerGone{}
}

// WithPayload adds the payload to the get answer gone response
func (o *GetAnswerGone) WithPayload(payload *models.Error) *GetAnswerGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get answer gone response
func (o *GetAnswerGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAnswerGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAnswerInternalServerErrorCode is the HTTP code returned for type GetAnswerInternalServerError
const GetAnswerInternalServerErrorCode int = 500

/*GetAnswerInternalServerError Internal error

swagger:response getAnswerInternalServerError
*/
type GetAnswerInternalServerError struct {
}

// NewGetAnswerInternalServerError creates GetAnswerInternalServerError with default headers values
func NewGetAnswerInternalServerError() *GetAnswerInternalServerError {

	return &GetAnswerInternalServerError{}
}

// WriteResponse to the client
func (o *GetAnswerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
