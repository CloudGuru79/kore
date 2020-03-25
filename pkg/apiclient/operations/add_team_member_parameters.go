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

// NewAddTeamMemberParams creates a new AddTeamMemberParams object
// with the default values initialized.
func NewAddTeamMemberParams() *AddTeamMemberParams {
	var ()
	return &AddTeamMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddTeamMemberParamsWithTimeout creates a new AddTeamMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddTeamMemberParamsWithTimeout(timeout time.Duration) *AddTeamMemberParams {
	var ()
	return &AddTeamMemberParams{

		timeout: timeout,
	}
}

// NewAddTeamMemberParamsWithContext creates a new AddTeamMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddTeamMemberParamsWithContext(ctx context.Context) *AddTeamMemberParams {
	var ()
	return &AddTeamMemberParams{

		Context: ctx,
	}
}

// NewAddTeamMemberParamsWithHTTPClient creates a new AddTeamMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddTeamMemberParamsWithHTTPClient(client *http.Client) *AddTeamMemberParams {
	var ()
	return &AddTeamMemberParams{
		HTTPClient: client,
	}
}

/*AddTeamMemberParams contains all the parameters to send to the API endpoint
for the add team member operation typically these are written to a http.Request
*/
type AddTeamMemberParams struct {

	/*Body
	  The definition for the user in the team

	*/
	Body *models.V1TeamMember
	/*Team
	  Is the name of the team you are acting within

	*/
	Team string
	/*User
	  Is the user you are adding to the team

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add team member params
func (o *AddTeamMemberParams) WithTimeout(timeout time.Duration) *AddTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add team member params
func (o *AddTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add team member params
func (o *AddTeamMemberParams) WithContext(ctx context.Context) *AddTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add team member params
func (o *AddTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add team member params
func (o *AddTeamMemberParams) WithHTTPClient(client *http.Client) *AddTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add team member params
func (o *AddTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add team member params
func (o *AddTeamMemberParams) WithBody(body *models.V1TeamMember) *AddTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add team member params
func (o *AddTeamMemberParams) SetBody(body *models.V1TeamMember) {
	o.Body = body
}

// WithTeam adds the team to the add team member params
func (o *AddTeamMemberParams) WithTeam(team string) *AddTeamMemberParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the add team member params
func (o *AddTeamMemberParams) SetTeam(team string) {
	o.Team = team
}

// WithUser adds the user to the add team member params
func (o *AddTeamMemberParams) WithUser(user string) *AddTeamMemberParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the add team member params
func (o *AddTeamMemberParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *AddTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}