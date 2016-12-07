package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/iron-io/functions_go/models"
)

// GetTasksReader is a Reader for the GetTasks structure.
type GetTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetTasksOK creates a GetTasksOK with default headers values
func NewGetTasksOK() *GetTasksOK {
	return &GetTasksOK{}
}

/*GetTasksOK handles this case with default header values.

Task information
*/
type GetTasksOK struct {
	Payload *models.TaskWrapper
}

func (o *GetTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] getTasksOK  %+v", 200, o.Payload)
}

func (o *GetTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTasksDefault creates a GetTasksDefault with default headers values
func NewGetTasksDefault(code int) *GetTasksDefault {
	return &GetTasksDefault{
		_statusCode: code,
	}
}

/*GetTasksDefault handles this case with default header values.

Unexpected error
*/
type GetTasksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tasks default response
func (o *GetTasksDefault) Code() int {
	return o._statusCode
}

func (o *GetTasksDefault) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] GetTasks default  %+v", o._statusCode, o.Payload)
}

func (o *GetTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}