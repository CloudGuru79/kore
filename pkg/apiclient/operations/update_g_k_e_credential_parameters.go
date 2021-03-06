// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/appvia/kore/pkg/apiclient/models"
)

// NewUpdateGKECredentialParams creates a new UpdateGKECredentialParams object
// with the default values initialized.
func NewUpdateGKECredentialParams() *UpdateGKECredentialParams {
	var ()
	return &UpdateGKECredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGKECredentialParamsWithTimeout creates a new UpdateGKECredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGKECredentialParamsWithTimeout(timeout time.Duration) *UpdateGKECredentialParams {
	var ()
	return &UpdateGKECredentialParams{

		timeout: timeout,
	}
}

// NewUpdateGKECredentialParamsWithContext creates a new UpdateGKECredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGKECredentialParamsWithContext(ctx context.Context) *UpdateGKECredentialParams {
	var ()
	return &UpdateGKECredentialParams{

		Context: ctx,
	}
}

// NewUpdateGKECredentialParamsWithHTTPClient creates a new UpdateGKECredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGKECredentialParamsWithHTTPClient(client *http.Client) *UpdateGKECredentialParams {
	var ()
	return &UpdateGKECredentialParams{
		HTTPClient: client,
	}
}

/*UpdateGKECredentialParams contains all the parameters to send to the API endpoint
for the update g k e credential operation typically these are written to a http.Request
*/
type UpdateGKECredentialParams struct {

	/*Body
	  The definition for GKE Credentials

	*/
	Body *models.V1alpha1GKECredentials
	/*Name
	  Is name the of the GKE cluster you are acting upon

	*/
	Name string
	/*Team
	  Is the name of the team you are acting within

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update g k e credential params
func (o *UpdateGKECredentialParams) WithTimeout(timeout time.Duration) *UpdateGKECredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update g k e credential params
func (o *UpdateGKECredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update g k e credential params
func (o *UpdateGKECredentialParams) WithContext(ctx context.Context) *UpdateGKECredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update g k e credential params
func (o *UpdateGKECredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update g k e credential params
func (o *UpdateGKECredentialParams) WithHTTPClient(client *http.Client) *UpdateGKECredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update g k e credential params
func (o *UpdateGKECredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update g k e credential params
func (o *UpdateGKECredentialParams) WithBody(body *models.V1alpha1GKECredentials) *UpdateGKECredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update g k e credential params
func (o *UpdateGKECredentialParams) SetBody(body *models.V1alpha1GKECredentials) {
	o.Body = body
}

// WithName adds the name to the update g k e credential params
func (o *UpdateGKECredentialParams) WithName(name string) *UpdateGKECredentialParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update g k e credential params
func (o *UpdateGKECredentialParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the update g k e credential params
func (o *UpdateGKECredentialParams) WithTeam(team string) *UpdateGKECredentialParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the update g k e credential params
func (o *UpdateGKECredentialParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGKECredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
