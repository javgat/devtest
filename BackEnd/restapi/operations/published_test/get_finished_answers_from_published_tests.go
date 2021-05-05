// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetFinishedAnswersFromPublishedTestsHandlerFunc turns a function with the right signature into a get finished answers from published tests handler
type GetFinishedAnswersFromPublishedTestsHandlerFunc func(GetFinishedAnswersFromPublishedTestsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFinishedAnswersFromPublishedTestsHandlerFunc) Handle(params GetFinishedAnswersFromPublishedTestsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetFinishedAnswersFromPublishedTestsHandler interface for that can handle valid get finished answers from published tests params
type GetFinishedAnswersFromPublishedTestsHandler interface {
	Handle(GetFinishedAnswersFromPublishedTestsParams, *models.User) middleware.Responder
}

// NewGetFinishedAnswersFromPublishedTests creates a new http.Handler for the get finished answers from published tests operation
func NewGetFinishedAnswersFromPublishedTests(ctx *middleware.Context, handler GetFinishedAnswersFromPublishedTestsHandler) *GetFinishedAnswersFromPublishedTests {
	return &GetFinishedAnswersFromPublishedTests{Context: ctx, Handler: handler}
}

/* GetFinishedAnswersFromPublishedTests swagger:route GET /publishedTests/{testid}/finishedAnswers publishedTest getFinishedAnswersFromPublishedTests

Returns all answers from a published test that are finished

Returns all answers from a published test that are finished

*/
type GetFinishedAnswersFromPublishedTests struct {
	Context *middleware.Context
	Handler GetFinishedAnswersFromPublishedTestsHandler
}

func (o *GetFinishedAnswersFromPublishedTests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFinishedAnswersFromPublishedTestsParams()
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
