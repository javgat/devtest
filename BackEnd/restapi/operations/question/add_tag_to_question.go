// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// AddTagToQuestionHandlerFunc turns a function with the right signature into a add tag to question handler
type AddTagToQuestionHandlerFunc func(AddTagToQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTagToQuestionHandlerFunc) Handle(params AddTagToQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// AddTagToQuestionHandler interface for that can handle valid add tag to question params
type AddTagToQuestionHandler interface {
	Handle(AddTagToQuestionParams, *models.User) middleware.Responder
}

// NewAddTagToQuestion creates a new http.Handler for the add tag to question operation
func NewAddTagToQuestion(ctx *middleware.Context, handler AddTagToQuestionHandler) *AddTagToQuestion {
	return &AddTagToQuestion{Context: ctx, Handler: handler}
}

/* AddTagToQuestion swagger:route PUT /questions/{questionid}/tags/{tag} question addTagToQuestion

Adds a tag to a question

Adds a tag to a question

*/
type AddTagToQuestion struct {
	Context *middleware.Context
	Handler AddTagToQuestionHandler
}

func (o *AddTagToQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTagToQuestionParams()
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
