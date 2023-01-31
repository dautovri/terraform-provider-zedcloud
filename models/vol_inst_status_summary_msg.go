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

// VolInstStatusSummaryMsg vol inst status summary msg
//
// swagger:model VolInstStatusSummaryMsg
type VolInstStatusSummaryMsg struct {

	// create time
	// Format: date-time
	CreateTime strfmt.DateTime `json:"createTime,omitempty"`

	// device Id
	DeviceID string `json:"deviceId,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// device state
	DeviceState *SWState `json:"deviceState,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// progress percentage
	ProgressPercentage int64 `json:"progressPercentage,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// run state
	RunState *RunState `json:"runState,omitempty"`

	// type
	Type *VolumeInstanceType `json:"type,omitempty"`
}

// Validate validates this vol inst status summary msg
func (m *VolInstStatusSummaryMsg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstStatusSummaryMsg) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VolInstStatusSummaryMsg) validateDeviceState(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceState) { // not required
		return nil
	}

	if m.DeviceState != nil {
		if err := m.DeviceState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceState")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstStatusSummaryMsg) validateRunState(formats strfmt.Registry) error {
	if swag.IsZero(m.RunState) { // not required
		return nil
	}

	if m.RunState != nil {
		if err := m.RunState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstStatusSummaryMsg) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vol inst status summary msg based on the context it is used
func (m *VolInstStatusSummaryMsg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeviceState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstStatusSummaryMsg) contextValidateDeviceState(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceState != nil {
		if err := m.DeviceState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deviceState")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstStatusSummaryMsg) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

	if m.RunState != nil {
		if err := m.RunState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runState")
			}
			return err
		}
	}

	return nil
}

func (m *VolInstStatusSummaryMsg) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolInstStatusSummaryMsg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolInstStatusSummaryMsg) UnmarshalBinary(b []byte) error {
	var res VolInstStatusSummaryMsg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
