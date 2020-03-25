// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SecretReference SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
//
// swagger:model v1.SecretReference
type V1SecretReference struct {

	// Name is unique within a namespace to reference a secret resource.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which the secret name must be unique.
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this v1 secret reference
func (m *V1SecretReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1SecretReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SecretReference) UnmarshalBinary(b []byte) error {
	var res V1SecretReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}