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

// ListAccountsReader is a Reader for the ListAccounts structure.
type ListAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAccountsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountsOK creates a ListAccountsOK with default headers values
func NewListAccountsOK() *ListAccountsOK {
	return &ListAccountsOK{}
}

/*ListAccountsOK handles this case with default header values.

A list of all the accounts
*/
type ListAccountsOK struct {
	Payload *models.V1beta1AccountManagementList
}

func (o *ListAccountsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/accountmanagements][%d] listAccountsOK  %+v", 200, o.Payload)
}

func (o *ListAccountsOK) GetPayload() *models.V1beta1AccountManagementList {
	return o.Payload
}

func (o *ListAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1beta1AccountManagementList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountsUnauthorized creates a ListAccountsUnauthorized with default headers values
func NewListAccountsUnauthorized() *ListAccountsUnauthorized {
	return &ListAccountsUnauthorized{}
}

/*ListAccountsUnauthorized handles this case with default header values.

If not authenticated
*/
type ListAccountsUnauthorized struct {
}

func (o *ListAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/accountmanagements][%d] listAccountsUnauthorized ", 401)
}

func (o *ListAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAccountsForbidden creates a ListAccountsForbidden with default headers values
func NewListAccountsForbidden() *ListAccountsForbidden {
	return &ListAccountsForbidden{}
}

/*ListAccountsForbidden handles this case with default header values.

If authenticated but not authorized
*/
type ListAccountsForbidden struct {
}

func (o *ListAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/accountmanagements][%d] listAccountsForbidden ", 403)
}

func (o *ListAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAccountsInternalServerError creates a ListAccountsInternalServerError with default headers values
func NewListAccountsInternalServerError() *ListAccountsInternalServerError {
	return &ListAccountsInternalServerError{}
}

/*ListAccountsInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type ListAccountsInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *ListAccountsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1alpha1/accountmanagements][%d] listAccountsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAccountsInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *ListAccountsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}