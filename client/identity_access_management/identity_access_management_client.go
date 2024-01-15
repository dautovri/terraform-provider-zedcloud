// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new identity access management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for identity access management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IdentityAccessManagementCreateCredential(params *IdentityAccessManagementCreateCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateCredentialOK, error)

	IdentityAccessManagementCreateRole(params *IdentityAccessManagementCreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateRoleOK, error)

	IdentityAccessManagementCreateUser(params *IdentityAccessManagementCreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateUserOK, error)

	IdentityAccessManagementDeleteRole(params *IdentityAccessManagementDeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementDeleteRoleOK, error)

	IdentityAccessManagementDeleteUser(params *IdentityAccessManagementDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementDeleteUserOK, error)

	IdentityAccessManagementGetRole(params *IdentityAccessManagementGetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetRoleOK, error)

	IdentityAccessManagementGetRoleByName(params *IdentityAccessManagementGetRoleByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetRoleByNameOK, error)

	IdentityAccessManagementGetUser(params *IdentityAccessManagementGetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetUserOK, error)

	IdentityAccessManagementGetUserByName(params *IdentityAccessManagementGetUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetUserByNameOK, error)

	IdentityAccessManagementUpdateCredential(params *IdentityAccessManagementUpdateCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateCredentialOK, error)

	IdentityAccessManagementUpdateRole(params *IdentityAccessManagementUpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateRoleOK, error)

	IdentityAccessManagementUpdateUser2(params *IdentityAccessManagementUpdateUser2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateUser2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
IdentityAccessManagementCreateCredential creates i a m credential

Create an IAM credential record.
*/
func (a *Client) IdentityAccessManagementCreateCredential(params *IdentityAccessManagementCreateCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementCreateCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_CreateCredential",
		Method:             "POST",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementCreateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementCreateCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementCreateCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementCreateRole creates i a m role

Create an IAM role record.
*/
func (a *Client) IdentityAccessManagementCreateRole(params *IdentityAccessManagementCreateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementCreateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_CreateRole",
		Method:             "POST",
		PathPattern:        "/v1/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementCreateRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementCreateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementCreateRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementCreateUser creates i a m user

Create an IAM user record.
*/
func (a *Client) IdentityAccessManagementCreateUser(params *IdentityAccessManagementCreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementCreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_CreateUser",
		Method:             "POST",
		PathPattern:        "/v1/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementCreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementCreateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementCreateUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementDeleteCredential deletes i a m credential

Delete an IAM credential record.
*/
func (a *Client) IdentityAccessManagementDeleteCredential(params *IdentityAccessManagementDeleteCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementDeleteCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementDeleteCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_DeleteCredential",
		Method:             "DELETE",
		PathPattern:        "/v1/credentials/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementDeleteCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementDeleteCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementDeleteCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementDeleteRole deletes i a m role

Delete an IAM role record.
*/
func (a *Client) IdentityAccessManagementDeleteRole(params *IdentityAccessManagementDeleteRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementDeleteRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementDeleteRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_DeleteRole",
		Method:             "DELETE",
		PathPattern:        "/v1/roles/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementDeleteRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementDeleteRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementDeleteRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementDeleteUser deletes i a m user

Delete an IAM user record.
*/
func (a *Client) IdentityAccessManagementDeleteUser(params *IdentityAccessManagementDeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementDeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_DeleteUser",
		Method:             "DELETE",
		PathPattern:        "/v1/users/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementDeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementDeleteUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementDeleteUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementGetRole gets i a m role

Get the configuration (without security details) of an IAM role record.
*/
func (a *Client) IdentityAccessManagementGetRole(params *IdentityAccessManagementGetRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementGetRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_GetRole",
		Method:             "GET",
		PathPattern:        "/v1/roles/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementGetRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementGetRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementGetRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementGetRoleByName gets i a m role

Get the configuration (without security details) of an IAM role record.
*/
func (a *Client) IdentityAccessManagementGetRoleByName(params *IdentityAccessManagementGetRoleByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetRoleByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementGetRoleByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_GetRoleByName",
		Method:             "GET",
		PathPattern:        "/v1/roles/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementGetRoleByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementGetRoleByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementGetRoleByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementGetUser gets i a m user

Get the configuration (without security details) of an IAM user record.
*/
func (a *Client) IdentityAccessManagementGetUser(params *IdentityAccessManagementGetUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementGetUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_GetUser",
		Method:             "GET",
		PathPattern:        "/v1/users/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementGetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementGetUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementGetUserDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementGetUserByName gets i a m user

Get the configuration (without security details) of an IAM user record.
*/
func (a *Client) IdentityAccessManagementGetUserByName(params *IdentityAccessManagementGetUserByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementGetUserByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementGetUserByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_GetUserByName",
		Method:             "GET",
		PathPattern:        "/v1/users/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementGetUserByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementGetUserByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementGetUserByNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementUpdateCredential updates i a m credential

Update an IAM credential record.
*/
func (a *Client) IdentityAccessManagementUpdateCredential(params *IdentityAccessManagementUpdateCredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateCredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementUpdateCredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_UpdateCredential",
		Method:             "PUT",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementUpdateCredentialReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementUpdateCredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementUpdateCredentialDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementUpdateRole updates i a m role

Update an IAM role record. The usual pattern to update an IAM role record is to retrieve the record and update with the modified values in a new body to update the IAM role record.
*/
func (a *Client) IdentityAccessManagementUpdateRole(params *IdentityAccessManagementUpdateRoleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementUpdateRoleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_UpdateRole",
		Method:             "PUT",
		PathPattern:        "/v1/roles/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementUpdateRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementUpdateRoleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementUpdateRoleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IdentityAccessManagementUpdateUser2 updates i a m user

Update an IAM user record. The usual pattern to update an IAM user record is to retrieve the record and update with the modified values in a new body to update the IAM user record.
*/
func (a *Client) IdentityAccessManagementUpdateUser2(params *IdentityAccessManagementUpdateUser2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IdentityAccessManagementUpdateUser2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIdentityAccessManagementUpdateUser2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "IdentityAccessManagement_UpdateUser2",
		Method:             "PUT",
		PathPattern:        "/v1/users/id/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IdentityAccessManagementUpdateUser2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IdentityAccessManagementUpdateUser2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IdentityAccessManagementUpdateUser2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
