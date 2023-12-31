// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// GetSingleChoiceGroupIDOKCode is the HTTP code returned for type GetSingleChoiceGroupIDOK
const GetSingleChoiceGroupIDOKCode int = 200

/*
GetSingleChoiceGroupIDOK OK

swagger:response getSingleChoiceGroupIdOK
*/
type GetSingleChoiceGroupIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.CategorySingleChoiceGroup `json:"body,omitempty"`
}

// NewGetSingleChoiceGroupIDOK creates GetSingleChoiceGroupIDOK with default headers values
func NewGetSingleChoiceGroupIDOK() *GetSingleChoiceGroupIDOK {

	return &GetSingleChoiceGroupIDOK{}
}

// WithPayload adds the payload to the get single choice group Id o k response
func (o *GetSingleChoiceGroupIDOK) WithPayload(payload *models.CategorySingleChoiceGroup) *GetSingleChoiceGroupIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single choice group Id o k response
func (o *GetSingleChoiceGroupIDOK) SetPayload(payload *models.CategorySingleChoiceGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleChoiceGroupIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetSingleChoiceGroupIDDefault Error

swagger:response getSingleChoiceGroupIdDefault
*/
type GetSingleChoiceGroupIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSingleChoiceGroupIDDefault creates GetSingleChoiceGroupIDDefault with default headers values
func NewGetSingleChoiceGroupIDDefault(code int) *GetSingleChoiceGroupIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSingleChoiceGroupIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get single choice group ID default response
func (o *GetSingleChoiceGroupIDDefault) WithStatusCode(code int) *GetSingleChoiceGroupIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get single choice group ID default response
func (o *GetSingleChoiceGroupIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get single choice group ID default response
func (o *GetSingleChoiceGroupIDDefault) WithPayload(payload *models.Error) *GetSingleChoiceGroupIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get single choice group ID default response
func (o *GetSingleChoiceGroupIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSingleChoiceGroupIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
