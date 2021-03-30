// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetQuestionFromPublishedTestsHandlerFunc turns a function with the right signature into a get question from published tests handler
type GetQuestionFromPublishedTestsHandlerFunc func(GetQuestionFromPublishedTestsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionFromPublishedTestsHandlerFunc) Handle(params GetQuestionFromPublishedTestsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionFromPublishedTestsHandler interface for that can handle valid get question from published tests params
type GetQuestionFromPublishedTestsHandler interface {
	Handle(GetQuestionFromPublishedTestsParams, *models.User) middleware.Responder
}

// NewGetQuestionFromPublishedTests creates a new http.Handler for the get question from published tests operation
func NewGetQuestionFromPublishedTests(ctx *middleware.Context, handler GetQuestionFromPublishedTestsHandler) *GetQuestionFromPublishedTests {
	return &GetQuestionFromPublishedTests{Context: ctx, Handler: handler}
}

/* GetQuestionFromPublishedTests swagger:route GET /publishedTests/{testid}/questions/{questionid} publishedTest getQuestionFromPublishedTests

Returns a question from a published test

Returns a question from a published test

*/
type GetQuestionFromPublishedTests struct {
	Context *middleware.Context
	Handler GetQuestionFromPublishedTestsHandler
}

func (o *GetQuestionFromPublishedTests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetQuestionFromPublishedTestsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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
