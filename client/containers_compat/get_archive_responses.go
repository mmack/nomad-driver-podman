// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetArchiveReader is a Reader for the GetArchive structure.
type GetArchiveReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchiveOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetArchiveOK creates a GetArchiveOK with default headers values
func NewGetArchiveOK(writer io.Writer) *GetArchiveOK {
	return &GetArchiveOK{
		Payload: writer,
	}
}

/*GetArchiveOK handles this case with default header values.

no error
*/
type GetArchiveOK struct {
	Payload io.Writer
}

func (o *GetArchiveOK) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/archive][%d] getArchiveOK  %+v", 200, o.Payload)
}

func (o *GetArchiveOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchiveBadRequest creates a GetArchiveBadRequest with default headers values
func NewGetArchiveBadRequest() *GetArchiveBadRequest {
	return &GetArchiveBadRequest{}
}

/*GetArchiveBadRequest handles this case with default header values.

Bad parameter in request
*/
type GetArchiveBadRequest struct {
	Payload *GetArchiveBadRequestBody
}

func (o *GetArchiveBadRequest) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/archive][%d] getArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchiveBadRequest) GetPayload() *GetArchiveBadRequestBody {
	return o.Payload
}

func (o *GetArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetArchiveBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchiveNotFound creates a GetArchiveNotFound with default headers values
func NewGetArchiveNotFound() *GetArchiveNotFound {
	return &GetArchiveNotFound{}
}

/*GetArchiveNotFound handles this case with default header values.

No such container
*/
type GetArchiveNotFound struct {
	Payload *GetArchiveNotFoundBody
}

func (o *GetArchiveNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/archive][%d] getArchiveNotFound  %+v", 404, o.Payload)
}

func (o *GetArchiveNotFound) GetPayload() *GetArchiveNotFoundBody {
	return o.Payload
}

func (o *GetArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetArchiveNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchiveInternalServerError creates a GetArchiveInternalServerError with default headers values
func NewGetArchiveInternalServerError() *GetArchiveInternalServerError {
	return &GetArchiveInternalServerError{}
}

/*GetArchiveInternalServerError handles this case with default header values.

Internal server error
*/
type GetArchiveInternalServerError struct {
	Payload *GetArchiveInternalServerErrorBody
}

func (o *GetArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/archive][%d] getArchiveInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchiveInternalServerError) GetPayload() *GetArchiveInternalServerErrorBody {
	return o.Payload
}

func (o *GetArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetArchiveInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetArchiveBadRequestBody get archive bad request body
swagger:model GetArchiveBadRequestBody
*/
type GetArchiveBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this get archive bad request body
func (o *GetArchiveBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetArchiveBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetArchiveBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetArchiveBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetArchiveInternalServerErrorBody get archive internal server error body
swagger:model GetArchiveInternalServerErrorBody
*/
type GetArchiveInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this get archive internal server error body
func (o *GetArchiveInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetArchiveInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetArchiveInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetArchiveInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetArchiveNotFoundBody get archive not found body
swagger:model GetArchiveNotFoundBody
*/
type GetArchiveNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this get archive not found body
func (o *GetArchiveNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetArchiveNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetArchiveNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetArchiveNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}