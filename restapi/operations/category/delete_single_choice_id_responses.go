// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// DeleteSingleChoiceIDNoContentCode is the HTTP code returned for type DeleteSingleChoiceIDNoContent
const DeleteSingleChoiceIDNoContentCode int = 204

/*
DeleteSingleChoiceIDNoContent Deleted

swagger:response deleteSingleChoiceIdNoContent
*/
type DeleteSingleChoiceIDNoContent struct {
}

// NewDeleteSingleChoiceIDNoContent creates DeleteSingleChoiceIDNoContent with default headers values
func NewDeleteSingleChoiceIDNoContent() *DeleteSingleChoiceIDNoContent {

	return &DeleteSingleChoiceIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSingleChoiceIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteSingleChoiceIDDefault Error

swagger:response deleteSingleChoiceIdDefault
*/
type DeleteSingleChoiceIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSingleChoiceIDDefault creates DeleteSingleChoiceIDDefault with default headers values
func NewDeleteSingleChoiceIDDefault(code int) *DeleteSingleChoiceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSingleChoiceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete single choice ID default response
func (o *DeleteSingleChoiceIDDefault) WithStatusCode(code int) *DeleteSingleChoiceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete single choice ID default response
func (o *DeleteSingleChoiceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete single choice ID default response
func (o *DeleteSingleChoiceIDDefault) WithPayload(payload *models.Error) *DeleteSingleChoiceIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete single choice ID default response
func (o *DeleteSingleChoiceIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSingleChoiceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
