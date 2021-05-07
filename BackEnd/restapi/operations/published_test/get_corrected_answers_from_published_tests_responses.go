// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"uva-devtest/models"
)

// GetCorrectedAnswersFromPublishedTestsOKCode is the HTTP code returned for type GetCorrectedAnswersFromPublishedTestsOK
const GetCorrectedAnswersFromPublishedTestsOKCode int = 200

/*GetCorrectedAnswersFromPublishedTestsOK Answers found

swagger:response getCorrectedAnswersFromPublishedTestsOK
*/
type GetCorrectedAnswersFromPublishedTestsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Answer `json:"body,omitempty"`
}

// NewGetCorrectedAnswersFromPublishedTestsOK creates GetCorrectedAnswersFromPublishedTestsOK with default headers values
func NewGetCorrectedAnswersFromPublishedTestsOK() *GetCorrectedAnswersFromPublishedTestsOK {

	return &GetCorrectedAnswersFromPublishedTestsOK{}
}

// WithPayload adds the payload to the get corrected answers from published tests o k response
func (o *GetCorrectedAnswersFromPublishedTestsOK) WithPayload(payload []*models.Answer) *GetCorrectedAnswersFromPublishedTestsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corrected answers from published tests o k response
func (o *GetCorrectedAnswersFromPublishedTestsOK) SetPayload(payload []*models.Answer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorrectedAnswersFromPublishedTestsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Answer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCorrectedAnswersFromPublishedTestsBadRequestCode is the HTTP code returned for type GetCorrectedAnswersFromPublishedTestsBadRequest
const GetCorrectedAnswersFromPublishedTestsBadRequestCode int = 400

/*GetCorrectedAnswersFromPublishedTestsBadRequest Incorrect Request, or invalida data

swagger:response getCorrectedAnswersFromPublishedTestsBadRequest
*/
type GetCorrectedAnswersFromPublishedTestsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCorrectedAnswersFromPublishedTestsBadRequest creates GetCorrectedAnswersFromPublishedTestsBadRequest with default headers values
func NewGetCorrectedAnswersFromPublishedTestsBadRequest() *GetCorrectedAnswersFromPublishedTestsBadRequest {

	return &GetCorrectedAnswersFromPublishedTestsBadRequest{}
}

// WithPayload adds the payload to the get corrected answers from published tests bad request response
func (o *GetCorrectedAnswersFromPublishedTestsBadRequest) WithPayload(payload *models.Error) *GetCorrectedAnswersFromPublishedTestsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corrected answers from published tests bad request response
func (o *GetCorrectedAnswersFromPublishedTestsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorrectedAnswersFromPublishedTestsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCorrectedAnswersFromPublishedTestsForbiddenCode is the HTTP code returned for type GetCorrectedAnswersFromPublishedTestsForbidden
const GetCorrectedAnswersFromPublishedTestsForbiddenCode int = 403

/*GetCorrectedAnswersFromPublishedTestsForbidden Not authorized to this content

swagger:response getCorrectedAnswersFromPublishedTestsForbidden
*/
type GetCorrectedAnswersFromPublishedTestsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCorrectedAnswersFromPublishedTestsForbidden creates GetCorrectedAnswersFromPublishedTestsForbidden with default headers values
func NewGetCorrectedAnswersFromPublishedTestsForbidden() *GetCorrectedAnswersFromPublishedTestsForbidden {

	return &GetCorrectedAnswersFromPublishedTestsForbidden{}
}

// WithPayload adds the payload to the get corrected answers from published tests forbidden response
func (o *GetCorrectedAnswersFromPublishedTestsForbidden) WithPayload(payload *models.Error) *GetCorrectedAnswersFromPublishedTestsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corrected answers from published tests forbidden response
func (o *GetCorrectedAnswersFromPublishedTestsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorrectedAnswersFromPublishedTestsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCorrectedAnswersFromPublishedTestsGoneCode is the HTTP code returned for type GetCorrectedAnswersFromPublishedTestsGone
const GetCorrectedAnswersFromPublishedTestsGoneCode int = 410

/*GetCorrectedAnswersFromPublishedTestsGone That resource does not exist

swagger:response getCorrectedAnswersFromPublishedTestsGone
*/
type GetCorrectedAnswersFromPublishedTestsGone struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCorrectedAnswersFromPublishedTestsGone creates GetCorrectedAnswersFromPublishedTestsGone with default headers values
func NewGetCorrectedAnswersFromPublishedTestsGone() *GetCorrectedAnswersFromPublishedTestsGone {

	return &GetCorrectedAnswersFromPublishedTestsGone{}
}

// WithPayload adds the payload to the get corrected answers from published tests gone response
func (o *GetCorrectedAnswersFromPublishedTestsGone) WithPayload(payload *models.Error) *GetCorrectedAnswersFromPublishedTestsGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get corrected answers from published tests gone response
func (o *GetCorrectedAnswersFromPublishedTestsGone) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCorrectedAnswersFromPublishedTestsGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCorrectedAnswersFromPublishedTestsInternalServerErrorCode is the HTTP code returned for type GetCorrectedAnswersFromPublishedTestsInternalServerError
const GetCorrectedAnswersFromPublishedTestsInternalServerErrorCode int = 500

/*GetCorrectedAnswersFromPublishedTestsInternalServerError Internal error

swagger:response getCorrectedAnswersFromPublishedTestsInternalServerError
*/
type GetCorrectedAnswersFromPublishedTestsInternalServerError struct {
}

// NewGetCorrectedAnswersFromPublishedTestsInternalServerError creates GetCorrectedAnswersFromPublishedTestsInternalServerError with default headers values
func NewGetCorrectedAnswersFromPublishedTestsInternalServerError() *GetCorrectedAnswersFromPublishedTestsInternalServerError {

	return &GetCorrectedAnswersFromPublishedTestsInternalServerError{}
}

// WriteResponse to the client
func (o *GetCorrectedAnswersFromPublishedTestsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}