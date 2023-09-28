// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// PutMultiChoiceIDOKCode is the HTTP code returned for type PutMultiChoiceIDOK
const PutMultiChoiceIDOKCode int = 200

/*
PutMultiChoiceIDOK OK

swagger:response putMultiChoiceIdOK
*/
type PutMultiChoiceIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CategoryMultiChoice `json:"body,omitempty"`
}

// NewPutMultiChoiceIDOK creates PutMultiChoiceIDOK with default headers values
func NewPutMultiChoiceIDOK() *PutMultiChoiceIDOK {

	return &PutMultiChoiceIDOK{}
}

// WithPayload adds the payload to the put multi choice Id o k response
func (o *PutMultiChoiceIDOK) WithPayload(payload *models.CategoryMultiChoice) *PutMultiChoiceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put multi choice Id o k response
func (o *PutMultiChoiceIDOK) SetPayload(payload *models.CategoryMultiChoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMultiChoiceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutMultiChoiceIDDefault Error

swagger:response putMultiChoiceIdDefault
*/
type PutMultiChoiceIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutMultiChoiceIDDefault creates PutMultiChoiceIDDefault with default headers values
func NewPutMultiChoiceIDDefault(code int) *PutMultiChoiceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutMultiChoiceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put multi choice ID default response
func (o *PutMultiChoiceIDDefault) WithStatusCode(code int) *PutMultiChoiceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put multi choice ID default response
func (o *PutMultiChoiceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put multi choice ID default response
func (o *PutMultiChoiceIDDefault) WithPayload(payload *models.Error) *PutMultiChoiceIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put multi choice ID default response
func (o *PutMultiChoiceIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutMultiChoiceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}