package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// DeleteAPIFeaturesIDReader is a Reader for the DeleteAPIFeaturesID structure.
type DeleteAPIFeaturesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIFeaturesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAPIFeaturesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteAPIFeaturesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteAPIFeaturesIDNoContent creates a DeleteAPIFeaturesIDNoContent with default headers values
func NewDeleteAPIFeaturesIDNoContent() *DeleteAPIFeaturesIDNoContent {
	return &DeleteAPIFeaturesIDNoContent{}
}

/*DeleteAPIFeaturesIDNoContent handles this case with default header values.

Feature successfully deleted
*/
type DeleteAPIFeaturesIDNoContent struct {
}

func (o *DeleteAPIFeaturesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/features/{id}][%d] deleteApiFeaturesIdNoContent ", 204)
}

func (o *DeleteAPIFeaturesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAPIFeaturesIDDefault creates a DeleteAPIFeaturesIDDefault with default headers values
func NewDeleteAPIFeaturesIDDefault(code int) *DeleteAPIFeaturesIDDefault {
	return &DeleteAPIFeaturesIDDefault{
		_statusCode: code,
	}
}

/*DeleteAPIFeaturesIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteAPIFeaturesIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete API features ID default response
func (o *DeleteAPIFeaturesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAPIFeaturesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/features/{id}][%d] DeleteAPIFeaturesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAPIFeaturesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
