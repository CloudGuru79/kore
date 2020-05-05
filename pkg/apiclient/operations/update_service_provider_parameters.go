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

// NewUpdateServiceProviderParams creates a new UpdateServiceProviderParams object
// with the default values initialized.
func NewUpdateServiceProviderParams() *UpdateServiceProviderParams {
	var ()
	return &UpdateServiceProviderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceProviderParamsWithTimeout creates a new UpdateServiceProviderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceProviderParamsWithTimeout(timeout time.Duration) *UpdateServiceProviderParams {
	var ()
	return &UpdateServiceProviderParams{

		timeout: timeout,
	}
}

// NewUpdateServiceProviderParamsWithContext creates a new UpdateServiceProviderParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceProviderParamsWithContext(ctx context.Context) *UpdateServiceProviderParams {
	var ()
	return &UpdateServiceProviderParams{

		Context: ctx,
	}
}

// NewUpdateServiceProviderParamsWithHTTPClient creates a new UpdateServiceProviderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceProviderParamsWithHTTPClient(client *http.Client) *UpdateServiceProviderParams {
	var ()
	return &UpdateServiceProviderParams{
		HTTPClient: client,
	}
}

/*UpdateServiceProviderParams contains all the parameters to send to the API endpoint
for the update service provider operation typically these are written to a http.Request
*/
type UpdateServiceProviderParams struct {

	/*Body
	  The specification for the service provider you are creating or updating

	*/
	Body *models.V1ServiceProvider
	/*Name
	  The name of the service provider you wish to create or update

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service provider params
func (o *UpdateServiceProviderParams) WithTimeout(timeout time.Duration) *UpdateServiceProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service provider params
func (o *UpdateServiceProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service provider params
func (o *UpdateServiceProviderParams) WithContext(ctx context.Context) *UpdateServiceProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service provider params
func (o *UpdateServiceProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service provider params
func (o *UpdateServiceProviderParams) WithHTTPClient(client *http.Client) *UpdateServiceProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service provider params
func (o *UpdateServiceProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service provider params
func (o *UpdateServiceProviderParams) WithBody(body *models.V1ServiceProvider) *UpdateServiceProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service provider params
func (o *UpdateServiceProviderParams) SetBody(body *models.V1ServiceProvider) {
	o.Body = body
}

// WithName adds the name to the update service provider params
func (o *UpdateServiceProviderParams) WithName(name string) *UpdateServiceProviderParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update service provider params
func (o *UpdateServiceProviderParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}