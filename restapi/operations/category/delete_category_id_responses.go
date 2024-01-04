// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// DeleteCategoryIDNoContentCode is the HTTP code returned for type DeleteCategoryIDNoContent
const DeleteCategoryIDNoContentCode int = 204

/*
DeleteCategoryIDNoContent Deleted

swagger:response deleteCategoryIdNoContent
*/
type DeleteCategoryIDNoContent struct {
}

// NewDeleteCategoryIDNoContent creates DeleteCategoryIDNoContent with default headers values
func NewDeleteCategoryIDNoContent() *DeleteCategoryIDNoContent {

	return &DeleteCategoryIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteCategoryIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*
DeleteCategoryIDDefault Error

swagger:response deleteCategoryIdDefault
*/
type DeleteCategoryIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCategoryIDDefault creates DeleteCategoryIDDefault with default headers values
func NewDeleteCategoryIDDefault(code int) *DeleteCategoryIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteCategoryIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete category ID default response
func (o *DeleteCategoryIDDefault) WithStatusCode(code int) *DeleteCategoryIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete category ID default response
func (o *DeleteCategoryIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete category ID default response
func (o *DeleteCategoryIDDefault) WithPayload(payload *models.Error) *DeleteCategoryIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete category ID default response
func (o *DeleteCategoryIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCategoryIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
