// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// PostSingleChoiceGroupGroupIDSingleChoiceCreatedCode is the HTTP code returned for type PostSingleChoiceGroupGroupIDSingleChoiceCreated
const PostSingleChoiceGroupGroupIDSingleChoiceCreatedCode int = 201

/*
PostSingleChoiceGroupGroupIDSingleChoiceCreated Created

swagger:response postSingleChoiceGroupGroupIdSingleChoiceCreated
*/
type PostSingleChoiceGroupGroupIDSingleChoiceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CategorySingleChoice `json:"body,omitempty"`
}

// NewPostSingleChoiceGroupGroupIDSingleChoiceCreated creates PostSingleChoiceGroupGroupIDSingleChoiceCreated with default headers values
func NewPostSingleChoiceGroupGroupIDSingleChoiceCreated() *PostSingleChoiceGroupGroupIDSingleChoiceCreated {

	return &PostSingleChoiceGroupGroupIDSingleChoiceCreated{}
}

// WithPayload adds the payload to the post single choice group group Id single choice created response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceCreated) WithPayload(payload *models.CategorySingleChoice) *PostSingleChoiceGroupGroupIDSingleChoiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post single choice group group Id single choice created response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceCreated) SetPayload(payload *models.CategorySingleChoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSingleChoiceGroupGroupIDSingleChoiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostSingleChoiceGroupGroupIDSingleChoiceDefault Error

swagger:response postSingleChoiceGroupGroupIdSingleChoiceDefault
*/
type PostSingleChoiceGroupGroupIDSingleChoiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostSingleChoiceGroupGroupIDSingleChoiceDefault creates PostSingleChoiceGroupGroupIDSingleChoiceDefault with default headers values
func NewPostSingleChoiceGroupGroupIDSingleChoiceDefault(code int) *PostSingleChoiceGroupGroupIDSingleChoiceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostSingleChoiceGroupGroupIDSingleChoiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post single choice group group ID single choice default response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceDefault) WithStatusCode(code int) *PostSingleChoiceGroupGroupIDSingleChoiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post single choice group group ID single choice default response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post single choice group group ID single choice default response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceDefault) WithPayload(payload *models.Error) *PostSingleChoiceGroupGroupIDSingleChoiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post single choice group group ID single choice default response
func (o *PostSingleChoiceGroupGroupIDSingleChoiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSingleChoiceGroupGroupIDSingleChoiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
