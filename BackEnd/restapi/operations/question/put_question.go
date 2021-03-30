// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// PutQuestionHandlerFunc turns a function with the right signature into a put question handler
type PutQuestionHandlerFunc func(PutQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PutQuestionHandlerFunc) Handle(params PutQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PutQuestionHandler interface for that can handle valid put question params
type PutQuestionHandler interface {
	Handle(PutQuestionParams, *models.User) middleware.Responder
}

// NewPutQuestion creates a new http.Handler for the put question operation
func NewPutQuestion(ctx *middleware.Context, handler PutQuestionHandler) *PutQuestion {
	return &PutQuestion{Context: ctx, Handler: handler}
}

/* PutQuestion swagger:route PUT /questions/{questionid} question putQuestion

Updates a question

Updates a question

*/
type PutQuestion struct {
	Context *middleware.Context
	Handler PutQuestionHandler
}

func (o *PutQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutQuestionParams()
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
