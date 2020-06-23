// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// NewLogsFromContainerParams creates a new LogsFromContainerParams object
// with the default values initialized.
func NewLogsFromContainerParams() *LogsFromContainerParams {
	var (
		tailDefault       = string("all")
		timestampsDefault = bool(false)
	)
	return &LogsFromContainerParams{
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewLogsFromContainerParamsWithTimeout creates a new LogsFromContainerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogsFromContainerParamsWithTimeout(timeout time.Duration) *LogsFromContainerParams {
	var (
		tailDefault       = string("all")
		timestampsDefault = bool(false)
	)
	return &LogsFromContainerParams{
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,

		timeout: timeout,
	}
}

// NewLogsFromContainerParamsWithContext creates a new LogsFromContainerParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogsFromContainerParamsWithContext(ctx context.Context) *LogsFromContainerParams {
	var (
		tailDefault       = string("all")
		timestampsDefault = bool(false)
	)
	return &LogsFromContainerParams{
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,

		Context: ctx,
	}
}

// NewLogsFromContainerParamsWithHTTPClient creates a new LogsFromContainerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogsFromContainerParamsWithHTTPClient(client *http.Client) *LogsFromContainerParams {
	var (
		tailDefault       = string("all")
		timestampsDefault = bool(false)
	)
	return &LogsFromContainerParams{
		Tail:       &tailDefault,
		Timestamps: &timestampsDefault,
		HTTPClient: client,
	}
}

/*LogsFromContainerParams contains all the parameters to send to the API endpoint
for the logs from container operation typically these are written to a http.Request
*/
type LogsFromContainerParams struct {

	/*Follow
	  Keep connection after returning logs.

	*/
	Follow *bool
	/*Name
	  the name or ID of the container

	*/
	Name string
	/*Since
	  Only return logs since this time, as a UNIX timestamp

	*/
	Since *string
	/*Stderr
	  Return logs from stderr

	*/
	Stderr *bool
	/*Stdout
	  Return logs from stdout

	*/
	Stdout *bool
	/*Tail
	  Only return this number of log lines from the end of the logs

	*/
	Tail *string
	/*Timestamps
	  Add timestamps to every log line

	*/
	Timestamps *bool
	/*Until
	  Only return logs before this time, as a UNIX timestamp

	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logs from container params
func (o *LogsFromContainerParams) WithTimeout(timeout time.Duration) *LogsFromContainerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logs from container params
func (o *LogsFromContainerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logs from container params
func (o *LogsFromContainerParams) WithContext(ctx context.Context) *LogsFromContainerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logs from container params
func (o *LogsFromContainerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logs from container params
func (o *LogsFromContainerParams) WithHTTPClient(client *http.Client) *LogsFromContainerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logs from container params
func (o *LogsFromContainerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFollow adds the follow to the logs from container params
func (o *LogsFromContainerParams) WithFollow(follow *bool) *LogsFromContainerParams {
	o.SetFollow(follow)
	return o
}

// SetFollow adds the follow to the logs from container params
func (o *LogsFromContainerParams) SetFollow(follow *bool) {
	o.Follow = follow
}

// WithName adds the name to the logs from container params
func (o *LogsFromContainerParams) WithName(name string) *LogsFromContainerParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the logs from container params
func (o *LogsFromContainerParams) SetName(name string) {
	o.Name = name
}

// WithSince adds the since to the logs from container params
func (o *LogsFromContainerParams) WithSince(since *string) *LogsFromContainerParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the logs from container params
func (o *LogsFromContainerParams) SetSince(since *string) {
	o.Since = since
}

// WithStderr adds the stderr to the logs from container params
func (o *LogsFromContainerParams) WithStderr(stderr *bool) *LogsFromContainerParams {
	o.SetStderr(stderr)
	return o
}

// SetStderr adds the stderr to the logs from container params
func (o *LogsFromContainerParams) SetStderr(stderr *bool) {
	o.Stderr = stderr
}

// WithStdout adds the stdout to the logs from container params
func (o *LogsFromContainerParams) WithStdout(stdout *bool) *LogsFromContainerParams {
	o.SetStdout(stdout)
	return o
}

// SetStdout adds the stdout to the logs from container params
func (o *LogsFromContainerParams) SetStdout(stdout *bool) {
	o.Stdout = stdout
}

// WithTail adds the tail to the logs from container params
func (o *LogsFromContainerParams) WithTail(tail *string) *LogsFromContainerParams {
	o.SetTail(tail)
	return o
}

// SetTail adds the tail to the logs from container params
func (o *LogsFromContainerParams) SetTail(tail *string) {
	o.Tail = tail
}

// WithTimestamps adds the timestamps to the logs from container params
func (o *LogsFromContainerParams) WithTimestamps(timestamps *bool) *LogsFromContainerParams {
	o.SetTimestamps(timestamps)
	return o
}

// SetTimestamps adds the timestamps to the logs from container params
func (o *LogsFromContainerParams) SetTimestamps(timestamps *bool) {
	o.Timestamps = timestamps
}

// WithUntil adds the until to the logs from container params
func (o *LogsFromContainerParams) WithUntil(until *string) *LogsFromContainerParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the logs from container params
func (o *LogsFromContainerParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *LogsFromContainerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Follow != nil {

		// query param follow
		var qrFollow bool
		if o.Follow != nil {
			qrFollow = *o.Follow
		}
		qFollow := swag.FormatBool(qrFollow)
		if qFollow != "" {
			if err := r.SetQueryParam("follow", qFollow); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.Stderr != nil {

		// query param stderr
		var qrStderr bool
		if o.Stderr != nil {
			qrStderr = *o.Stderr
		}
		qStderr := swag.FormatBool(qrStderr)
		if qStderr != "" {
			if err := r.SetQueryParam("stderr", qStderr); err != nil {
				return err
			}
		}

	}

	if o.Stdout != nil {

		// query param stdout
		var qrStdout bool
		if o.Stdout != nil {
			qrStdout = *o.Stdout
		}
		qStdout := swag.FormatBool(qrStdout)
		if qStdout != "" {
			if err := r.SetQueryParam("stdout", qStdout); err != nil {
				return err
			}
		}

	}

	if o.Tail != nil {

		// query param tail
		var qrTail string
		if o.Tail != nil {
			qrTail = *o.Tail
		}
		qTail := qrTail
		if qTail != "" {
			if err := r.SetQueryParam("tail", qTail); err != nil {
				return err
			}
		}

	}

	if o.Timestamps != nil {

		// query param timestamps
		var qrTimestamps bool
		if o.Timestamps != nil {
			qrTimestamps = *o.Timestamps
		}
		qTimestamps := swag.FormatBool(qrTimestamps)
		if qTimestamps != "" {
			if err := r.SetQueryParam("timestamps", qTimestamps); err != nil {
				return err
			}
		}

	}

	if o.Until != nil {

		// query param until
		var qrUntil string
		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {
			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}