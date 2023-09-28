// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// GetMultiChoiceIDOKCode is the HTTP code returned for type GetMultiChoiceIDOK
const GetMultiChoiceIDOKCode int = 200

/*
GetMultiChoiceIDOK OK

swagger:response getMultiChoiceIdOK
*/
type GetMultiChoiceIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CategoryMultiChoice `json:"body,omitempty"`
}

// NewGetMultiChoiceIDOK creates GetMultiChoiceIDOK with default headers values
func NewGetMultiChoiceIDOK() *GetMultiChoiceIDOK {

	return &GetMultiChoiceIDOK{}
}

// WithPayload adds the payload to the get multi choice Id o k response
func (o *GetMultiChoiceIDOK) WithPayload(payload *models.CategoryMultiChoice) *GetMultiChoiceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get multi choice Id o k response
func (o *GetMultiChoiceIDOK) SetPayload(payload *models.CategoryMultiChoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMultiChoiceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetMultiChoiceIDDefault Error

swagger:response getMultiChoiceIdDefault
*/
type GetMultiChoiceIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMultiChoiceIDDefault creates GetMultiChoiceIDDefault with default headers values
func NewGetMultiChoiceIDDefault(code int) *GetMultiChoiceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMultiChoiceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get multi choice ID default response
func (o *GetMultiChoiceIDDefault) WithStatusCode(code int) *GetMultiChoiceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get multi choice ID default response
func (o *GetMultiChoiceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get multi choice ID default response
func (o *GetMultiChoiceIDDefault) WithPayload(payload *models.Error) *GetMultiChoiceIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get multi choice ID default response
func (o *GetMultiChoiceIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMultiChoiceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}