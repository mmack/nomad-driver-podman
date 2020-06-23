// Code generated by go-swagger; DO NOT EDIT.

package exec_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new exec compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for exec compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateExec(params *CreateExecParams) (*CreateExecCreated, error)

	InspectExec(params *InspectExecParams) (*InspectExecOK, error)

	ResizeExec(params *ResizeExecParams) (*ResizeExecCreated, error)

	StartExec(params *StartExecParams) (*StartExecOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateExec creates an exec instance

  Create an exec session to run a command inside a running container. Exec sessions will be automatically removed 5 minutes after they exit.
*/
func (a *Client) CreateExec(params *CreateExecParams) (*CreateExecCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExecParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createExec",
		Method:             "POST",
		PathPattern:        "/containers/{name}/exec",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateExecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateExecCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InspectExec inspects an exec instance

  Return low-level information about an exec instance.
*/
func (a *Client) InspectExec(params *InspectExecParams) (*InspectExecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInspectExecParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "inspectExec",
		Method:             "GET",
		PathPattern:        "/exec/{id}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InspectExecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InspectExecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inspectExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResizeExec resizes an exec instance

  Resize the TTY session used by an exec instance. This endpoint only works if tty was specified as part of creating and starting the exec instance.

*/
func (a *Client) ResizeExec(params *ResizeExecParams) (*ResizeExecCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeExecParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resizeExec",
		Method:             "POST",
		PathPattern:        "/exec/{id}/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResizeExecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResizeExecCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resizeExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartExec starts an exec instance

  Starts a previously set up exec instance. If detach is true, this endpoint returns immediately after starting the command. Otherwise, it sets up an interactive session with the command.
*/
func (a *Client) StartExec(params *StartExecParams) (*StartExecOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartExecParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startExec",
		Method:             "POST",
		PathPattern:        "/exec/{id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartExecReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartExecOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startExec: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}