// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"hta_backend_2/schemas"
)

// PostEntriesHandlerFunc turns a function with the right signature into a post entries handler
type PostEntriesHandlerFunc func(PostEntriesParams, *schemas.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PostEntriesHandlerFunc) Handle(params PostEntriesParams, principal *schemas.User) middleware.Responder {
	return fn(params, principal)
}

// PostEntriesHandler interface for that can handle valid post entries params
type PostEntriesHandler interface {
	Handle(PostEntriesParams, *schemas.User) middleware.Responder
}

// NewPostEntries creates a new http.Handler for the post entries operation
func NewPostEntries(ctx *middleware.Context, handler PostEntriesHandler) *PostEntries {
	return &PostEntries{Context: ctx, Handler: handler}
}

/*
	PostEntries swagger:route POST /entries/ entry postEntries

PostEntries post entries API
*/
type PostEntries struct {
	Context *middleware.Context
	Handler PostEntriesHandler
}

func (o *PostEntries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostEntriesParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *schemas.User
	if uprinc != nil {
		principal = uprinc.(*schemas.User) // this is really a schemas.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
