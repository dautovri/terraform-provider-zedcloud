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

// HardwareModelQueryHardwareModelsReader is a Reader for the HardwareModelQueryHardwareModels structure.
type HardwareModelQueryHardwareModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelQueryHardwareModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelQueryHardwareModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelQueryHardwareModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelQueryHardwareModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelQueryHardwareModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelQueryHardwareModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelQueryHardwareModelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelQueryHardwareModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelQueryHardwareModelsOK creates a HardwareModelQueryHardwareModelsOK with default headers values
func NewHardwareModelQueryHardwareModelsOK() *HardwareModelQueryHardwareModelsOK {
	return &HardwareModelQueryHardwareModelsOK{}
}

/*
HardwareModelQueryHardwareModelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelQueryHardwareModelsOK struct {
	Payload *models.SysModels
}

// IsSuccess returns true when this hardware model query hardware models o k response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model query hardware models o k response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models o k response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query hardware models o k response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query hardware models o k response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hardware model query hardware models o k response
func (o *HardwareModelQueryHardwareModelsOK) Code() int {
	return 200
}

func (o *HardwareModelQueryHardwareModelsOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsOK) GetPayload() *models.SysModels {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysModels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsBadRequest creates a HardwareModelQueryHardwareModelsBadRequest with default headers values
func NewHardwareModelQueryHardwareModelsBadRequest() *HardwareModelQueryHardwareModelsBadRequest {
	return &HardwareModelQueryHardwareModelsBadRequest{}
}

/*
HardwareModelQueryHardwareModelsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelQueryHardwareModelsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query hardware models bad request response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query hardware models bad request response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models bad request response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query hardware models bad request response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query hardware models bad request response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the hardware model query hardware models bad request response
func (o *HardwareModelQueryHardwareModelsBadRequest) Code() int {
	return 400
}

func (o *HardwareModelQueryHardwareModelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsUnauthorized creates a HardwareModelQueryHardwareModelsUnauthorized with default headers values
func NewHardwareModelQueryHardwareModelsUnauthorized() *HardwareModelQueryHardwareModelsUnauthorized {
	return &HardwareModelQueryHardwareModelsUnauthorized{}
}

/*
HardwareModelQueryHardwareModelsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelQueryHardwareModelsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query hardware models unauthorized response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query hardware models unauthorized response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models unauthorized response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query hardware models unauthorized response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query hardware models unauthorized response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the hardware model query hardware models unauthorized response
func (o *HardwareModelQueryHardwareModelsUnauthorized) Code() int {
	return 401
}

func (o *HardwareModelQueryHardwareModelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsForbidden creates a HardwareModelQueryHardwareModelsForbidden with default headers values
func NewHardwareModelQueryHardwareModelsForbidden() *HardwareModelQueryHardwareModelsForbidden {
	return &HardwareModelQueryHardwareModelsForbidden{}
}

/*
HardwareModelQueryHardwareModelsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelQueryHardwareModelsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query hardware models forbidden response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query hardware models forbidden response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models forbidden response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model query hardware models forbidden response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model query hardware models forbidden response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the hardware model query hardware models forbidden response
func (o *HardwareModelQueryHardwareModelsForbidden) Code() int {
	return 403
}

func (o *HardwareModelQueryHardwareModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsInternalServerError creates a HardwareModelQueryHardwareModelsInternalServerError with default headers values
func NewHardwareModelQueryHardwareModelsInternalServerError() *HardwareModelQueryHardwareModelsInternalServerError {
	return &HardwareModelQueryHardwareModelsInternalServerError{}
}

/*
HardwareModelQueryHardwareModelsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelQueryHardwareModelsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query hardware models internal server error response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query hardware models internal server error response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models internal server error response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query hardware models internal server error response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model query hardware models internal server error response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the hardware model query hardware models internal server error response
func (o *HardwareModelQueryHardwareModelsInternalServerError) Code() int {
	return 500
}

func (o *HardwareModelQueryHardwareModelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsGatewayTimeout creates a HardwareModelQueryHardwareModelsGatewayTimeout with default headers values
func NewHardwareModelQueryHardwareModelsGatewayTimeout() *HardwareModelQueryHardwareModelsGatewayTimeout {
	return &HardwareModelQueryHardwareModelsGatewayTimeout{}
}

/*
HardwareModelQueryHardwareModelsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelQueryHardwareModelsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model query hardware models gateway timeout response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model query hardware models gateway timeout response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model query hardware models gateway timeout response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model query hardware models gateway timeout response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model query hardware models gateway timeout response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the hardware model query hardware models gateway timeout response
func (o *HardwareModelQueryHardwareModelsGatewayTimeout) Code() int {
	return 504
}

func (o *HardwareModelQueryHardwareModelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] hardwareModelQueryHardwareModelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelQueryHardwareModelsDefault creates a HardwareModelQueryHardwareModelsDefault with default headers values
func NewHardwareModelQueryHardwareModelsDefault(code int) *HardwareModelQueryHardwareModelsDefault {
	return &HardwareModelQueryHardwareModelsDefault{
		_statusCode: code,
	}
}

/*
HardwareModelQueryHardwareModelsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelQueryHardwareModelsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this hardware model query hardware models default response has a 2xx status code
func (o *HardwareModelQueryHardwareModelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model query hardware models default response has a 3xx status code
func (o *HardwareModelQueryHardwareModelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model query hardware models default response has a 4xx status code
func (o *HardwareModelQueryHardwareModelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model query hardware models default response has a 5xx status code
func (o *HardwareModelQueryHardwareModelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model query hardware models default response a status code equal to that given
func (o *HardwareModelQueryHardwareModelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hardware model query hardware models default response
func (o *HardwareModelQueryHardwareModelsDefault) Code() int {
	return o._statusCode
}

func (o *HardwareModelQueryHardwareModelsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] HardwareModel_QueryHardwareModels default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsDefault) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels][%d] HardwareModel_QueryHardwareModels default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelQueryHardwareModelsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelQueryHardwareModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
