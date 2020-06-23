// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// NewPruneImagesParams creates a new PruneImagesParams object
// with the default values initialized.
func NewPruneImagesParams() *PruneImagesParams {
	var ()
	return &PruneImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPruneImagesParamsWithTimeout creates a new PruneImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPruneImagesParamsWithTimeout(timeout time.Duration) *PruneImagesParams {
	var ()
	return &PruneImagesParams{

		timeout: timeout,
	}
}

// NewPruneImagesParamsWithContext creates a new PruneImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPruneImagesParamsWithContext(ctx context.Context) *PruneImagesParams {
	var ()
	return &PruneImagesParams{

		Context: ctx,
	}
}

// NewPruneImagesParamsWithHTTPClient creates a new PruneImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPruneImagesParamsWithHTTPClient(client *http.Client) *PruneImagesParams {
	var ()
	return &PruneImagesParams{
		HTTPClient: client,
	}
}

/*PruneImagesParams contains all the parameters to send to the API endpoint
for the prune images operation typically these are written to a http.Request
*/
type PruneImagesParams struct {

	/*Filters
	  filters to apply to image pruning, encoded as JSON (map[string][]string). Available filters:
	  - `dangling=<boolean>` When set to `true` (or `1`), prune only
	     unused *and* untagged images. When set to `false`
	     (or `0`), all unused images are pruned.
	  - `until=<string>` Prune images created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	  - `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune images with (or without, in case `label!=...` is used) the specified labels.


	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prune images params
func (o *PruneImagesParams) WithTimeout(timeout time.Duration) *PruneImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prune images params
func (o *PruneImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prune images params
func (o *PruneImagesParams) WithContext(ctx context.Context) *PruneImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prune images params
func (o *PruneImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prune images params
func (o *PruneImagesParams) WithHTTPClient(client *http.Client) *PruneImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prune images params
func (o *PruneImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the prune images params
func (o *PruneImagesParams) WithFilters(filters *string) *PruneImagesParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the prune images params
func (o *PruneImagesParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *PruneImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}