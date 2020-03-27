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

// RemoveKubernetesReader is a Reader for the RemoveKubernetes structure.
type RemoveKubernetesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveKubernetesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveKubernetesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRemoveKubernetesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveKubernetesOK creates a RemoveKubernetesOK with default headers values
func NewRemoveKubernetesOK() *RemoveKubernetesOK {
	return &RemoveKubernetesOK{}
}

/*RemoveKubernetesOK handles this case with default header values.

Contains the former team definition from the kore
*/
type RemoveKubernetesOK struct {
	Payload *models.V1Kubernetes
}

func (o *RemoveKubernetesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/kubernetes/{name}][%d] removeKubernetesOK  %+v", 200, o.Payload)
}

func (o *RemoveKubernetesOK) GetPayload() *models.V1Kubernetes {
	return o.Payload
}

func (o *RemoveKubernetesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Kubernetes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveKubernetesInternalServerError creates a RemoveKubernetesInternalServerError with default headers values
func NewRemoveKubernetesInternalServerError() *RemoveKubernetesInternalServerError {
	return &RemoveKubernetesInternalServerError{}
}

/*RemoveKubernetesInternalServerError handles this case with default header values.

A generic API error containing the cause of the error
*/
type RemoveKubernetesInternalServerError struct {
	Payload *models.ApiserverError
}

func (o *RemoveKubernetesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1alpha1/teams/{team}/kubernetes/{name}][%d] removeKubernetesInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveKubernetesInternalServerError) GetPayload() *models.ApiserverError {
	return o.Payload
}

func (o *RemoveKubernetesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}