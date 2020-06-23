// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageStore ImageStore describes the image store.  Right now only the number
// of images present
//
// swagger:model ImageStore
type ImageStore struct {

	// number
	Number int64 `json:"number,omitempty"`
}

// Validate validates this image store
func (m *ImageStore) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageStore) UnmarshalBinary(b []byte) error {
	var res ImageStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}