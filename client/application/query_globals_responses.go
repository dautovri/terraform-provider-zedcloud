package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesReader is a Reader for the EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundles structure.
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK struct {
	Payload *models.Apps
}

// IsSuccess returns true when this edge application configuration query global edge application bundles o k response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration query global edge application bundles o k response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles o k response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query global edge application bundles o k response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query global edge application bundles o k response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration query global edge application bundles o k response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) GetPayload() *models.Apps {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Apps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query global edge application bundles bad request response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query global edge application bundles bad request response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles bad request response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query global edge application bundles bad request response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query global edge application bundles bad request response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge application configuration query global edge application bundles bad request response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) Code() int {
	return 400
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query global edge application bundles unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query global edge application bundles unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query global edge application bundles unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query global edge application bundles unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration query global edge application bundles unauthorized response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query global edge application bundles forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query global edge application bundles forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query global edge application bundles forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query global edge application bundles forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration query global edge application bundles forbidden response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query global edge application bundles internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query global edge application bundles internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query global edge application bundles internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration query global edge application bundles internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration query global edge application bundles internal server error response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout() *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout{}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query global edge application bundles gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query global edge application bundles gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query global edge application bundles gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query global edge application bundles gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration query global edge application bundles gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration query global edge application bundles gateway timeout response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] edgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault creates a EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault with default headers values
func NewEdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault(code int) *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault {
	return &EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration query global edge application bundles default response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration query global edge application bundles default response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration query global edge application bundles default response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration query global edge application bundles default response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration query global edge application bundles default response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration query global edge application bundles default response
func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] EdgeApplicationConfiguration_QueryGlobalEdgeApplicationBundles default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/global][%d] EdgeApplicationConfiguration_QueryGlobalEdgeApplicationBundles default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryGlobalEdgeApplicationBundlesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
