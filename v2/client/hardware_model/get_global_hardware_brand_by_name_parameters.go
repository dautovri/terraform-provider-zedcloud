// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHardwareModelGetGlobalHardwareBrandByNameParams creates a new HardwareModelGetGlobalHardwareBrandByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelGetGlobalHardwareBrandByNameParams() *HardwareModelGetGlobalHardwareBrandByNameParams {
	return &HardwareModelGetGlobalHardwareBrandByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelGetGlobalHardwareBrandByNameParamsWithTimeout creates a new HardwareModelGetGlobalHardwareBrandByNameParams object
// with the ability to set a timeout on a request.
func NewHardwareModelGetGlobalHardwareBrandByNameParamsWithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareBrandByNameParams {
	return &HardwareModelGetGlobalHardwareBrandByNameParams{
		timeout: timeout,
	}
}

// NewHardwareModelGetGlobalHardwareBrandByNameParamsWithContext creates a new HardwareModelGetGlobalHardwareBrandByNameParams object
// with the ability to set a context for a request.
func NewHardwareModelGetGlobalHardwareBrandByNameParamsWithContext(ctx context.Context) *HardwareModelGetGlobalHardwareBrandByNameParams {
	return &HardwareModelGetGlobalHardwareBrandByNameParams{
		Context: ctx,
	}
}

// NewHardwareModelGetGlobalHardwareBrandByNameParamsWithHTTPClient creates a new HardwareModelGetGlobalHardwareBrandByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelGetGlobalHardwareBrandByNameParamsWithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareBrandByNameParams {
	return &HardwareModelGetGlobalHardwareBrandByNameParams{
		HTTPClient: client,
	}
}

/*
HardwareModelGetGlobalHardwareBrandByNameParams contains all the parameters to send to the API endpoint

	for the hardware model get global hardware brand by name operation.

	Typically these are written to a http.Request.
*/
type HardwareModelGetGlobalHardwareBrandByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   deprecated field: EnterpriseId
	*/
	EnterpriseID *string

	/* Name.

	   user defined sys brand name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model get global hardware brand by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithDefaults() *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model get global hardware brand by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithContext(ctx context.Context) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithXRequestID(xRequestID *string) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithEnterpriseID(enterpriseID *string) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithName adds the name to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WithName(name string) *HardwareModelGetGlobalHardwareBrandByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the hardware model get global hardware brand by name params
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelGetGlobalHardwareBrandByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.EnterpriseID != nil {

		// query param enterpriseId
		var qrEnterpriseID string

		if o.EnterpriseID != nil {
			qrEnterpriseID = *o.EnterpriseID
		}
		qEnterpriseID := qrEnterpriseID
		if qEnterpriseID != "" {

			if err := r.SetQueryParam("enterpriseId", qEnterpriseID); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
