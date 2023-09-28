// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Entry entry
//
// swagger:model entry
type Entry struct {

	// diastole
	Diastole float32 `json:"diastole,omitempty"`

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// have blood pressure
	// Required: true
	HaveBloodPressure *bool `json:"have_blood_pressure"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// multi choices
	// Required: true
	MultiChoices []int64 `json:"multi_choices"`

	// pulse
	Pulse float32 `json:"pulse,omitempty"`

	// remarks
	Remarks string `json:"remarks,omitempty"`

	// single choices
	// Required: true
	SingleChoices []int64 `json:"single_choices"`

	// start time
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time"`

	// systole
	Systole float32 `json:"systole,omitempty"`

	// user id
	// Read Only: true
	UserID int32 `json:"user_id,omitempty"`
}

// Validate validates this entry
func (m *Entry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaveBloodPressure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiChoices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleChoices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entry) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Entry) validateHaveBloodPressure(formats strfmt.Registry) error {

	if err := validate.Required("have_blood_pressure", "body", m.HaveBloodPressure); err != nil {
		return err
	}

	return nil
}

func (m *Entry) validateMultiChoices(formats strfmt.Registry) error {

	if err := validate.Required("multi_choices", "body", m.MultiChoices); err != nil {
		return err
	}

	return nil
}

func (m *Entry) validateSingleChoices(formats strfmt.Registry) error {

	if err := validate.Required("single_choices", "body", m.SingleChoices); err != nil {
		return err
	}

	return nil
}

func (m *Entry) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this entry based on the context it is used
func (m *Entry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entry) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Entry) contextValidateUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_id", "body", int32(m.UserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Entry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entry) UnmarshalBinary(b []byte) error {
	var res Entry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}