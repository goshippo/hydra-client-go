// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra-client-go/models"
)

// NewCreateJSONWebKeySetParams creates a new CreateJSONWebKeySetParams object
// with the default values initialized.
func NewCreateJSONWebKeySetParams() *CreateJSONWebKeySetParams {
	var ()
	return &CreateJSONWebKeySetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateJSONWebKeySetParamsWithTimeout creates a new CreateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateJSONWebKeySetParamsWithTimeout(timeout time.Duration) *CreateJSONWebKeySetParams {
	var ()
	return &CreateJSONWebKeySetParams{

		timeout: timeout,
	}
}

// NewCreateJSONWebKeySetParamsWithContext creates a new CreateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateJSONWebKeySetParamsWithContext(ctx context.Context) *CreateJSONWebKeySetParams {
	var ()
	return &CreateJSONWebKeySetParams{

		Context: ctx,
	}
}

// NewCreateJSONWebKeySetParamsWithHTTPClient creates a new CreateJSONWebKeySetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateJSONWebKeySetParamsWithHTTPClient(client *http.Client) *CreateJSONWebKeySetParams {
	var ()
	return &CreateJSONWebKeySetParams{
		HTTPClient: client,
	}
}

/*CreateJSONWebKeySetParams contains all the parameters to send to the API endpoint
for the create Json web key set operation typically these are written to a http.Request
*/
type CreateJSONWebKeySetParams struct {

	/*Body*/
	Body *models.JSONWebKeySetGeneratorRequest
	/*Set
	  The set

	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create Json web key set params
func (o *CreateJSONWebKeySetParams) WithTimeout(timeout time.Duration) *CreateJSONWebKeySetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Json web key set params
func (o *CreateJSONWebKeySetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Json web key set params
func (o *CreateJSONWebKeySetParams) WithContext(ctx context.Context) *CreateJSONWebKeySetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Json web key set params
func (o *CreateJSONWebKeySetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Json web key set params
func (o *CreateJSONWebKeySetParams) WithHTTPClient(client *http.Client) *CreateJSONWebKeySetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Json web key set params
func (o *CreateJSONWebKeySetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create Json web key set params
func (o *CreateJSONWebKeySetParams) WithBody(body *models.JSONWebKeySetGeneratorRequest) *CreateJSONWebKeySetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create Json web key set params
func (o *CreateJSONWebKeySetParams) SetBody(body *models.JSONWebKeySetGeneratorRequest) {
	o.Body = body
}

// WithSet adds the set to the create Json web key set params
func (o *CreateJSONWebKeySetParams) WithSet(set string) *CreateJSONWebKeySetParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the create Json web key set params
func (o *CreateJSONWebKeySetParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *CreateJSONWebKeySetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
