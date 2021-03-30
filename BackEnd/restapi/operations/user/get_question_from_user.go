// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetQuestionFromUserHandlerFunc turns a function with the right signature into a get question from user handler
type GetQuestionFromUserHandlerFunc func(GetQuestionFromUserParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionFromUserHandlerFunc) Handle(params GetQuestionFromUserParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionFromUserHandler interface for that can handle valid get question from user params
type GetQuestionFromUserHandler interface {
	Handle(GetQuestionFromUserParams, *models.User) middleware.Responder
}

// NewGetQuestionFromUser creates a new http.Handler for the get question from user operation
func NewGetQuestionFromUser(ctx *middleware.Context, handler GetQuestionFromUserHandler) *GetQuestionFromUser {
	return &GetQuestionFromUser{Context: ctx, Handler: handler}
}

/* GetQuestionFromUser swagger:route GET /users/{username}/questions/{questionid} user getQuestionFromUser

Returns a question of a user

Returns a question of a user

*/
type GetQuestionFromUser struct {
	Context *middleware.Context
	Handler GetQuestionFromUserHandler
}

func (o *GetQuestionFromUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetQuestionFromUserParams()
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
