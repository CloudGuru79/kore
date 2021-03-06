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
)

// NewGetGKEParams creates a new GetGKEParams object
// with the default values initialized.
func NewGetGKEParams() *GetGKEParams {
	var ()
	return &GetGKEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGKEParamsWithTimeout creates a new GetGKEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGKEParamsWithTimeout(timeout time.Duration) *GetGKEParams {
	var ()
	return &GetGKEParams{

		timeout: timeout,
	}
}

// NewGetGKEParamsWithContext creates a new GetGKEParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGKEParamsWithContext(ctx context.Context) *GetGKEParams {
	var ()
	return &GetGKEParams{

		Context: ctx,
	}
}

// NewGetGKEParamsWithHTTPClient creates a new GetGKEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGKEParamsWithHTTPClient(client *http.Client) *GetGKEParams {
	var ()
	return &GetGKEParams{
		HTTPClient: client,
	}
}

/*GetGKEParams contains all the parameters to send to the API endpoint
for the get g k e operation typically these are written to a http.Request
*/
type GetGKEParams struct {

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

// WithTimeout adds the timeout to the get g k e params
func (o *GetGKEParams) WithTimeout(timeout time.Duration) *GetGKEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get g k e params
func (o *GetGKEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get g k e params
func (o *GetGKEParams) WithContext(ctx context.Context) *GetGKEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get g k e params
func (o *GetGKEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get g k e params
func (o *GetGKEParams) WithHTTPClient(client *http.Client) *GetGKEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get g k e params
func (o *GetGKEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get g k e params
func (o *GetGKEParams) WithName(name string) *GetGKEParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get g k e params
func (o *GetGKEParams) SetName(name string) {
	o.Name = name
}

// WithTeam adds the team to the get g k e params
func (o *GetGKEParams) WithTeam(team string) *GetGKEParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get g k e params
func (o *GetGKEParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetGKEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
