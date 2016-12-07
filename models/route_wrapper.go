package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// RouteWrapper route wrapper
// swagger:model RouteWrapper
type RouteWrapper struct {

	// error
	Error *ErrorBody `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// route
	// Required: true
	Route *Route `json:"route"`
}

// Validate validates this route wrapper
func (m *RouteWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteWrapper) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *RouteWrapper) validateRoute(formats strfmt.Registry) error {

	if err := validate.Required("route", "body", m.Route); err != nil {
		return err
	}

	if m.Route != nil {

		if err := m.Route.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
