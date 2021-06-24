// Code generated by go-swagger; DO NOT EDIT.

package answer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPublishedPruebasFromQuestionTestHandlerFunc turns a function with the right signature into a get published pruebas from question test handler
type GetPublishedPruebasFromQuestionTestHandlerFunc func(GetPublishedPruebasFromQuestionTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublishedPruebasFromQuestionTestHandlerFunc) Handle(params GetPublishedPruebasFromQuestionTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPublishedPruebasFromQuestionTestHandler interface for that can handle valid get published pruebas from question test params
type GetPublishedPruebasFromQuestionTestHandler interface {
	Handle(GetPublishedPruebasFromQuestionTestParams, *models.User) middleware.Responder
}

// NewGetPublishedPruebasFromQuestionTest creates a new http.Handler for the get published pruebas from question test operation
func NewGetPublishedPruebasFromQuestionTest(ctx *middleware.Context, handler GetPublishedPruebasFromQuestionTestHandler) *GetPublishedPruebasFromQuestionTest {
	return &GetPublishedPruebasFromQuestionTest{Context: ctx, Handler: handler}
}

/* GetPublishedPruebasFromQuestionTest swagger:route GET /answers/{answerid}/questions/{questionid}/pruebas answer getPublishedPruebasFromQuestionTest

Returns all pruebas from a question. It will contain the correctness information related to the QuestionAnswer

Returns all pruebas from a question. It will contain the correctness information related to the QuestionAnswer

*/
type GetPublishedPruebasFromQuestionTest struct {
	Context *middleware.Context
	Handler GetPublishedPruebasFromQuestionTestHandler
}

func (o *GetPublishedPruebasFromQuestionTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublishedPruebasFromQuestionTestParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}