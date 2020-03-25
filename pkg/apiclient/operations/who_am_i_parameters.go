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

// NewWhoAmIParams creates a new WhoAmIParams object
// with the default values initialized.
func NewWhoAmIParams() *WhoAmIParams {

	return &WhoAmIParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWhoAmIParamsWithTimeout creates a new WhoAmIParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWhoAmIParamsWithTimeout(timeout time.Duration) *WhoAmIParams {

	return &WhoAmIParams{

		timeout: timeout,
	}
}

// NewWhoAmIParamsWithContext creates a new WhoAmIParams object
// with the default values initialized, and the ability to set a context for a request
func NewWhoAmIParamsWithContext(ctx context.Context) *WhoAmIParams {

	return &WhoAmIParams{

		Context: ctx,
	}
}

// NewWhoAmIParamsWithHTTPClient creates a new WhoAmIParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWhoAmIParamsWithHTTPClient(client *http.Client) *WhoAmIParams {

	return &WhoAmIParams{
		HTTPClient: client,
	}
}

/*WhoAmIParams contains all the parameters to send to the API endpoint
for the who am i operation typically these are written to a http.Request
*/
type WhoAmIParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the who am i params
func (o *WhoAmIParams) WithTimeout(timeout time.Duration) *WhoAmIParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the who am i params
func (o *WhoAmIParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the who am i params
func (o *WhoAmIParams) WithContext(ctx context.Context) *WhoAmIParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the who am i params
func (o *WhoAmIParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the who am i params
func (o *WhoAmIParams) WithHTTPClient(client *http.Client) *WhoAmIParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the who am i params
func (o *WhoAmIParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WhoAmIParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}