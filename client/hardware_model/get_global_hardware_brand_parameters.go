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

// NewHardwareModelGetGlobalHardwareBrandParams creates a new HardwareModelGetGlobalHardwareBrandParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelGetGlobalHardwareBrandParams() *HardwareModelGetGlobalHardwareBrandParams {
	return &HardwareModelGetGlobalHardwareBrandParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelGetGlobalHardwareBrandParamsWithTimeout creates a new HardwareModelGetGlobalHardwareBrandParams object
// with the ability to set a timeout on a request.
func NewHardwareModelGetGlobalHardwareBrandParamsWithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareBrandParams {
	return &HardwareModelGetGlobalHardwareBrandParams{
		timeout: timeout,
	}
}

// NewHardwareModelGetGlobalHardwareBrandParamsWithContext creates a new HardwareModelGetGlobalHardwareBrandParams object
// with the ability to set a context for a request.
func NewHardwareModelGetGlobalHardwareBrandParamsWithContext(ctx context.Context) *HardwareModelGetGlobalHardwareBrandParams {
	return &HardwareModelGetGlobalHardwareBrandParams{
		Context: ctx,
	}
}

// NewHardwareModelGetGlobalHardwareBrandParamsWithHTTPClient creates a new HardwareModelGetGlobalHardwareBrandParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelGetGlobalHardwareBrandParamsWithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareBrandParams {
	return &HardwareModelGetGlobalHardwareBrandParams{
		HTTPClient: client,
	}
}

/*
HardwareModelGetGlobalHardwareBrandParams contains all the parameters to send to the API endpoint

	for the hardware model get global hardware brand operation.

	Typically these are written to a http.Request.
*/
type HardwareModelGetGlobalHardwareBrandParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* EnterpriseID.

	   deprecated field: EnterpriseId
	*/
	EnterpriseID *string

	/* ID.

	   System defined universally unique Id of the model.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model get global hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareBrandParams) WithDefaults() *HardwareModelGetGlobalHardwareBrandParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model get global hardware brand params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetGlobalHardwareBrandParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithTimeout(timeout time.Duration) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithContext(ctx context.Context) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithHTTPClient(client *http.Client) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithXRequestID(xRequestID *string) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEnterpriseID adds the enterpriseID to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithEnterpriseID(enterpriseID *string) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetEnterpriseID(enterpriseID)
	return o
}

// SetEnterpriseID adds the enterpriseId to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetEnterpriseID(enterpriseID *string) {
	o.EnterpriseID = enterpriseID
}

// WithID adds the id to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) WithID(id string) *HardwareModelGetGlobalHardwareBrandParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hardware model get global hardware brand params
func (o *HardwareModelGetGlobalHardwareBrandParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelGetGlobalHardwareBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
