// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pl33/hta-backend/models"
)

// GetCategoryCategoryIDSingleChoiceGroupOKCode is the HTTP code returned for type GetCategoryCategoryIDSingleChoiceGroupOK
const GetCategoryCategoryIDSingleChoiceGroupOKCode int = 200

/*
GetCategoryCategoryIDSingleChoiceGroupOK List of single choice groups

swagger:response getCategoryCategoryIdSingleChoiceGroupOK
*/
type GetCategoryCategoryIDSingleChoiceGroupOK struct {

	/*
	  In: Body
	*/
	Payload []*models.CategorySingleChoiceGroup `json:"body,omitempty"`
}

// NewGetCategoryCategoryIDSingleChoiceGroupOK creates GetCategoryCategoryIDSingleChoiceGroupOK with default headers values
func NewGetCategoryCategoryIDSingleChoiceGroupOK() *GetCategoryCategoryIDSingleChoiceGroupOK {

	return &GetCategoryCategoryIDSingleChoiceGroupOK{}
}

// WithPayload adds the payload to the get category category Id single choice group o k response
func (o *GetCategoryCategoryIDSingleChoiceGroupOK) WithPayload(payload []*models.CategorySingleChoiceGroup) *GetCategoryCategoryIDSingleChoiceGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category category Id single choice group o k response
func (o *GetCategoryCategoryIDSingleChoiceGroupOK) SetPayload(payload []*models.CategorySingleChoiceGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryCategoryIDSingleChoiceGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.CategorySingleChoiceGroup, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetCategoryCategoryIDSingleChoiceGroupDefault Error

swagger:response getCategoryCategoryIdSingleChoiceGroupDefault
*/
type GetCategoryCategoryIDSingleChoiceGroupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCategoryCategoryIDSingleChoiceGroupDefault creates GetCategoryCategoryIDSingleChoiceGroupDefault with default headers values
func NewGetCategoryCategoryIDSingleChoiceGroupDefault(code int) *GetCategoryCategoryIDSingleChoiceGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCategoryCategoryIDSingleChoiceGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get category category ID single choice group default response
func (o *GetCategoryCategoryIDSingleChoiceGroupDefault) WithStatusCode(code int) *GetCategoryCategoryIDSingleChoiceGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get category category ID single choice group default response
func (o *GetCategoryCategoryIDSingleChoiceGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get category category ID single choice group default response
func (o *GetCategoryCategoryIDSingleChoiceGroupDefault) WithPayload(payload *models.Error) *GetCategoryCategoryIDSingleChoiceGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category category ID single choice group default response
func (o *GetCategoryCategoryIDSingleChoiceGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryCategoryIDSingleChoiceGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
