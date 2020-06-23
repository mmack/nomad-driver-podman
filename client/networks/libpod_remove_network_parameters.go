// Code generated by go-swagger; DO NOT EDIT.

package networks

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
	"github.com/go-openapi/swag"
)

// NewLibpodRemoveNetworkParams creates a new LibpodRemoveNetworkParams object
// with the default values initialized.
func NewLibpodRemoveNetworkParams() *LibpodRemoveNetworkParams {
	var ()
	return &LibpodRemoveNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodRemoveNetworkParamsWithTimeout creates a new LibpodRemoveNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodRemoveNetworkParamsWithTimeout(timeout time.Duration) *LibpodRemoveNetworkParams {
	var ()
	return &LibpodRemoveNetworkParams{

		timeout: timeout,
	}
}

// NewLibpodRemoveNetworkParamsWithContext creates a new LibpodRemoveNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodRemoveNetworkParamsWithContext(ctx context.Context) *LibpodRemoveNetworkParams {
	var ()
	return &LibpodRemoveNetworkParams{

		Context: ctx,
	}
}

// NewLibpodRemoveNetworkParamsWithHTTPClient creates a new LibpodRemoveNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodRemoveNetworkParamsWithHTTPClient(client *http.Client) *LibpodRemoveNetworkParams {
	var ()
	return &LibpodRemoveNetworkParams{
		HTTPClient: client,
	}
}

/*LibpodRemoveNetworkParams contains all the parameters to send to the API endpoint
for the libpod remove network operation typically these are written to a http.Request
*/
type LibpodRemoveNetworkParams struct {

	/*Force
	  remove containers associated with network

	*/
	Force *bool
	/*Name
	  the name of the network

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod remove network params
func (o *LibpodRemoveNetworkParams) WithTimeout(timeout time.Duration) *LibpodRemoveNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod remove network params
func (o *LibpodRemoveNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod remove network params
func (o *LibpodRemoveNetworkParams) WithContext(ctx context.Context) *LibpodRemoveNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod remove network params
func (o *LibpodRemoveNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod remove network params
func (o *LibpodRemoveNetworkParams) WithHTTPClient(client *http.Client) *LibpodRemoveNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod remove network params
func (o *LibpodRemoveNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the libpod remove network params
func (o *LibpodRemoveNetworkParams) WithForce(force *bool) *LibpodRemoveNetworkParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the libpod remove network params
func (o *LibpodRemoveNetworkParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the libpod remove network params
func (o *LibpodRemoveNetworkParams) WithName(name string) *LibpodRemoveNetworkParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the libpod remove network params
func (o *LibpodRemoveNetworkParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodRemoveNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param Force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("Force", qForce); err != nil {
				return err
			}
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