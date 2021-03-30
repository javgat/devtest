// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// RemoveQuestionFromTestHandlerFunc turns a function with the right signature into a remove question from test handler
type RemoveQuestionFromTestHandlerFunc func(RemoveQuestionFromTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveQuestionFromTestHandlerFunc) Handle(params RemoveQuestionFromTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// RemoveQuestionFromTestHandler interface for that can handle valid remove question from test params
type RemoveQuestionFromTestHandler interface {
	Handle(RemoveQuestionFromTestParams, *models.User) middleware.Responder
}

// NewRemoveQuestionFromTest creates a new http.Handler for the remove question from test operation
func NewRemoveQuestionFromTest(ctx *middleware.Context, handler RemoveQuestionFromTestHandler) *RemoveQuestionFromTest {
	return &RemoveQuestionFromTest{Context: ctx, Handler: handler}
}

/* RemoveQuestionFromTest swagger:route DELETE /tests/{testid}/questions/{questionid} test removeQuestionFromTest

Removes an existing question from a test

Removes an existing question from a test

*/
type RemoveQuestionFromTest struct {
	Context *middleware.Context
	Handler RemoveQuestionFromTestHandler
}

func (o *RemoveQuestionFromTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveQuestionFromTestParams()
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
