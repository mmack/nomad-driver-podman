// Code generated by go-swagger; DO NOT EDIT.

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodResizeExecReader is a Reader for the LibpodResizeExec structure.
type LibpodResizeExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LibpodResizeExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLibpodResizeExecCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodResizeExecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodResizeExecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodResizeExecCreated creates a LibpodResizeExecCreated with default headers values
func NewLibpodResizeExecCreated() *LibpodResizeExecCreated {
	return &LibpodResizeExecCreated{}
}

/*LibpodResizeExecCreated handles this case with default header values.

no error
*/
type LibpodResizeExecCreated struct {
}

func (o *LibpodResizeExecCreated) Error() string {
	return fmt.Sprintf("[POST /libpod/exec/{id}/resize][%d] libpodResizeExecCreated ", 201)
}

func (o *LibpodResizeExecCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLibpodResizeExecNotFound creates a LibpodResizeExecNotFound with default headers values
func NewLibpodResizeExecNotFound() *LibpodResizeExecNotFound {
	return &LibpodResizeExecNotFound{}
}

/*LibpodResizeExecNotFound handles this case with default header values.

No such exec instance
*/
type LibpodResizeExecNotFound struct {
	Payload *LibpodResizeExecNotFoundBody
}

func (o *LibpodResizeExecNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/exec/{id}/resize][%d] libpodResizeExecNotFound  %+v", 404, o.Payload)
}

func (o *LibpodResizeExecNotFound) GetPayload() *LibpodResizeExecNotFoundBody {
	return o.Payload
}

func (o *LibpodResizeExecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodResizeExecNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodResizeExecInternalServerError creates a LibpodResizeExecInternalServerError with default headers values
func NewLibpodResizeExecInternalServerError() *LibpodResizeExecInternalServerError {
	return &LibpodResizeExecInternalServerError{}
}

/*LibpodResizeExecInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodResizeExecInternalServerError struct {
	Payload *LibpodResizeExecInternalServerErrorBody
}

func (o *LibpodResizeExecInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/exec/{id}/resize][%d] libpodResizeExecInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodResizeExecInternalServerError) GetPayload() *LibpodResizeExecInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodResizeExecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodResizeExecInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodResizeExecInternalServerErrorBody libpod resize exec internal server error body
swagger:model LibpodResizeExecInternalServerErrorBody
*/
type LibpodResizeExecInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod resize exec internal server error body
func (o *LibpodResizeExecInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodResizeExecInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodResizeExecInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodResizeExecInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodResizeExecNotFoundBody libpod resize exec not found body
swagger:model LibpodResizeExecNotFoundBody
*/
type LibpodResizeExecNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod resize exec not found body
func (o *LibpodResizeExecNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodResizeExecNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodResizeExecNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodResizeExecNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}