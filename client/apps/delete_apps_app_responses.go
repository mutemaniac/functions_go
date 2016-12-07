package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAppsAppReader is a Reader for the DeleteAppsApp structure.
type DeleteAppsAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppsAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAppsAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAppsAppOK creates a DeleteAppsAppOK with default headers values
func NewDeleteAppsAppOK() *DeleteAppsAppOK {
	return &DeleteAppsAppOK{}
}

/*DeleteAppsAppOK handles this case with default header values.

Apps successfully deleted.
*/
type DeleteAppsAppOK struct {
}

func (o *DeleteAppsAppOK) Error() string {
	return fmt.Sprintf("[DELETE /apps/{app}][%d] deleteAppsAppOK ", 200)
}

func (o *DeleteAppsAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
