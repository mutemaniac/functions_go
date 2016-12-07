package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/iron-io/functions_go/models"
)

// PostAppsReader is a Reader for the PostApps structure.
type PostAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAppsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAppsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostAppsOK creates a PostAppsOK with default headers values
func NewPostAppsOK() *PostAppsOK {
	return &PostAppsOK{}
}

/*PostAppsOK handles this case with default header values.

App details and stats.
*/
type PostAppsOK struct {
	Payload *models.AppWrapper
}

func (o *PostAppsOK) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsOK  %+v", 200, o.Payload)
}

func (o *PostAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsBadRequest creates a PostAppsBadRequest with default headers values
func NewPostAppsBadRequest() *PostAppsBadRequest {
	return &PostAppsBadRequest{}
}

/*PostAppsBadRequest handles this case with default header values.

Parameters are missing or invalid.
*/
type PostAppsBadRequest struct {
	Payload *models.Error
}

func (o *PostAppsBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAppsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsInternalServerError creates a PostAppsInternalServerError with default headers values
func NewPostAppsInternalServerError() *PostAppsInternalServerError {
	return &PostAppsInternalServerError{}
}

/*PostAppsInternalServerError handles this case with default header values.

Could not accept app due to internal error.
*/
type PostAppsInternalServerError struct {
	Payload *models.Error
}

func (o *PostAppsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps][%d] postAppsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAppsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAppsDefault creates a PostAppsDefault with default headers values
func NewPostAppsDefault(code int) *PostAppsDefault {
	return &PostAppsDefault{
		_statusCode: code,
	}
}

/*PostAppsDefault handles this case with default header values.

Unexpected error
*/
type PostAppsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post apps default response
func (o *PostAppsDefault) Code() int {
	return o._statusCode
}

func (o *PostAppsDefault) Error() string {
	return fmt.Sprintf("[POST /apps][%d] PostApps default  %+v", o._statusCode, o.Payload)
}

func (o *PostAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}