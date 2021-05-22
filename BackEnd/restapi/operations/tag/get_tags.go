// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTagsHandlerFunc turns a function with the right signature into a get tags handler
type GetTagsHandlerFunc func(GetTagsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTagsHandlerFunc) Handle(params GetTagsParams) middleware.Responder {
	return fn(params)
}

// GetTagsHandler interface for that can handle valid get tags params
type GetTagsHandler interface {
	Handle(GetTagsParams) middleware.Responder
}

// NewGetTags creates a new http.Handler for the get tags operation
func NewGetTags(ctx *middleware.Context, handler GetTagsHandler) *GetTags {
	return &GetTags{Context: ctx, Handler: handler}
}

/* GetTags swagger:route GET /tags tag getTags

Returns all tags.

Returns all tags.

*/
type GetTags struct {
	Context *middleware.Context
	Handler GetTagsHandler
}

func (o *GetTags) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTagsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
