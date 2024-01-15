// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

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

// NewIdentityAccessManagementGetUserSessionParams creates a new IdentityAccessManagementGetUserSessionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementGetUserSessionParams() *IdentityAccessManagementGetUserSessionParams {
	return &IdentityAccessManagementGetUserSessionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementGetUserSessionParamsWithTimeout creates a new IdentityAccessManagementGetUserSessionParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementGetUserSessionParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserSessionParams {
	return &IdentityAccessManagementGetUserSessionParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementGetUserSessionParamsWithContext creates a new IdentityAccessManagementGetUserSessionParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementGetUserSessionParamsWithContext(ctx context.Context) *IdentityAccessManagementGetUserSessionParams {
	return &IdentityAccessManagementGetUserSessionParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementGetUserSessionParamsWithHTTPClient creates a new IdentityAccessManagementGetUserSessionParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementGetUserSessionParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserSessionParams {
	return &IdentityAccessManagementGetUserSessionParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementGetUserSessionParams contains all the parameters to send to the API endpoint

	for the identity access management get user session operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementGetUserSessionParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// SessionTokenBase64.
	SessionTokenBase64 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management get user session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserSessionParams) WithDefaults() *IdentityAccessManagementGetUserSessionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management get user session params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetUserSessionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementGetUserSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) WithContext(ctx context.Context) *IdentityAccessManagementGetUserSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementGetUserSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementGetUserSessionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSessionTokenBase64 adds the sessionTokenBase64 to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) WithSessionTokenBase64(sessionTokenBase64 string) *IdentityAccessManagementGetUserSessionParams {
	o.SetSessionTokenBase64(sessionTokenBase64)
	return o
}

// SetSessionTokenBase64 adds the sessionTokenBase64 to the identity access management get user session params
func (o *IdentityAccessManagementGetUserSessionParams) SetSessionTokenBase64(sessionTokenBase64 string) {
	o.SessionTokenBase64 = sessionTokenBase64
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementGetUserSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param sessionToken.base64
	if err := r.SetPathParam("sessionToken.base64", o.SessionTokenBase64); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
