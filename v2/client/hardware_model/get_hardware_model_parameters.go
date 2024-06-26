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

// NewHardwareModelGetHardwareModelParams creates a new HardwareModelGetHardwareModelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHardwareModelGetHardwareModelParams() *HardwareModelGetHardwareModelParams {
	return &HardwareModelGetHardwareModelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHardwareModelGetHardwareModelParamsWithTimeout creates a new HardwareModelGetHardwareModelParams object
// with the ability to set a timeout on a request.
func NewHardwareModelGetHardwareModelParamsWithTimeout(timeout time.Duration) *HardwareModelGetHardwareModelParams {
	return &HardwareModelGetHardwareModelParams{
		timeout: timeout,
	}
}

// NewHardwareModelGetHardwareModelParamsWithContext creates a new HardwareModelGetHardwareModelParams object
// with the ability to set a context for a request.
func NewHardwareModelGetHardwareModelParamsWithContext(ctx context.Context) *HardwareModelGetHardwareModelParams {
	return &HardwareModelGetHardwareModelParams{
		Context: ctx,
	}
}

// NewHardwareModelGetHardwareModelParamsWithHTTPClient creates a new HardwareModelGetHardwareModelParams object
// with the ability to set a custom HTTPClient for a request.
func NewHardwareModelGetHardwareModelParamsWithHTTPClient(client *http.Client) *HardwareModelGetHardwareModelParams {
	return &HardwareModelGetHardwareModelParams{
		HTTPClient: client,
	}
}

/*
HardwareModelGetHardwareModelParams contains all the parameters to send to the API endpoint

	for the hardware model get hardware model operation.

	Typically these are written to a http.Request.
*/
type HardwareModelGetHardwareModelParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the  model
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the hardware model get hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetHardwareModelParams) WithDefaults() *HardwareModelGetHardwareModelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the hardware model get hardware model params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HardwareModelGetHardwareModelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) WithTimeout(timeout time.Duration) *HardwareModelGetHardwareModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) WithContext(ctx context.Context) *HardwareModelGetHardwareModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) WithHTTPClient(client *http.Client) *HardwareModelGetHardwareModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) WithXRequestID(xRequestID *string) *HardwareModelGetHardwareModelParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) WithID(id string) *HardwareModelGetHardwareModelParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the hardware model get hardware model params
func (o *HardwareModelGetHardwareModelParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *HardwareModelGetHardwareModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
