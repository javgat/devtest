// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPublicPublishedTestHandlerFunc turns a function with the right signature into a get public published test handler
type GetPublicPublishedTestHandlerFunc func(GetPublicPublishedTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublicPublishedTestHandlerFunc) Handle(params GetPublicPublishedTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPublicPublishedTestHandler interface for that can handle valid get public published test params
type GetPublicPublishedTestHandler interface {
	Handle(GetPublicPublishedTestParams, *models.User) middleware.Responder
}

// NewGetPublicPublishedTest creates a new http.Handler for the get public published test operation
func NewGetPublicPublishedTest(ctx *middleware.Context, handler GetPublicPublishedTestHandler) *GetPublicPublishedTest {
	return &GetPublicPublishedTest{Context: ctx, Handler: handler}
}

/* GetPublicPublishedTest swagger:route GET /publicPublishedTests/{testid} publishedTest getPublicPublishedTest

Returns a public publishedTest

Returns a public publishedTest

*/
type GetPublicPublishedTest struct {
	Context *middleware.Context
	Handler GetPublicPublishedTestHandler
}

func (o *GetPublicPublishedTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublicPublishedTestParams()
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