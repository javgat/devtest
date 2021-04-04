// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddTeamToQuestionHandlerFunc turns a function with the right signature into a add team to question handler
type AddTeamToQuestionHandlerFunc func(AddTeamToQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTeamToQuestionHandlerFunc) Handle(params AddTeamToQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddTeamToQuestionHandler interface for that can handle valid add team to question params
type AddTeamToQuestionHandler interface {
	Handle(AddTeamToQuestionParams, *models.User) middleware.Responder
}

// NewAddTeamToQuestion creates a new http.Handler for the add team to question operation
func NewAddTeamToQuestion(ctx *middleware.Context, handler AddTeamToQuestionHandler) *AddTeamToQuestion {
	return &AddTeamToQuestion{Context: ctx, Handler: handler}
}

/* AddTeamToQuestion swagger:route PUT /questions/{questionid}/teams/{teamname} question addTeamToQuestion

Adds a team to administer a question

Adds a team to administer a question

*/
type AddTeamToQuestion struct {
	Context *middleware.Context
	Handler AddTeamToQuestionHandler
}

func (o *AddTeamToQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTeamToQuestionParams()
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