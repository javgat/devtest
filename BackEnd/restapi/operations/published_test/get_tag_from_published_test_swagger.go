// Code generated by go-swagger; DO NOT EDIT.

package published_test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetTagFromPublishedTestHandlerFunc turns a function with the right signature into a get tag from published test handler
type GetTagFromPublishedTestHandlerFunc func(GetTagFromPublishedTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTagFromPublishedTestHandlerFunc) Handle(params GetTagFromPublishedTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTagFromPublishedTestHandler interface for that can handle valid get tag from published test params
type GetTagFromPublishedTestHandler interface {
	Handle(GetTagFromPublishedTestParams, *models.User) middleware.Responder
}

// NewGetTagFromPublishedTest creates a new http.Handler for the get tag from published test operation
func NewGetTagFromPublishedTest(ctx *middleware.Context, handler GetTagFromPublishedTestHandler) *GetTagFromPublishedTest {
	return &GetTagFromPublishedTest{Context: ctx, Handler: handler}
}

/* GetTagFromPublishedTest swagger:route GET /publishedTests/{testid}/tags/{tag} publishedTest getTagFromPublishedTest

Returns a tag from a published test.

Returns a tag from a published test.

*/
type GetTagFromPublishedTest struct {
	Context *middleware.Context
	Handler GetTagFromPublishedTestHandler
}

func (o *GetTagFromPublishedTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTagFromPublishedTestParams()
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