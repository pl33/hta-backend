// Code generated by go-swagger; DO NOT EDIT.

package entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*
PostUserCreated Created

swagger:response postUserCreated
*/
type PostUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Entry `json:"body,omitempty"`
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {

	return &PostUserCreated{}
}

// WithPayload adds the payload to the post user created response
func (o *PostUserCreated) WithPayload(payload *models.Entry) *PostUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user created response
func (o *PostUserCreated) SetPayload(payload *models.Entry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostUserDefault Error

swagger:response postUserDefault
*/
type PostUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUserDefault creates PostUserDefault with default headers values
func NewPostUserDefault(code int) *PostUserDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user default response
func (o *PostUserDefault) WithStatusCode(code int) *PostUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user default response
func (o *PostUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user default response
func (o *PostUserDefault) WithPayload(payload *models.Error) *PostUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user default response
func (o *PostUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
