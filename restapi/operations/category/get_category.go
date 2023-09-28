// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"hta_backend_2/schemas"
)

// GetCategoryHandlerFunc turns a function with the right signature into a get category handler
type GetCategoryHandlerFunc func(GetCategoryParams, *schemas.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCategoryHandlerFunc) Handle(params GetCategoryParams, principal *schemas.User) middleware.Responder {
	return fn(params, principal)
}

// GetCategoryHandler interface for that can handle valid get category params
type GetCategoryHandler interface {
	Handle(GetCategoryParams, *schemas.User) middleware.Responder
}

// NewGetCategory creates a new http.Handler for the get category operation
func NewGetCategory(ctx *middleware.Context, handler GetCategoryHandler) *GetCategory {
	return &GetCategory{Context: ctx, Handler: handler}
}

/*
	GetCategory swagger:route GET /category/ category getCategory

GetCategory get category API
*/
type GetCategory struct {
	Context *middleware.Context
	Handler GetCategoryHandler
}

func (o *GetCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCategoryParams()
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
