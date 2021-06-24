// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetVisiblePublishedPruebasFromQuestionTestOKCode is the HTTP code returned for type GetVisiblePublishedPruebasFromQuestionTestOK
const GetVisiblePublishedPruebasFromQuestionTestOKCode int = 200

/*GetVisiblePublishedPruebasFromQuestionTestOK pruebas visibles found

swagger:response getVisiblePublishedPruebasFromQuestionTestOK
*/
type GetVisiblePublishedPruebasFromQuestionTestOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Prueba `json:"body,omitempty"`
}

// NewGetVisiblePublishedPruebasFromQuestionTestOK creates GetVisiblePublishedPruebasFromQuestionTestOK with default headers values
func NewGetVisiblePublishedPruebasFromQuestionTestOK() *GetVisiblePublishedPruebasFromQuestionTestOK {

	return &GetVisiblePublishedPruebasFromQuestionTestOK{}
}

// WithPayload adds the payload to the get visible published pruebas from question test o k response
func (o *GetVisiblePublishedPruebasFromQuestionTestOK) WithPayload(payload []*models.Prueba) *GetVisiblePublishedPruebasFromQuestionTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get visible published pruebas from question test o k response
func (o *GetVisiblePublishedPruebasFromQuestionTestOK) SetPayload(payload []*models.Prueba) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVisiblePublishedPruebasFromQuestionTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Prueba, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVisiblePublishedPruebasFromQuestionTestBadRequestCode is the HTTP code returned for type GetVisiblePublishedPruebasFromQuestionTestBadRequest
const GetVisiblePublishedPruebasFromQuestionTestBadRequestCode int = 400

/*GetVisiblePublishedPruebasFromQuestionTestBadRequest Incorrect Request, or invalida data

swagger:response getVisiblePublishedPruebasFromQuestionTestBadRequest
*/
type GetVisiblePublishedPruebasFromQuestionTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVisiblePublishedPruebasFromQuestionTestBadRequest creates GetVisiblePublishedPruebasFromQuestionTestBadRequest with default headers values
func NewGetVisiblePublishedPruebasFromQuestionTestBadRequest() *GetVisiblePublishedPruebasFromQuestionTestBadRequest {

	return &GetVisiblePublishedPruebasFromQuestionTestBadRequest{}
}

// WithPayload adds the payload to the get visible published pruebas from question test bad request response
func (o *GetVisiblePublishedPruebasFromQuestionTestBadRequest) WithPayload(payload *models.Error) *GetVisiblePublishedPruebasFromQuestionTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get visible published pruebas from question test bad request response
func (o *GetVisiblePublishedPruebasFromQuestionTestBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVisiblePublishedPruebasFromQuestionTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVisiblePublishedPruebasFromQuestionTestForbiddenCode is the HTTP code returned for type GetVisiblePublishedPruebasFromQuestionTestForbidden
const GetVisiblePublishedPruebasFromQuestionTestForbiddenCode int = 403

/*GetVisiblePublishedPruebasFromQuestionTestForbidden Not authorized to this content

swagger:response getVisiblePublishedPruebasFromQuestionTestForbidden
*/
type GetVisiblePublishedPruebasFromQuestionTestForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVisiblePublishedPruebasFromQuestionTestForbidden creates GetVisiblePublishedPruebasFromQuestionTestForbidden with default headers values
func NewGetVisiblePublishedPruebasFromQuestionTestForbidden() *GetVisiblePublishedPruebasFromQuestionTestForbidden {

	return &GetVisiblePublishedPruebasFromQuestionTestForbidden{}
}

// WithPayload adds the payload to the get visible published pruebas from question test forbidden response
func (o *GetVisiblePublishedPruebasFromQuestionTestForbidden) WithPayload(payload *models.Error) *GetVisiblePublishedPruebasFromQuestionTestForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get visible published pruebas from question test forbidden response
func (o *GetVisiblePublishedPruebasFromQuestionTestForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVisiblePublishedPruebasFromQuestionTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVisiblePublishedPruebasFromQuestionTestGoneCode is the HTTP code returned for type GetVisiblePublishedPruebasFromQuestionTestGone
const GetVisiblePublishedPruebasFromQuestionTestGoneCode int = 410

/*GetVisiblePublishedPruebasFromQuestionTestGone That resource does not exist

swagger:response getVisiblePublishedPruebasFromQuestionTestGone
*/
type GetVisiblePublishedPruebasFromQuestionTestGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVisiblePublishedPruebasFromQuestionTestGone creates GetVisiblePublishedPruebasFromQuestionTestGone with default headers values
func NewGetVisiblePublishedPruebasFromQuestionTestGone() *GetVisiblePublishedPruebasFromQuestionTestGone {

	return &GetVisiblePublishedPruebasFromQuestionTestGone{}
}

// WithPayload adds the payload to the get visible published pruebas from question test gone response
func (o *GetVisiblePublishedPruebasFromQuestionTestGone) WithPayload(payload *models.Error) *GetVisiblePublishedPruebasFromQuestionTestGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get visible published pruebas from question test gone response
func (o *GetVisiblePublishedPruebasFromQuestionTestGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVisiblePublishedPruebasFromQuestionTestGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVisiblePublishedPruebasFromQuestionTestInternalServerErrorCode is the HTTP code returned for type GetVisiblePublishedPruebasFromQuestionTestInternalServerError
const GetVisiblePublishedPruebasFromQuestionTestInternalServerErrorCode int = 500

/*GetVisiblePublishedPruebasFromQuestionTestInternalServerError Internal error

swagger:response getVisiblePublishedPruebasFromQuestionTestInternalServerError
*/
type GetVisiblePublishedPruebasFromQuestionTestInternalServerError struct {
}

// NewGetVisiblePublishedPruebasFromQuestionTestInternalServerError creates GetVisiblePublishedPruebasFromQuestionTestInternalServerError with default headers values
func NewGetVisiblePublishedPruebasFromQuestionTestInternalServerError() *GetVisiblePublishedPruebasFromQuestionTestInternalServerError {

	return &GetVisiblePublishedPruebasFromQuestionTestInternalServerError{}
}

// WriteResponse to the client
func (o *GetVisiblePublishedPruebasFromQuestionTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}