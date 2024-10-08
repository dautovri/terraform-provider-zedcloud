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

// EnterpriseState enterprise state
//
// swagger:model EnterpriseState
type EnterpriseState string

func NewEnterpriseState(value EnterpriseState) *EnterpriseState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EnterpriseState.
func (m EnterpriseState) Pointer() *EnterpriseState {
	return &m
}

const (

	// EnterpriseStateENTERPRISESTATEUNSPECIFIED captures enum value "ENTERPRISE_STATE_UNSPECIFIED"
	EnterpriseStateENTERPRISESTATEUNSPECIFIED EnterpriseState = "ENTERPRISE_STATE_UNSPECIFIED"

	// EnterpriseStateENTERPRISESTATECREATED captures enum value "ENTERPRISE_STATE_CREATED"
	EnterpriseStateENTERPRISESTATECREATED EnterpriseState = "ENTERPRISE_STATE_CREATED"

	// EnterpriseStateENTERPRISESTATEDELETED captures enum value "ENTERPRISE_STATE_DELETED"
	EnterpriseStateENTERPRISESTATEDELETED EnterpriseState = "ENTERPRISE_STATE_DELETED"

	// EnterpriseStateENTERPRISESTATEACTIVE captures enum value "ENTERPRISE_STATE_ACTIVE"
	EnterpriseStateENTERPRISESTATEACTIVE EnterpriseState = "ENTERPRISE_STATE_ACTIVE"

	// EnterpriseStateENTERPRISESTATEINACTIVE captures enum value "ENTERPRISE_STATE_INACTIVE"
	EnterpriseStateENTERPRISESTATEINACTIVE EnterpriseState = "ENTERPRISE_STATE_INACTIVE"

	// EnterpriseStateENTERPRISESTATESIGNEDUP captures enum value "ENTERPRISE_STATE_SIGNEDUP"
	EnterpriseStateENTERPRISESTATESIGNEDUP EnterpriseState = "ENTERPRISE_STATE_SIGNEDUP"
)

// for schema
var enterpriseStateEnum []interface{}

func init() {
	var res []EnterpriseState
	if err := json.Unmarshal([]byte(`["ENTERPRISE_STATE_UNSPECIFIED","ENTERPRISE_STATE_CREATED","ENTERPRISE_STATE_DELETED","ENTERPRISE_STATE_ACTIVE","ENTERPRISE_STATE_INACTIVE","ENTERPRISE_STATE_SIGNEDUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		enterpriseStateEnum = append(enterpriseStateEnum, v)
	}
}

func (m EnterpriseState) validateEnterpriseStateEnum(path, location string, value EnterpriseState) error {
	if err := validate.EnumCase(path, location, value, enterpriseStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this enterprise state
func (m EnterpriseState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEnterpriseStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enterprise state based on context it is used
func (m EnterpriseState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
