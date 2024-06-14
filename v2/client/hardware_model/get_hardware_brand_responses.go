// Code generated by go-swagger; DO NOT EDIT.

package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// HardwareModelGetHardwareBrandReader is a Reader for the HardwareModelGetHardwareBrand structure.
type HardwareModelGetHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetHardwareBrandNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetHardwareBrandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetHardwareBrandOK creates a HardwareModelGetHardwareBrandOK with default headers values
func NewHardwareModelGetHardwareBrandOK() *HardwareModelGetHardwareBrandOK {
	return &HardwareModelGetHardwareBrandOK{}
}

/*
HardwareModelGetHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetHardwareBrandOK struct {
	Payload *models.SysBrand
}

// IsSuccess returns true when this hardware model get hardware brand o k response has a 2xx status code
func (o *HardwareModelGetHardwareBrandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get hardware brand o k response has a 3xx status code
func (o *HardwareModelGetHardwareBrandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand o k response has a 4xx status code
func (o *HardwareModelGetHardwareBrandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand o k response has a 5xx status code
func (o *HardwareModelGetHardwareBrandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand o k response a status code equal to that given
func (o *HardwareModelGetHardwareBrandOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hardware model get hardware brand o k response
func (o *HardwareModelGetHardwareBrandOK) Code() int {
	return 200
}

func (o *HardwareModelGetHardwareBrandOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareBrandOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareBrandOK) GetPayload() *models.SysBrand {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandUnauthorized creates a HardwareModelGetHardwareBrandUnauthorized with default headers values
func NewHardwareModelGetHardwareBrandUnauthorized() *HardwareModelGetHardwareBrandUnauthorized {
	return &HardwareModelGetHardwareBrandUnauthorized{}
}

/*
HardwareModelGetHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetHardwareBrandUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand unauthorized response has a 2xx status code
func (o *HardwareModelGetHardwareBrandUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand unauthorized response has a 3xx status code
func (o *HardwareModelGetHardwareBrandUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand unauthorized response has a 4xx status code
func (o *HardwareModelGetHardwareBrandUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand unauthorized response has a 5xx status code
func (o *HardwareModelGetHardwareBrandUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand unauthorized response a status code equal to that given
func (o *HardwareModelGetHardwareBrandUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the hardware model get hardware brand unauthorized response
func (o *HardwareModelGetHardwareBrandUnauthorized) Code() int {
	return 401
}

func (o *HardwareModelGetHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareBrandUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareBrandUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandForbidden creates a HardwareModelGetHardwareBrandForbidden with default headers values
func NewHardwareModelGetHardwareBrandForbidden() *HardwareModelGetHardwareBrandForbidden {
	return &HardwareModelGetHardwareBrandForbidden{}
}

/*
HardwareModelGetHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetHardwareBrandForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand forbidden response has a 2xx status code
func (o *HardwareModelGetHardwareBrandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand forbidden response has a 3xx status code
func (o *HardwareModelGetHardwareBrandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand forbidden response has a 4xx status code
func (o *HardwareModelGetHardwareBrandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand forbidden response has a 5xx status code
func (o *HardwareModelGetHardwareBrandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand forbidden response a status code equal to that given
func (o *HardwareModelGetHardwareBrandForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the hardware model get hardware brand forbidden response
func (o *HardwareModelGetHardwareBrandForbidden) Code() int {
	return 403
}

func (o *HardwareModelGetHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareBrandForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareBrandForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandNotFound creates a HardwareModelGetHardwareBrandNotFound with default headers values
func NewHardwareModelGetHardwareBrandNotFound() *HardwareModelGetHardwareBrandNotFound {
	return &HardwareModelGetHardwareBrandNotFound{}
}

/*
HardwareModelGetHardwareBrandNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetHardwareBrandNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand not found response has a 2xx status code
func (o *HardwareModelGetHardwareBrandNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand not found response has a 3xx status code
func (o *HardwareModelGetHardwareBrandNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand not found response has a 4xx status code
func (o *HardwareModelGetHardwareBrandNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand not found response has a 5xx status code
func (o *HardwareModelGetHardwareBrandNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand not found response a status code equal to that given
func (o *HardwareModelGetHardwareBrandNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the hardware model get hardware brand not found response
func (o *HardwareModelGetHardwareBrandNotFound) Code() int {
	return 404
}

func (o *HardwareModelGetHardwareBrandNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareBrandNotFound) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareBrandNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandInternalServerError creates a HardwareModelGetHardwareBrandInternalServerError with default headers values
func NewHardwareModelGetHardwareBrandInternalServerError() *HardwareModelGetHardwareBrandInternalServerError {
	return &HardwareModelGetHardwareBrandInternalServerError{}
}

/*
HardwareModelGetHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetHardwareBrandInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand internal server error response has a 2xx status code
func (o *HardwareModelGetHardwareBrandInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand internal server error response has a 3xx status code
func (o *HardwareModelGetHardwareBrandInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand internal server error response has a 4xx status code
func (o *HardwareModelGetHardwareBrandInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand internal server error response has a 5xx status code
func (o *HardwareModelGetHardwareBrandInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware brand internal server error response a status code equal to that given
func (o *HardwareModelGetHardwareBrandInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the hardware model get hardware brand internal server error response
func (o *HardwareModelGetHardwareBrandInternalServerError) Code() int {
	return 500
}

func (o *HardwareModelGetHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareBrandInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareBrandInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandGatewayTimeout creates a HardwareModelGetHardwareBrandGatewayTimeout with default headers values
func NewHardwareModelGetHardwareBrandGatewayTimeout() *HardwareModelGetHardwareBrandGatewayTimeout {
	return &HardwareModelGetHardwareBrandGatewayTimeout{}
}

/*
HardwareModelGetHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetHardwareBrandGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand gateway timeout response has a 2xx status code
func (o *HardwareModelGetHardwareBrandGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand gateway timeout response has a 3xx status code
func (o *HardwareModelGetHardwareBrandGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand gateway timeout response has a 4xx status code
func (o *HardwareModelGetHardwareBrandGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand gateway timeout response has a 5xx status code
func (o *HardwareModelGetHardwareBrandGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware brand gateway timeout response a status code equal to that given
func (o *HardwareModelGetHardwareBrandGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the hardware model get hardware brand gateway timeout response
func (o *HardwareModelGetHardwareBrandGatewayTimeout) Code() int {
	return 504
}

func (o *HardwareModelGetHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareBrandGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] hardwareModelGetHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareBrandGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandDefault creates a HardwareModelGetHardwareBrandDefault with default headers values
func NewHardwareModelGetHardwareBrandDefault(code int) *HardwareModelGetHardwareBrandDefault {
	return &HardwareModelGetHardwareBrandDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetHardwareBrandDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetHardwareBrandDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this hardware model get hardware brand default response has a 2xx status code
func (o *HardwareModelGetHardwareBrandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get hardware brand default response has a 3xx status code
func (o *HardwareModelGetHardwareBrandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get hardware brand default response has a 4xx status code
func (o *HardwareModelGetHardwareBrandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get hardware brand default response has a 5xx status code
func (o *HardwareModelGetHardwareBrandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get hardware brand default response a status code equal to that given
func (o *HardwareModelGetHardwareBrandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hardware model get hardware brand default response
func (o *HardwareModelGetHardwareBrandDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelGetHardwareBrandDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] HardwareModel_GetHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareBrandDefault) String() string {
	return fmt.Sprintf("[GET /v1/brands/id/{id}][%d] HardwareModel_GetHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareBrandDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
