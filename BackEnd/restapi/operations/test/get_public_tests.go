// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPublicTestsHandlerFunc turns a function with the right signature into a get public tests handler
type GetPublicTestsHandlerFunc func(GetPublicTestsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublicTestsHandlerFunc) Handle(params GetPublicTestsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPublicTestsHandler interface for that can handle valid get public tests params
type GetPublicTestsHandler interface {
	Handle(GetPublicTestsParams, *models.User) middleware.Responder
}

// NewGetPublicTests creates a new http.Handler for the get public tests operation
func NewGetPublicTests(ctx *middleware.Context, handler GetPublicTestsHandler) *GetPublicTests {
	return &GetPublicTests{Context: ctx, Handler: handler}
}

/* GetPublicTests swagger:route GET /publicTests test getPublicTests

Returns all public tests

Returns all tests

*/
type GetPublicTests struct {
	Context *middleware.Context
	Handler GetPublicTestsHandler
}

func (o *GetPublicTests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublicTestsParams()
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
