// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/pl33/hta-backend/schemas"
)

// PostCategoryCategoryIDSingleChoiceGroupHandlerFunc turns a function with the right signature into a post category category ID single choice group handler
type PostCategoryCategoryIDSingleChoiceGroupHandlerFunc func(PostCategoryCategoryIDSingleChoiceGroupParams, *schemas.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCategoryCategoryIDSingleChoiceGroupHandlerFunc) Handle(params PostCategoryCategoryIDSingleChoiceGroupParams, principal *schemas.User) middleware.Responder {
	return fn(params, principal)
}

// PostCategoryCategoryIDSingleChoiceGroupHandler interface for that can handle valid post category category ID single choice group params
type PostCategoryCategoryIDSingleChoiceGroupHandler interface {
	Handle(PostCategoryCategoryIDSingleChoiceGroupParams, *schemas.User) middleware.Responder
}

// NewPostCategoryCategoryIDSingleChoiceGroup creates a new http.Handler for the post category category ID single choice group operation
func NewPostCategoryCategoryIDSingleChoiceGroup(ctx *middleware.Context, handler PostCategoryCategoryIDSingleChoiceGroupHandler) *PostCategoryCategoryIDSingleChoiceGroup {
	return &PostCategoryCategoryIDSingleChoiceGroup{Context: ctx, Handler: handler}
}

/*
	PostCategoryCategoryIDSingleChoiceGroup swagger:route POST /category/{category_id}/single_choice_group/ category postCategoryCategoryIdSingleChoiceGroup

PostCategoryCategoryIDSingleChoiceGroup post category category ID single choice group API
*/
type PostCategoryCategoryIDSingleChoiceGroup struct {
	Context *middleware.Context
	Handler PostCategoryCategoryIDSingleChoiceGroupHandler
}

func (o *PostCategoryCategoryIDSingleChoiceGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCategoryCategoryIDSingleChoiceGroupParams()
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
