// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AAAFrontendLogoutResponseCause a a a frontend logout response cause
//
// swagger:model AAA_Frontend_LogoutResponseCause
type AAAFrontendLogoutResponseCause string

func NewAAAFrontendLogoutResponseCause(value AAAFrontendLogoutResponseCause) *AAAFrontendLogoutResponseCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFrontendLogoutResponseCause.
func (m AAAFrontendLogoutResponseCause) Pointer() *AAAFrontendLogoutResponseCause {
	return &m
}

const (

	// AAAFrontendLogoutResponseCauseUNSPECIFIED captures enum value "UNSPECIFIED"
	AAAFrontendLogoutResponseCauseUNSPECIFIED AAAFrontendLogoutResponseCause = "UNSPECIFIED"

	// AAAFrontendLogoutResponseCauseOK captures enum value "OK"
	AAAFrontendLogoutResponseCauseOK AAAFrontendLogoutResponseCause = "OK"

	// AAAFrontendLogoutResponseCauseFAILED captures enum value "FAILED"
	AAAFrontendLogoutResponseCauseFAILED AAAFrontendLogoutResponseCause = "FAILED"
)

// for schema
var aAAFrontendLogoutResponseCauseEnum []interface{}

func init() {
	var res []AAAFrontendLogoutResponseCause
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","OK","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFrontendLogoutResponseCauseEnum = append(aAAFrontendLogoutResponseCauseEnum, v)
	}
}

func (m AAAFrontendLogoutResponseCause) validateAAAFrontendLogoutResponseCauseEnum(path, location string, value AAAFrontendLogoutResponseCause) error {
	if err := validate.EnumCase(path, location, value, aAAFrontendLogoutResponseCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a frontend logout response cause
func (m AAAFrontendLogoutResponseCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFrontendLogoutResponseCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a frontend logout response cause based on context it is used
func (m AAAFrontendLogoutResponseCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
