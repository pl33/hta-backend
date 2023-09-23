// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hta_backend_2/models"
)

// GetCategoryCategoryIDMultiChoiceOKCode is the HTTP code returned for type GetCategoryCategoryIDMultiChoiceOK
const GetCategoryCategoryIDMultiChoiceOKCode int = 200

/*
GetCategoryCategoryIDMultiChoiceOK List of multi choices

swagger:response getCategoryCategoryIdMultiChoiceOK
*/
type GetCategoryCategoryIDMultiChoiceOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CategoryMultiChoice `json:"body,omitempty"`
}

// NewGetCategoryCategoryIDMultiChoiceOK creates GetCategoryCategoryIDMultiChoiceOK with default headers values
func NewGetCategoryCategoryIDMultiChoiceOK() *GetCategoryCategoryIDMultiChoiceOK {

	return &GetCategoryCategoryIDMultiChoiceOK{}
}

// WithPayload adds the payload to the get category category Id multi choice o k response
func (o *GetCategoryCategoryIDMultiChoiceOK) WithPayload(payload []*models.CategoryMultiChoice) *GetCategoryCategoryIDMultiChoiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category category Id multi choice o k response
func (o *GetCategoryCategoryIDMultiChoiceOK) SetPayload(payload []*models.CategoryMultiChoice) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryCategoryIDMultiChoiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CategoryMultiChoice, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetCategoryCategoryIDMultiChoiceDefault Error

swagger:response getCategoryCategoryIdMultiChoiceDefault
*/
type GetCategoryCategoryIDMultiChoiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCategoryCategoryIDMultiChoiceDefault creates GetCategoryCategoryIDMultiChoiceDefault with default headers values
func NewGetCategoryCategoryIDMultiChoiceDefault(code int) *GetCategoryCategoryIDMultiChoiceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCategoryCategoryIDMultiChoiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get category category ID multi choice default response
func (o *GetCategoryCategoryIDMultiChoiceDefault) WithStatusCode(code int) *GetCategoryCategoryIDMultiChoiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get category category ID multi choice default response
func (o *GetCategoryCategoryIDMultiChoiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get category category ID multi choice default response
func (o *GetCategoryCategoryIDMultiChoiceDefault) WithPayload(payload *models.Error) *GetCategoryCategoryIDMultiChoiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category category ID multi choice default response
func (o *GetCategoryCategoryIDMultiChoiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryCategoryIDMultiChoiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
