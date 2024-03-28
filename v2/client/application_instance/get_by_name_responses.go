package application_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameReader is a Reader for the EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByName structure.
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK struct {
	Payload *models.AppInstance
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name o k response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name o k response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name o k response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name o k response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance by name o k response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application instance configuration get edge application instance by name o k response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) Code() int {
	return 200
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) GetPayload() *models.AppInstance {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name unauthorized response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name unauthorized response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name unauthorized response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name unauthorized response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance by name unauthorized response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application instance configuration get edge application instance by name unauthorized response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name forbidden response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name forbidden response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name forbidden response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name forbidden response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance by name forbidden response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application instance configuration get edge application instance by name forbidden response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name not found response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name not found response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name not found response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name not found response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application instance configuration get edge application instance by name not found response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application instance configuration get edge application instance by name not found response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name internal server error response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name internal server error response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name internal server error response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name internal server error response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration get edge application instance by name internal server error response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application instance configuration get edge application instance by name internal server error response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout() *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout{}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name gateway timeout response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name gateway timeout response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name gateway timeout response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name gateway timeout response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application instance configuration get edge application instance by name gateway timeout response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application instance configuration get edge application instance by name gateway timeout response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] edgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault creates a EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault with default headers values
func NewEdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault(code int) *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault {
	return &EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application instance configuration get edge application instance by name default response has a 2xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application instance configuration get edge application instance by name default response has a 3xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application instance configuration get edge application instance by name default response has a 4xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application instance configuration get edge application instance by name default response has a 5xx status code
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application instance configuration get edge application instance by name default response a status code equal to that given
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application instance configuration get edge application instance by name default response
func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] EdgeApplicationInstanceConfiguration_GetEdgeApplicationInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/instances/name/{name}][%d] EdgeApplicationInstanceConfiguration_GetEdgeApplicationInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationInstanceConfigurationGetEdgeApplicationInstanceByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}