// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetOidcInfoParams creates a new GetOidcInfoParams object
//
// There are no default values defined in the spec.
func NewGetOidcInfoParams() GetOidcInfoParams {

	return GetOidcInfoParams{}
}

// GetOidcInfoParams contains all the bound params for the get oidc info operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetOidcInfo
type GetOidcInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetOidcInfoParams() beforehand.
func (o *GetOidcInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
