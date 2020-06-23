// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

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

// NewCompatInspectNetworkParams creates a new CompatInspectNetworkParams object
// with the default values initialized.
func NewCompatInspectNetworkParams() *CompatInspectNetworkParams {
	var ()
	return &CompatInspectNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompatInspectNetworkParamsWithTimeout creates a new CompatInspectNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompatInspectNetworkParamsWithTimeout(timeout time.Duration) *CompatInspectNetworkParams {
	var ()
	return &CompatInspectNetworkParams{

		timeout: timeout,
	}
}

// NewCompatInspectNetworkParamsWithContext creates a new CompatInspectNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompatInspectNetworkParamsWithContext(ctx context.Context) *CompatInspectNetworkParams {
	var ()
	return &CompatInspectNetworkParams{

		Context: ctx,
	}
}

// NewCompatInspectNetworkParamsWithHTTPClient creates a new CompatInspectNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompatInspectNetworkParamsWithHTTPClient(client *http.Client) *CompatInspectNetworkParams {
	var ()
	return &CompatInspectNetworkParams{
		HTTPClient: client,
	}
}

/*CompatInspectNetworkParams contains all the parameters to send to the API endpoint
for the compat inspect network operation typically these are written to a http.Request
*/
type CompatInspectNetworkParams struct {

	/*Name
	  the name of the network

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the compat inspect network params
func (o *CompatInspectNetworkParams) WithTimeout(timeout time.Duration) *CompatInspectNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the compat inspect network params
func (o *CompatInspectNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the compat inspect network params
func (o *CompatInspectNetworkParams) WithContext(ctx context.Context) *CompatInspectNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the compat inspect network params
func (o *CompatInspectNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the compat inspect network params
func (o *CompatInspectNetworkParams) WithHTTPClient(client *http.Client) *CompatInspectNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the compat inspect network params
func (o *CompatInspectNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the compat inspect network params
func (o *CompatInspectNetworkParams) WithName(name string) *CompatInspectNetworkParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the compat inspect network params
func (o *CompatInspectNetworkParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *CompatInspectNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}