// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetCViewHandlerFunc turns a function with the right signature into a get c view handler
type GetCViewHandlerFunc func(GetCViewParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCViewHandlerFunc) Handle(params GetCViewParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetCViewHandler interface for that can handle valid get c view params
type GetCViewHandler interface {
	Handle(GetCViewParams, *models.User) middleware.Responder
}

// NewGetCView creates a new http.Handler for the get c view operation
func NewGetCView(ctx *middleware.Context, handler GetCViewHandler) *GetCView {
	return &GetCView{Context: ctx, Handler: handler}
}

/* GetCView swagger:route GET /customizedViews/{rolBase} configuration getCView

Returns a CustomizedView

Returns a CustomizedView

*/
type GetCView struct {
	Context *middleware.Context
	Handler GetCViewHandler
}

func (o *GetCView) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCViewParams()
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