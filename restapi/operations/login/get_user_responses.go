// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// GetUserOKCode is the HTTP code returned for type GetUserOK
const GetUserOKCode int = 200

/*
GetUserOK User info

swagger:response getUserOK
*/
type GetUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserOK creates GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {

	return &GetUserOK{}
}

// WithPayload adds the payload to the get user o k response
func (o *GetUserOK) WithPayload(payload *models.User) *GetUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user o k response
func (o *GetUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetUserDefault Error

swagger:response getUserDefault
*/
type GetUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserDefault creates GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user default response
func (o *GetUserDefault) WithStatusCode(code int) *GetUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user default response
func (o *GetUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user default response
func (o *GetUserDefault) WithPayload(payload *models.Error) *GetUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user default response
func (o *GetUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
