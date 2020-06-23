// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// NewPodExistsParams creates a new PodExistsParams object
// with the default values initialized.
func NewPodExistsParams() *PodExistsParams {
	var ()
	return &PodExistsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPodExistsParamsWithTimeout creates a new PodExistsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPodExistsParamsWithTimeout(timeout time.Duration) *PodExistsParams {
	var ()
	return &PodExistsParams{

		timeout: timeout,
	}
}

// NewPodExistsParamsWithContext creates a new PodExistsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPodExistsParamsWithContext(ctx context.Context) *PodExistsParams {
	var ()
	return &PodExistsParams{

		Context: ctx,
	}
}

// NewPodExistsParamsWithHTTPClient creates a new PodExistsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPodExistsParamsWithHTTPClient(client *http.Client) *PodExistsParams {
	var ()
	return &PodExistsParams{
		HTTPClient: client,
	}
}

/*PodExistsParams contains all the parameters to send to the API endpoint
for the pod exists operation typically these are written to a http.Request
*/
type PodExistsParams struct {

	/*Name
	  the name or ID of the pod

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pod exists params
func (o *PodExistsParams) WithTimeout(timeout time.Duration) *PodExistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod exists params
func (o *PodExistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod exists params
func (o *PodExistsParams) WithContext(ctx context.Context) *PodExistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod exists params
func (o *PodExistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod exists params
func (o *PodExistsParams) WithHTTPClient(client *http.Client) *PodExistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod exists params
func (o *PodExistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod exists params
func (o *PodExistsParams) WithName(name string) *PodExistsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod exists params
func (o *PodExistsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodExistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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