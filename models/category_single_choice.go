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

// CategorySingleChoice category single choice
//
// swagger:model category_single_choice
type CategorySingleChoice struct {

	// description
	Description string `json:"description,omitempty"`

	// group id
	// Read Only: true
	GroupID int32 `json:"group_id,omitempty"`

	// id
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this category single choice
func (m *CategorySingleChoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CategorySingleChoice) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this category single choice based on the context it is used
func (m *CategorySingleChoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CategorySingleChoice) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "group_id", "body", int32(m.GroupID)); err != nil {
		return err
	}

	return nil
}

func (m *CategorySingleChoice) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CategorySingleChoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategorySingleChoice) UnmarshalBinary(b []byte) error {
	var res CategorySingleChoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
