// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkDNSConfig DNS configuration for a network
// swagger:model network_dns_config
type NetworkDNSConfig struct {

	// enable caching
	// Required: true
	EnableCaching *bool `json:"enable_caching"`

	// local ttl
	// Required: true
	LocalTTL *uint32 `json:"local_ttl"`

	// records
	Records []*DNSConfigRecord `json:"records"`
}

// Validate validates this network dns config
func (m *NetworkDNSConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnableCaching(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkDNSConfig) validateEnableCaching(formats strfmt.Registry) error {

	if err := validate.Required("enable_caching", "body", m.EnableCaching); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDNSConfig) validateLocalTTL(formats strfmt.Registry) error {

	if err := validate.Required("local_ttl", "body", m.LocalTTL); err != nil {
		return err
	}

	return nil
}

func (m *NetworkDNSConfig) validateRecords(formats strfmt.Registry) error {

	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkDNSConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkDNSConfig) UnmarshalBinary(b []byte) error {
	var res NetworkDNSConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
