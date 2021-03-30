// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddAmbirHandlerFunc turns a function with the right signature into a add ambir handler
type AddAmbirHandlerFunc func(AddAmbirParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddAmbirHandlerFunc) Handle(params AddAmbirParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddAmbirHandler interface for that can handle valid add ambir params
type AddAmbirHandler interface {
	Handle(AddAmbirParams, *models.User) middleware.Responder
}

// NewAddAmbir creates a new http.Handler for the add ambir operation
func NewAddAmbir(ctx *middleware.Context, handler AddAmbirHandler) *AddAmbir {
	return &AddAmbir{Context: ctx, Handler: handler}
}

/* AddAmbir swagger:route PUT /teams/{teamname}/admins/{username} team addAmbir

Adds user {username} to team {teamname} as an Admin

Adds user {username} to team {teamname} as an Admin.

*/
type AddAmbir struct {
	Context *middleware.Context
	Handler AddAmbirHandler
}

func (o *AddAmbir) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddAmbirParams()
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
