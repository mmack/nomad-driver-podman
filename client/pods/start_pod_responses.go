// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pascomnet/nomad-driver-podman/models"
)

// StartPodReader is a Reader for the StartPod structure.
type StartPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStartPodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartPodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartPodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStartPodOK creates a StartPodOK with default headers values
func NewStartPodOK() *StartPodOK {
	return &StartPodOK{}
}

/*StartPodOK handles this case with default header values.

Start pod
*/
type StartPodOK struct {
	Payload *models.PodStartReport
}

func (o *StartPodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/start][%d] startPodOK  %+v", 200, o.Payload)
}

func (o *StartPodOK) GetPayload() *models.PodStartReport {
	return o.Payload
}

func (o *StartPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodStartReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPodNotModified creates a StartPodNotModified with default headers values
func NewStartPodNotModified() *StartPodNotModified {
	return &StartPodNotModified{}
}

/*StartPodNotModified handles this case with default header values.

Pod already started
*/
type StartPodNotModified struct {
	Payload *StartPodNotModifiedBody
}

func (o *StartPodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/start][%d] startPodNotModified  %+v", 304, o.Payload)
}

func (o *StartPodNotModified) GetPayload() *StartPodNotModifiedBody {
	return o.Payload
}

func (o *StartPodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPodNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPodNotFound creates a StartPodNotFound with default headers values
func NewStartPodNotFound() *StartPodNotFound {
	return &StartPodNotFound{}
}

/*StartPodNotFound handles this case with default header values.

No such pod
*/
type StartPodNotFound struct {
	Payload *StartPodNotFoundBody
}

func (o *StartPodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/start][%d] startPodNotFound  %+v", 404, o.Payload)
}

func (o *StartPodNotFound) GetPayload() *StartPodNotFoundBody {
	return o.Payload
}

func (o *StartPodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartPodInternalServerError creates a StartPodInternalServerError with default headers values
func NewStartPodInternalServerError() *StartPodInternalServerError {
	return &StartPodInternalServerError{}
}

/*StartPodInternalServerError handles this case with default header values.

Internal server error
*/
type StartPodInternalServerError struct {
	Payload *StartPodInternalServerErrorBody
}

func (o *StartPodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/start][%d] startPodInternalServerError  %+v", 500, o.Payload)
}

func (o *StartPodInternalServerError) GetPayload() *StartPodInternalServerErrorBody {
	return o.Payload
}

func (o *StartPodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartPodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartPodInternalServerErrorBody start pod internal server error body
swagger:model StartPodInternalServerErrorBody
*/
type StartPodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this start pod internal server error body
func (o *StartPodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res StartPodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPodNotFoundBody start pod not found body
swagger:model StartPodNotFoundBody
*/
type StartPodNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this start pod not found body
func (o *StartPodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StartPodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartPodNotModifiedBody start pod not modified body
swagger:model StartPodNotModifiedBody
*/
type StartPodNotModifiedBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this start pod not modified body
func (o *StartPodNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartPodNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartPodNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res StartPodNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}