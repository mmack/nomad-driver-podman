// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodContainersPruneReport libpod containers prune report
//
// swagger:model LibpodContainersPruneReport
type LibpodContainersPruneReport struct {

	// ID
	ID string `json:"id,omitempty"`

	// prune error
	PruneError string `json:"error,omitempty"`

	// space reclaimed
	SpaceReclaimed int64 `json:"space,omitempty"`
}

// Validate validates this libpod containers prune report
func (m *LibpodContainersPruneReport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LibpodContainersPruneReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibpodContainersPruneReport) UnmarshalBinary(b []byte) error {
	var res LibpodContainersPruneReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}