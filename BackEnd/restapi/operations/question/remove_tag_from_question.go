// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// RemoveTagFromQuestionHandlerFunc turns a function with the right signature into a remove tag from question handler
type RemoveTagFromQuestionHandlerFunc func(RemoveTagFromQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveTagFromQuestionHandlerFunc) Handle(params RemoveTagFromQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// RemoveTagFromQuestionHandler interface for that can handle valid remove tag from question params
type RemoveTagFromQuestionHandler interface {
	Handle(RemoveTagFromQuestionParams, *models.User) middleware.Responder
}

// NewRemoveTagFromQuestion creates a new http.Handler for the remove tag from question operation
func NewRemoveTagFromQuestion(ctx *middleware.Context, handler RemoveTagFromQuestionHandler) *RemoveTagFromQuestion {
	return &RemoveTagFromQuestion{Context: ctx, Handler: handler}
}

/* RemoveTagFromQuestion swagger:route DELETE /questions/{questionid}/tags/{tag} question removeTagFromQuestion

Removes a tag from a question

Removes a tag from a question

*/
type RemoveTagFromQuestion struct {
	Context *middleware.Context
	Handler RemoveTagFromQuestionHandler
}

func (o *RemoveTagFromQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRemoveTagFromQuestionParams()
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