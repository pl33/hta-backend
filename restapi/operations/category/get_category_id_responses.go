// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// GetCategoryIDOKCode is the HTTP code returned for type GetCategoryIDOK
const GetCategoryIDOKCode int = 200

/*
GetCategoryIDOK OK

swagger:response getCategoryIdOK
*/
type GetCategoryIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Category `json:"body,omitempty"`
}

// NewGetCategoryIDOK creates GetCategoryIDOK with default headers values
func NewGetCategoryIDOK() *GetCategoryIDOK {

	return &GetCategoryIDOK{}
}

// WithPayload adds the payload to the get category Id o k response
func (o *GetCategoryIDOK) WithPayload(payload *models.Category) *GetCategoryIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category Id o k response
func (o *GetCategoryIDOK) SetPayload(payload *models.Category) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetCategoryIDDefault Error

swagger:response getCategoryIdDefault
*/
type GetCategoryIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCategoryIDDefault creates GetCategoryIDDefault with default headers values
func NewGetCategoryIDDefault(code int) *GetCategoryIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCategoryIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get category ID default response
func (o *GetCategoryIDDefault) WithStatusCode(code int) *GetCategoryIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get category ID default response
func (o *GetCategoryIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get category ID default response
func (o *GetCategoryIDDefault) WithPayload(payload *models.Error) *GetCategoryIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category ID default response
func (o *GetCategoryIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
