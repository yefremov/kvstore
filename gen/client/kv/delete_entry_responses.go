package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/kvstore/gen/models"
)

// DeleteEntryReader is a Reader for the DeleteEntry structure.
type DeleteEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEntryNoContent creates a DeleteEntryNoContent with default headers values
func NewDeleteEntryNoContent() *DeleteEntryNoContent {
	return &DeleteEntryNoContent{}
}

/*DeleteEntryNoContent handles this case with default header values.

the delete was successful
*/
type DeleteEntryNoContent struct {
	/*The request id this is a response to
	 */
	XRequestID string
}

func (o *DeleteEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /kv/{key}][%d] deleteEntryNoContent ", 204)
}

func (o *DeleteEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewDeleteEntryDefault creates a DeleteEntryDefault with default headers values
func NewDeleteEntryDefault(code int) *DeleteEntryDefault {
	return &DeleteEntryDefault{
		_statusCode: code,
	}
}

/*DeleteEntryDefault handles this case with default header values.

Error
*/
type DeleteEntryDefault struct {
	_statusCode int

	/*The request id this is a response to
	 */
	XRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete entry default response
func (o *DeleteEntryDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEntryDefault) Error() string {
	return fmt.Sprintf("[DELETE /kv/{key}][%d] deleteEntry default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}