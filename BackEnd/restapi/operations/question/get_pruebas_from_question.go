// Code generated by go-swagger; DO NOT EDIT.

package question

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetPruebasFromQuestionHandlerFunc turns a function with the right signature into a get pruebas from question handler
type GetPruebasFromQuestionHandlerFunc func(GetPruebasFromQuestionParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPruebasFromQuestionHandlerFunc) Handle(params GetPruebasFromQuestionParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetPruebasFromQuestionHandler interface for that can handle valid get pruebas from question params
type GetPruebasFromQuestionHandler interface {
	Handle(GetPruebasFromQuestionParams, *models.User) middleware.Responder
}

// NewGetPruebasFromQuestion creates a new http.Handler for the get pruebas from question operation
func NewGetPruebasFromQuestion(ctx *middleware.Context, handler GetPruebasFromQuestionHandler) *GetPruebasFromQuestion {
	return &GetPruebasFromQuestion{Context: ctx, Handler: handler}
}

/* GetPruebasFromQuestion swagger:route GET /questions/{questionid}/pruebas question getPruebasFromQuestion

Returns all pruebas from a question.

Returns all pruebas from a question.

*/
type GetPruebasFromQuestion struct {
	Context *middleware.Context
	Handler GetPruebasFromQuestionHandler
}

func (o *GetPruebasFromQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPruebasFromQuestionParams()
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
