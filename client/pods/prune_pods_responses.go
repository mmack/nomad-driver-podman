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

// PrunePodsReader is a Reader for the PrunePods structure.
type PrunePodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrunePodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrunePodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrunePodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPrunePodsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrunePodsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPrunePodsOK creates a PrunePodsOK with default headers values
func NewPrunePodsOK() *PrunePodsOK {
	return &PrunePodsOK{}
}

/*PrunePodsOK handles this case with default header values.

Prune pod
*/
type PrunePodsOK struct {
	Payload *models.PodPruneReport
}

func (o *PrunePodsOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] prunePodsOK  %+v", 200, o.Payload)
}

func (o *PrunePodsOK) GetPayload() *models.PodPruneReport {
	return o.Payload
}

func (o *PrunePodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodPruneReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrunePodsBadRequest creates a PrunePodsBadRequest with default headers values
func NewPrunePodsBadRequest() *PrunePodsBadRequest {
	return &PrunePodsBadRequest{}
}

/*PrunePodsBadRequest handles this case with default header values.

Bad parameter in request
*/
type PrunePodsBadRequest struct {
	Payload *PrunePodsBadRequestBody
}

func (o *PrunePodsBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] prunePodsBadRequest  %+v", 400, o.Payload)
}

func (o *PrunePodsBadRequest) GetPayload() *PrunePodsBadRequestBody {
	return o.Payload
}

func (o *PrunePodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PrunePodsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrunePodsConflict creates a PrunePodsConflict with default headers values
func NewPrunePodsConflict() *PrunePodsConflict {
	return &PrunePodsConflict{}
}

/*PrunePodsConflict handles this case with default header values.

pod already exists
*/
type PrunePodsConflict struct {
}

func (o *PrunePodsConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] prunePodsConflict ", 409)
}

func (o *PrunePodsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrunePodsInternalServerError creates a PrunePodsInternalServerError with default headers values
func NewPrunePodsInternalServerError() *PrunePodsInternalServerError {
	return &PrunePodsInternalServerError{}
}

/*PrunePodsInternalServerError handles this case with default header values.

Internal server error
*/
type PrunePodsInternalServerError struct {
	Payload *PrunePodsInternalServerErrorBody
}

func (o *PrunePodsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/prune][%d] prunePodsInternalServerError  %+v", 500, o.Payload)
}

func (o *PrunePodsInternalServerError) GetPayload() *PrunePodsInternalServerErrorBody {
	return o.Payload
}

func (o *PrunePodsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PrunePodsInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PrunePodsBadRequestBody prune pods bad request body
swagger:model PrunePodsBadRequestBody
*/
type PrunePodsBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this prune pods bad request body
func (o *PrunePodsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PrunePodsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PrunePodsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PrunePodsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PrunePodsInternalServerErrorBody prune pods internal server error body
swagger:model PrunePodsInternalServerErrorBody
*/
type PrunePodsInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this prune pods internal server error body
func (o *PrunePodsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PrunePodsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PrunePodsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PrunePodsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}