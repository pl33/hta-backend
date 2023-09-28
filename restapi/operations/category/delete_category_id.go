// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"hta_backend_2/schemas"
)

// DeleteCategoryIDHandlerFunc turns a function with the right signature into a delete category ID handler
type DeleteCategoryIDHandlerFunc func(DeleteCategoryIDParams, *schemas.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCategoryIDHandlerFunc) Handle(params DeleteCategoryIDParams, principal *schemas.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteCategoryIDHandler interface for that can handle valid delete category ID params
type DeleteCategoryIDHandler interface {
	Handle(DeleteCategoryIDParams, *schemas.User) middleware.Responder
}

// NewDeleteCategoryID creates a new http.Handler for the delete category ID operation
func NewDeleteCategoryID(ctx *middleware.Context, handler DeleteCategoryIDHandler) *DeleteCategoryID {
	return &DeleteCategoryID{Context: ctx, Handler: handler}
}

/*
	DeleteCategoryID swagger:route DELETE /category/{id} category deleteCategoryId

DeleteCategoryID delete category ID API
*/
type DeleteCategoryID struct {
	Context *middleware.Context
	Handler DeleteCategoryIDHandler
}

func (o *DeleteCategoryID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCategoryIDParams()
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
