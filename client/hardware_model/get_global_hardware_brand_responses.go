package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider-zedcloud/models"
)

// HardwareModelGetGlobalHardwareBrandReader is a Reader for the HardwareModelGetGlobalHardwareBrand structure.
type HardwareModelGetGlobalHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetGlobalHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetGlobalHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetGlobalHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetGlobalHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetGlobalHardwareBrandNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetGlobalHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetGlobalHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetGlobalHardwareBrandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetGlobalHardwareBrandOK creates a HardwareModelGetGlobalHardwareBrandOK with default headers values
func NewHardwareModelGetGlobalHardwareBrandOK() *HardwareModelGetGlobalHardwareBrandOK {
	return &HardwareModelGetGlobalHardwareBrandOK{}
}

/*
HardwareModelGetGlobalHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetGlobalHardwareBrandOK struct {
	Payload *models.SysBrand
}

// IsSuccess returns true when this hardware model get global hardware brand o k response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get global hardware brand o k response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand o k response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand o k response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand o k response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetGlobalHardwareBrandOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandOK) GetPayload() *models.SysBrand {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandUnauthorized creates a HardwareModelGetGlobalHardwareBrandUnauthorized with default headers values
func NewHardwareModelGetGlobalHardwareBrandUnauthorized() *HardwareModelGetGlobalHardwareBrandUnauthorized {
	return &HardwareModelGetGlobalHardwareBrandUnauthorized{}
}

/*
HardwareModelGetGlobalHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetGlobalHardwareBrandUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand unauthorized response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand unauthorized response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand unauthorized response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand unauthorized response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand unauthorized response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandForbidden creates a HardwareModelGetGlobalHardwareBrandForbidden with default headers values
func NewHardwareModelGetGlobalHardwareBrandForbidden() *HardwareModelGetGlobalHardwareBrandForbidden {
	return &HardwareModelGetGlobalHardwareBrandForbidden{}
}

/*
HardwareModelGetGlobalHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetGlobalHardwareBrandForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand forbidden response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand forbidden response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand forbidden response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand forbidden response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand forbidden response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetGlobalHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandNotFound creates a HardwareModelGetGlobalHardwareBrandNotFound with default headers values
func NewHardwareModelGetGlobalHardwareBrandNotFound() *HardwareModelGetGlobalHardwareBrandNotFound {
	return &HardwareModelGetGlobalHardwareBrandNotFound{}
}

/*
HardwareModelGetGlobalHardwareBrandNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetGlobalHardwareBrandNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand not found response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand not found response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand not found response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware brand not found response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware brand not found response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelGetGlobalHardwareBrandNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandNotFound) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandInternalServerError creates a HardwareModelGetGlobalHardwareBrandInternalServerError with default headers values
func NewHardwareModelGetGlobalHardwareBrandInternalServerError() *HardwareModelGetGlobalHardwareBrandInternalServerError {
	return &HardwareModelGetGlobalHardwareBrandInternalServerError{}
}

/*
HardwareModelGetGlobalHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetGlobalHardwareBrandInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand internal server error response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand internal server error response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand internal server error response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand internal server error response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware brand internal server error response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandGatewayTimeout creates a HardwareModelGetGlobalHardwareBrandGatewayTimeout with default headers values
func NewHardwareModelGetGlobalHardwareBrandGatewayTimeout() *HardwareModelGetGlobalHardwareBrandGatewayTimeout {
	return &HardwareModelGetGlobalHardwareBrandGatewayTimeout{}
}

/*
HardwareModelGetGlobalHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetGlobalHardwareBrandGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware brand gateway timeout response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware brand gateway timeout response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware brand gateway timeout response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware brand gateway timeout response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware brand gateway timeout response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] hardwareModelGetGlobalHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareBrandDefault creates a HardwareModelGetGlobalHardwareBrandDefault with default headers values
func NewHardwareModelGetGlobalHardwareBrandDefault(code int) *HardwareModelGetGlobalHardwareBrandDefault {
	return &HardwareModelGetGlobalHardwareBrandDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetGlobalHardwareBrandDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetGlobalHardwareBrandDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get global hardware brand default response
func (o *HardwareModelGetGlobalHardwareBrandDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get global hardware brand default response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareBrandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get global hardware brand default response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareBrandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get global hardware brand default response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareBrandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get global hardware brand default response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareBrandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get global hardware brand default response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareBrandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetGlobalHardwareBrandDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] HardwareModel_GetGlobalHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandDefault) String() string {
	return fmt.Sprintf("[GET /v1/brands/global/id/{id}][%d] HardwareModel_GetGlobalHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareBrandDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareBrandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
