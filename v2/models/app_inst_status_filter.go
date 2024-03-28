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

// AppInstStatusFilter app inst status filter
//
// swagger:model AppInstStatusFilter
type AppInstStatusFilter struct {

	// User defined name of the app bundle, unique across the enterprise. Once app bundle is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	AppName string `json:"appName,omitempty"`

	// type of app
	AppType *AppType `json:"appType,omitempty"`

	// type of deployment for the app, eg: azure, k3s, standalone
	DeploymentType *DeploymentType `json:"deploymentType,omitempty"`

	// User defined name of the device, unique across the enterprise. Once device is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	DeviceName string `json:"deviceName,omitempty"`

	// device name pattern
	DeviceNamePattern string `json:"deviceNamePattern,omitempty"`

	// name pattern
	NamePattern string `json:"namePattern,omitempty"`

	// User defined name of the project, unique across the enterprise. Once project is created, name can’t be changed
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ProjectName string `json:"projectName,omitempty"`

	// project name pattern
	ProjectNamePattern string `json:"projectNamePattern,omitempty"`

	// aperation status
	RunState *RunState `json:"runState,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this app inst status filter
func (m *AppInstStatusFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstStatusFilter) validateAppName(formats strfmt.Registry) error {
	if swag.IsZero(m.AppName) { // not required
		return nil
	}

	if err := validate.MinLength("appName", "body", m.AppName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("appName", "body", m.AppName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("appName", "body", m.AppName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusFilter) validateAppType(formats strfmt.Registry) error {
	if swag.IsZero(m.AppType) { // not required
		return nil
	}

	if m.AppType != nil {
		if err := m.AppType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusFilter) validateDeploymentType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentType) { // not required
		return nil
	}

	if m.DeploymentType != nil {
		if err := m.DeploymentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusFilter) validateDeviceName(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceName) { // not required
		return nil
	}

	if err := validate.MinLength("deviceName", "body", m.DeviceName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("deviceName", "body", m.DeviceName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("deviceName", "body", m.DeviceName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusFilter) validateProjectName(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectName) { // not required
		return nil
	}

	if err := validate.MinLength("projectName", "body", m.ProjectName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("projectName", "body", m.ProjectName, 256); err != nil {
		return err
	}

	if err := validate.Pattern("projectName", "body", m.ProjectName, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AppInstStatusFilter) validateRunState(formats strfmt.Registry) error {
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

// ContextValidate validate this app inst status filter based on the context it is used
func (m *AppInstStatusFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstStatusFilter) contextValidateAppType(ctx context.Context, formats strfmt.Registry) error {

	if m.AppType != nil {
		if err := m.AppType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusFilter) contextValidateDeploymentType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentType != nil {
		if err := m.DeploymentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deploymentType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deploymentType")
			}
			return err
		}
	}

	return nil
}

func (m *AppInstStatusFilter) contextValidateRunState(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *AppInstStatusFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstStatusFilter) UnmarshalBinary(b []byte) error {
	var res AppInstStatusFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}