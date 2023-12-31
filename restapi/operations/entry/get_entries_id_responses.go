// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// GetEntriesIDOKCode is the HTTP code returned for type GetEntriesIDOK
const GetEntriesIDOKCode int = 200

/*
GetEntriesIDOK OK

swagger:response getEntriesIdOK
*/
type GetEntriesIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Entry `json:"body,omitempty"`
}

// NewGetEntriesIDOK creates GetEntriesIDOK with default headers values
func NewGetEntriesIDOK() *GetEntriesIDOK {

	return &GetEntriesIDOK{}
}

// WithPayload adds the payload to the get entries Id o k response
func (o *GetEntriesIDOK) WithPayload(payload *models.Entry) *GetEntriesIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get entries Id o k response
func (o *GetEntriesIDOK) SetPayload(payload *models.Entry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEntriesIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetEntriesIDDefault Error

swagger:response getEntriesIdDefault
*/
type GetEntriesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEntriesIDDefault creates GetEntriesIDDefault with default headers values
func NewGetEntriesIDDefault(code int) *GetEntriesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetEntriesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get entries ID default response
func (o *GetEntriesIDDefault) WithStatusCode(code int) *GetEntriesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get entries ID default response
func (o *GetEntriesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get entries ID default response
func (o *GetEntriesIDDefault) WithPayload(payload *models.Error) *GetEntriesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get entries ID default response
func (o *GetEntriesIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEntriesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
