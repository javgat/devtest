// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"uva-devtest/models"
)

// GetCorrectedAnswersFromUserAnsweredTestHandlerFunc turns a function with the right signature into a get corrected answers from user answered test handler
type GetCorrectedAnswersFromUserAnsweredTestHandlerFunc func(GetCorrectedAnswersFromUserAnsweredTestParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCorrectedAnswersFromUserAnsweredTestHandlerFunc) Handle(params GetCorrectedAnswersFromUserAnsweredTestParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetCorrectedAnswersFromUserAnsweredTestHandler interface for that can handle valid get corrected answers from user answered test params
type GetCorrectedAnswersFromUserAnsweredTestHandler interface {
	Handle(GetCorrectedAnswersFromUserAnsweredTestParams, *models.User) middleware.Responder
}

// NewGetCorrectedAnswersFromUserAnsweredTest creates a new http.Handler for the get corrected answers from user answered test operation
func NewGetCorrectedAnswersFromUserAnsweredTest(ctx *middleware.Context, handler GetCorrectedAnswersFromUserAnsweredTestHandler) *GetCorrectedAnswersFromUserAnsweredTest {
	return &GetCorrectedAnswersFromUserAnsweredTest{Context: ctx, Handler: handler}
}

/* GetCorrectedAnswersFromUserAnsweredTest swagger:route GET /users/{username}/answeredTests/{testid}/finishedAnswers user getCorrectedAnswersFromUserAnsweredTest

Returns all answers that the user has answered to a test and are corrected

Returns all answers that the user has answered to a test and are corrected

*/
type GetCorrectedAnswersFromUserAnsweredTest struct {
	Context *middleware.Context
	Handler GetCorrectedAnswersFromUserAnsweredTestHandler
}

func (o *GetCorrectedAnswersFromUserAnsweredTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCorrectedAnswersFromUserAnsweredTestParams()
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