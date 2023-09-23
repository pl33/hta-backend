// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// GetEntriesOKCode is the HTTP code returned for type GetEntriesOK
const GetEntriesOKCode int = 200

/*
GetEntriesOK List of entries

swagger:response getEntriesOK
*/
type GetEntriesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Entry `json:"body,omitempty"`
}

// NewGetEntriesOK creates GetEntriesOK with default headers values
func NewGetEntriesOK() *GetEntriesOK {

	return &GetEntriesOK{}
}

// WithPayload adds the payload to the get entries o k response
func (o *GetEntriesOK) WithPayload(payload []*models.Entry) *GetEntriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get entries o k response
func (o *GetEntriesOK) SetPayload(payload []*models.Entry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEntriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Entry, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetEntriesDefault Error

swagger:response getEntriesDefault
*/
type GetEntriesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEntriesDefault creates GetEntriesDefault with default headers values
func NewGetEntriesDefault(code int) *GetEntriesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetEntriesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get entries default response
func (o *GetEntriesDefault) WithStatusCode(code int) *GetEntriesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get entries default response
func (o *GetEntriesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get entries default response
func (o *GetEntriesDefault) WithPayload(payload *models.Error) *GetEntriesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get entries default response
func (o *GetEntriesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEntriesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
