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

// DeleteServiceCredentialsReader is a Reader for the DeleteServiceCredentials structure.
type DeleteServiceCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteServiceCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServiceCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteServiceCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteServiceCredentialsOK creates a DeleteServiceCredentialsOK with default headers values
func NewDeleteServiceCredentialsOK() *DeleteServiceCredentialsOK {
	return &DeleteServiceCredentialsOK{}
}

/*DeleteServiceCredentialsOK handles this case with default header values.

Contains the former service credentials definition
*/
type DeleteServiceCredentialsOK struct {
	Payload *models.V1ServiceCredentials
}

func (o *DeleteServiceCredentialsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] deleteServiceCredentialsOK  %+v", 200, o.Payload)
}

func (o *DeleteServiceCredentialsOK) GetPayload() *models.V1ServiceCredentials {
	return o.Payload
}

func (o *DeleteServiceCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ServiceCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceCredentialsUnauthorized creates a DeleteServiceCredentialsUnauthorized with default headers values
func NewDeleteServiceCredentialsUnauthorized() *DeleteServiceCredentialsUnauthorized {
	return &DeleteServiceCredentialsUnauthorized{}
}

/*DeleteServiceCredentialsUnauthorized handles this case with default header values.

If not authenticated
*/
type DeleteServiceCredentialsUnauthorized struct {
}

func (o *DeleteServiceCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] deleteServiceCredentialsUnauthorized ", 401)
}

func (o *DeleteServiceCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialsForbidden creates a DeleteServiceCredentialsForbidden with default headers values
func NewDeleteServiceCredentialsForbidden() *DeleteServiceCredentialsForbidden {
	return &DeleteServiceCredentialsForbidden{}
}

/*DeleteServiceCredentialsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type DeleteServiceCredentialsForbidden struct {
}

func (o *DeleteServiceCredentialsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] deleteServiceCredentialsForbidden ", 403)
}

func (o *DeleteServiceCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialsNotFound creates a DeleteServiceCredentialsNotFound with default headers values
func NewDeleteServiceCredentialsNotFound() *DeleteServiceCredentialsNotFound {
	return &DeleteServiceCredentialsNotFound{}
}

/*DeleteServiceCredentialsNotFound handles this case with default header values.

the service credentials with the given name doesn't exist
*/
type DeleteServiceCredentialsNotFound struct {
}

func (o *DeleteServiceCredentialsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] deleteServiceCredentialsNotFound ", 404)
}

func (o *DeleteServiceCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceCredentialsInternalServerError creates a DeleteServiceCredentialsInternalServerError with default headers values
func NewDeleteServiceCredentialsInternalServerError() *DeleteServiceCredentialsInternalServerError {
	return &DeleteServiceCredentialsInternalServerError{}
}

/*DeleteServiceCredentialsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type DeleteServiceCredentialsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *DeleteServiceCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/servicecredentials/{name}][%d] deleteServiceCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServiceCredentialsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *DeleteServiceCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}