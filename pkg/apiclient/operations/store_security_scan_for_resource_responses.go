// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// StoreSecurityScanForResourceReader is a Reader for the StoreSecurityScanForResource structure.
type StoreSecurityScanForResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StoreSecurityScanForResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStoreSecurityScanForResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStoreSecurityScanForResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStoreSecurityScanForResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStoreSecurityScanForResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStoreSecurityScanForResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStoreSecurityScanForResourceOK creates a StoreSecurityScanForResourceOK with default headers values
func NewStoreSecurityScanForResourceOK() *StoreSecurityScanForResourceOK {
	return &StoreSecurityScanForResourceOK{}
}

/*StoreSecurityScanForResourceOK handles this case with default header values.

Latest security scan for this resource
*/
type StoreSecurityScanForResourceOK struct {
	Payload *models.V1SecurityScanResult
}

func (o *StoreSecurityScanForResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/security/scans/{group}/{version}/{kind}/{namespace}/{name}][%d] storeSecurityScanForResourceOK  %+v", 200, o.Payload)
}

func (o *StoreSecurityScanForResourceOK) GetPayload() *models.V1SecurityScanResult {
	return o.Payload
}

func (o *StoreSecurityScanForResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SecurityScanResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoreSecurityScanForResourceBadRequest creates a StoreSecurityScanForResourceBadRequest with default headers values
func NewStoreSecurityScanForResourceBadRequest() *StoreSecurityScanForResourceBadRequest {
	return &StoreSecurityScanForResourceBadRequest{}
}

/*StoreSecurityScanForResourceBadRequest handles this case with default header values.

Validation error of supplied parameters/body
*/
type StoreSecurityScanForResourceBadRequest struct {
	Payload *models.ValidationError
}

func (o *StoreSecurityScanForResourceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/security/scans/{group}/{version}/{kind}/{namespace}/{name}][%d] storeSecurityScanForResourceBadRequest  %+v", 400, o.Payload)
}

func (o *StoreSecurityScanForResourceBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *StoreSecurityScanForResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStoreSecurityScanForResourceUnauthorized creates a StoreSecurityScanForResourceUnauthorized with default headers values
func NewStoreSecurityScanForResourceUnauthorized() *StoreSecurityScanForResourceUnauthorized {
	return &StoreSecurityScanForResourceUnauthorized{}
}

/*StoreSecurityScanForResourceUnauthorized handles this case with default header values.

If not authenticated
*/
type StoreSecurityScanForResourceUnauthorized struct {
}

func (o *StoreSecurityScanForResourceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/security/scans/{group}/{version}/{kind}/{namespace}/{name}][%d] storeSecurityScanForResourceUnauthorized ", 401)
}

func (o *StoreSecurityScanForResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStoreSecurityScanForResourceForbidden creates a StoreSecurityScanForResourceForbidden with default headers values
func NewStoreSecurityScanForResourceForbidden() *StoreSecurityScanForResourceForbidden {
	return &StoreSecurityScanForResourceForbidden{}
}

/*StoreSecurityScanForResourceForbidden handles this case with default header values.

If authenticated but not authorized
*/
type StoreSecurityScanForResourceForbidden struct {
}

func (o *StoreSecurityScanForResourceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/security/scans/{group}/{version}/{kind}/{namespace}/{name}][%d] storeSecurityScanForResourceForbidden ", 403)
}

func (o *StoreSecurityScanForResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStoreSecurityScanForResourceInternalServerError creates a StoreSecurityScanForResourceInternalServerError with default headers values
func NewStoreSecurityScanForResourceInternalServerError() *StoreSecurityScanForResourceInternalServerError {
	return &StoreSecurityScanForResourceInternalServerError{}
}

/*StoreSecurityScanForResourceInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type StoreSecurityScanForResourceInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *StoreSecurityScanForResourceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1alpha1/security/scans/{group}/{version}/{kind}/{namespace}/{name}][%d] storeSecurityScanForResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *StoreSecurityScanForResourceInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *StoreSecurityScanForResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}