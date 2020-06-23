// Code generated by go-swagger; DO NOT EDIT.

package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodInitContainerReader is a Reader for the LibpodInitContainer structure.
type LibpodInitContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodInitContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLibpodInitContainerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewLibpodInitContainerNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLibpodInitContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodInitContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodInitContainerNoContent creates a LibpodInitContainerNoContent with default headers values
func NewLibpodInitContainerNoContent() *LibpodInitContainerNoContent {
	return &LibpodInitContainerNoContent{}
}

/*LibpodInitContainerNoContent handles this case with default header values.

no error
*/
type LibpodInitContainerNoContent struct {
}

func (o *LibpodInitContainerNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] libpodInitContainerNoContent ", 204)
}

func (o *LibpodInitContainerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodInitContainerNotModified creates a LibpodInitContainerNotModified with default headers values
func NewLibpodInitContainerNotModified() *LibpodInitContainerNotModified {
	return &LibpodInitContainerNotModified{}
}

/*LibpodInitContainerNotModified handles this case with default header values.

container already initialized
*/
type LibpodInitContainerNotModified struct {
}

func (o *LibpodInitContainerNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] libpodInitContainerNotModified ", 304)
}

func (o *LibpodInitContainerNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodInitContainerNotFound creates a LibpodInitContainerNotFound with default headers values
func NewLibpodInitContainerNotFound() *LibpodInitContainerNotFound {
	return &LibpodInitContainerNotFound{}
}

/*LibpodInitContainerNotFound handles this case with default header values.

No such container
*/
type LibpodInitContainerNotFound struct {
	Payload *LibpodInitContainerNotFoundBody
}

func (o *LibpodInitContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] libpodInitContainerNotFound  %+v", 404, o.Payload)
}

func (o *LibpodInitContainerNotFound) GetPayload() *LibpodInitContainerNotFoundBody {
	return o.Payload
}

func (o *LibpodInitContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodInitContainerNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodInitContainerInternalServerError creates a LibpodInitContainerInternalServerError with default headers values
func NewLibpodInitContainerInternalServerError() *LibpodInitContainerInternalServerError {
	return &LibpodInitContainerInternalServerError{}
}

/*LibpodInitContainerInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodInitContainerInternalServerError struct {
	Payload *LibpodInitContainerInternalServerErrorBody
}

func (o *LibpodInitContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] libpodInitContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodInitContainerInternalServerError) GetPayload() *LibpodInitContainerInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodInitContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodInitContainerInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodInitContainerInternalServerErrorBody libpod init container internal server error body
swagger:model LibpodInitContainerInternalServerErrorBody
*/
type LibpodInitContainerInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod init container internal server error body
func (o *LibpodInitContainerInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodInitContainerInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodInitContainerInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodInitContainerInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodInitContainerNotFoundBody libpod init container not found body
swagger:model LibpodInitContainerNotFoundBody
*/
type LibpodInitContainerNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod init container not found body
func (o *LibpodInitContainerNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodInitContainerNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodInitContainerNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodInitContainerNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}