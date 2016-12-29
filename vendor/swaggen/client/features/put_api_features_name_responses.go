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

// PutAPIFeaturesNameReader is a Reader for the PutAPIFeaturesName structure.
type PutAPIFeaturesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIFeaturesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIFeaturesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPutAPIFeaturesNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutAPIFeaturesNameOK creates a PutAPIFeaturesNameOK with default headers values
func NewPutAPIFeaturesNameOK() *PutAPIFeaturesNameOK {
	return &PutAPIFeaturesNameOK{}
}

/*PutAPIFeaturesNameOK handles this case with default header values.

Feature successfully updated
*/
type PutAPIFeaturesNameOK struct {
	Payload *models.Feature
}

func (o *PutAPIFeaturesNameOK) Error() string {
	return fmt.Sprintf("[PUT /api/features/{name}][%d] putApiFeaturesNameOK  %+v", 200, o.Payload)
}

func (o *PutAPIFeaturesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIFeaturesNameDefault creates a PutAPIFeaturesNameDefault with default headers values
func NewPutAPIFeaturesNameDefault(code int) *PutAPIFeaturesNameDefault {
	return &PutAPIFeaturesNameDefault{
		_statusCode: code,
	}
}

/*PutAPIFeaturesNameDefault handles this case with default header values.

Unexpected exception
*/
type PutAPIFeaturesNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put API features name default response
func (o *PutAPIFeaturesNameDefault) Code() int {
	return o._statusCode
}

func (o *PutAPIFeaturesNameDefault) Error() string {
	return fmt.Sprintf("[PUT /api/features/{name}][%d] PutAPIFeaturesName default  %+v", o._statusCode, o.Payload)
}

func (o *PutAPIFeaturesNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}