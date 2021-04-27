// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetFavoriteQuestionOKCode is the HTTP code returned for type GetFavoriteQuestionOK
const GetFavoriteQuestionOKCode int = 200

/*GetFavoriteQuestionOK question found

swagger:response getFavoriteQuestionOK
*/
type GetFavoriteQuestionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Question `json:"body,omitempty"`
}

// NewGetFavoriteQuestionOK creates GetFavoriteQuestionOK with default headers values
func NewGetFavoriteQuestionOK() *GetFavoriteQuestionOK {

	return &GetFavoriteQuestionOK{}
}

// WithPayload adds the payload to the get favorite question o k response
func (o *GetFavoriteQuestionOK) WithPayload(payload *models.Question) *GetFavoriteQuestionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite question o k response
func (o *GetFavoriteQuestionOK) SetPayload(payload *models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteQuestionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteQuestionBadRequestCode is the HTTP code returned for type GetFavoriteQuestionBadRequest
const GetFavoriteQuestionBadRequestCode int = 400

/*GetFavoriteQuestionBadRequest Incorrect Request, or invalida data

swagger:response getFavoriteQuestionBadRequest
*/
type GetFavoriteQuestionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteQuestionBadRequest creates GetFavoriteQuestionBadRequest with default headers values
func NewGetFavoriteQuestionBadRequest() *GetFavoriteQuestionBadRequest {

	return &GetFavoriteQuestionBadRequest{}
}

// WithPayload adds the payload to the get favorite question bad request response
func (o *GetFavoriteQuestionBadRequest) WithPayload(payload *models.Error) *GetFavoriteQuestionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite question bad request response
func (o *GetFavoriteQuestionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteQuestionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteQuestionForbiddenCode is the HTTP code returned for type GetFavoriteQuestionForbidden
const GetFavoriteQuestionForbiddenCode int = 403

/*GetFavoriteQuestionForbidden Not authorized to this content

swagger:response getFavoriteQuestionForbidden
*/
type GetFavoriteQuestionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteQuestionForbidden creates GetFavoriteQuestionForbidden with default headers values
func NewGetFavoriteQuestionForbidden() *GetFavoriteQuestionForbidden {

	return &GetFavoriteQuestionForbidden{}
}

// WithPayload adds the payload to the get favorite question forbidden response
func (o *GetFavoriteQuestionForbidden) WithPayload(payload *models.Error) *GetFavoriteQuestionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite question forbidden response
func (o *GetFavoriteQuestionForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteQuestionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteQuestionGoneCode is the HTTP code returned for type GetFavoriteQuestionGone
const GetFavoriteQuestionGoneCode int = 410

/*GetFavoriteQuestionGone That resource does not exist

swagger:response getFavoriteQuestionGone
*/
type GetFavoriteQuestionGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFavoriteQuestionGone creates GetFavoriteQuestionGone with default headers values
func NewGetFavoriteQuestionGone() *GetFavoriteQuestionGone {

	return &GetFavoriteQuestionGone{}
}

// WithPayload adds the payload to the get favorite question gone response
func (o *GetFavoriteQuestionGone) WithPayload(payload *models.Error) *GetFavoriteQuestionGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get favorite question gone response
func (o *GetFavoriteQuestionGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFavoriteQuestionGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFavoriteQuestionInternalServerErrorCode is the HTTP code returned for type GetFavoriteQuestionInternalServerError
const GetFavoriteQuestionInternalServerErrorCode int = 500

/*GetFavoriteQuestionInternalServerError Internal error

swagger:response getFavoriteQuestionInternalServerError
*/
type GetFavoriteQuestionInternalServerError struct {
}

// NewGetFavoriteQuestionInternalServerError creates GetFavoriteQuestionInternalServerError with default headers values
func NewGetFavoriteQuestionInternalServerError() *GetFavoriteQuestionInternalServerError {

	return &GetFavoriteQuestionInternalServerError{}
}

// WriteResponse to the client
func (o *GetFavoriteQuestionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}