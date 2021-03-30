// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTestHandlerFunc turns a function with the right signature into a get test handler
type GetTestHandlerFunc func(GetTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTestHandlerFunc) Handle(params GetTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTestHandler interface for that can handle valid get test params
type GetTestHandler interface {
	Handle(GetTestParams, *models.User) middleware.Responder
}

// NewGetTest creates a new http.Handler for the get test operation
func NewGetTest(ctx *middleware.Context, handler GetTestHandler) *GetTest {
	return &GetTest{Context: ctx, Handler: handler}
}

/* GetTest swagger:route GET /tests/{testid} test getTest

Returns a test

Returns a test

*/
type GetTest struct {
	Context *middleware.Context
	Handler GetTestHandler
}

func (o *GetTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTestParams()
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
