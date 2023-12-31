// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// PutSingleChoiceGroupIDOKCode is the HTTP code returned for type PutSingleChoiceGroupIDOK
const PutSingleChoiceGroupIDOKCode int = 200

/*
PutSingleChoiceGroupIDOK OK

swagger:response putSingleChoiceGroupIdOK
*/
type PutSingleChoiceGroupIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CategorySingleChoiceGroup `json:"body,omitempty"`
}

// NewPutSingleChoiceGroupIDOK creates PutSingleChoiceGroupIDOK with default headers values
func NewPutSingleChoiceGroupIDOK() *PutSingleChoiceGroupIDOK {

	return &PutSingleChoiceGroupIDOK{}
}

// WithPayload adds the payload to the put single choice group Id o k response
func (o *PutSingleChoiceGroupIDOK) WithPayload(payload *models.CategorySingleChoiceGroup) *PutSingleChoiceGroupIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put single choice group Id o k response
func (o *PutSingleChoiceGroupIDOK) SetPayload(payload *models.CategorySingleChoiceGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSingleChoiceGroupIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PutSingleChoiceGroupIDDefault Error

swagger:response putSingleChoiceGroupIdDefault
*/
type PutSingleChoiceGroupIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSingleChoiceGroupIDDefault creates PutSingleChoiceGroupIDDefault with default headers values
func NewPutSingleChoiceGroupIDDefault(code int) *PutSingleChoiceGroupIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSingleChoiceGroupIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put single choice group ID default response
func (o *PutSingleChoiceGroupIDDefault) WithStatusCode(code int) *PutSingleChoiceGroupIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put single choice group ID default response
func (o *PutSingleChoiceGroupIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put single choice group ID default response
func (o *PutSingleChoiceGroupIDDefault) WithPayload(payload *models.Error) *PutSingleChoiceGroupIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put single choice group ID default response
func (o *PutSingleChoiceGroupIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSingleChoiceGroupIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
