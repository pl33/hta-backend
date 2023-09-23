// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"hta_backend_2/models"
)

// DeleteMultiChoiceIDHandlerFunc turns a function with the right signature into a delete multi choice ID handler
type DeleteMultiChoiceIDHandlerFunc func(DeleteMultiChoiceIDParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteMultiChoiceIDHandlerFunc) Handle(params DeleteMultiChoiceIDParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteMultiChoiceIDHandler interface for that can handle valid delete multi choice ID params
type DeleteMultiChoiceIDHandler interface {
	Handle(DeleteMultiChoiceIDParams, *models.User) middleware.Responder
}

// NewDeleteMultiChoiceID creates a new http.Handler for the delete multi choice ID operation
func NewDeleteMultiChoiceID(ctx *middleware.Context, handler DeleteMultiChoiceIDHandler) *DeleteMultiChoiceID {
	return &DeleteMultiChoiceID{Context: ctx, Handler: handler}
}

/*
	DeleteMultiChoiceID swagger:route DELETE /multi_choice/{id} category deleteMultiChoiceId

DeleteMultiChoiceID delete multi choice ID API
*/
type DeleteMultiChoiceID struct {
	Context *middleware.Context
	Handler DeleteMultiChoiceIDHandler
}

func (o *DeleteMultiChoiceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteMultiChoiceIDParams()
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
