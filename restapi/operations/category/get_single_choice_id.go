// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/pl33/hta-backend/schemas"
)

// GetSingleChoiceIDHandlerFunc turns a function with the right signature into a get single choice ID handler
type GetSingleChoiceIDHandlerFunc func(GetSingleChoiceIDParams, *schemas.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSingleChoiceIDHandlerFunc) Handle(params GetSingleChoiceIDParams, principal *schemas.User) middleware.Responder {
	return fn(params, principal)
}

// GetSingleChoiceIDHandler interface for that can handle valid get single choice ID params
type GetSingleChoiceIDHandler interface {
	Handle(GetSingleChoiceIDParams, *schemas.User) middleware.Responder
}

// NewGetSingleChoiceID creates a new http.Handler for the get single choice ID operation
func NewGetSingleChoiceID(ctx *middleware.Context, handler GetSingleChoiceIDHandler) *GetSingleChoiceID {
	return &GetSingleChoiceID{Context: ctx, Handler: handler}
}

/*
	GetSingleChoiceID swagger:route GET /single_choice/{id} category getSingleChoiceId

GetSingleChoiceID get single choice ID API
*/
type GetSingleChoiceID struct {
	Context *middleware.Context
	Handler GetSingleChoiceIDHandler
}

func (o *GetSingleChoiceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSingleChoiceIDParams()
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
