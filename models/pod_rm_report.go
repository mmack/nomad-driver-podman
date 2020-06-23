// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodRmReport pod rm report
//
// swagger:model PodRmReport
type PodRmReport struct {

	// err
	Err string `json:"Err,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`
}

// Validate validates this pod rm report
func (m *PodRmReport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodRmReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodRmReport) UnmarshalBinary(b []byte) error {
	var res PodRmReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}