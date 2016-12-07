package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// App app
// swagger:model App
type App struct {

	// Application configuration
	Config map[string]string `json:"config,omitempty"`

	// Name of this app. Must be different than the image name. Can ony contain alphanumeric, -, and _.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this app
func (m *App) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *App) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	return nil
}
