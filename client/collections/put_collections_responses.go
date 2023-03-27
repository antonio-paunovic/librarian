// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"librarian/models"
)

// PutCollectionsReader is a Reader for the PutCollections structure.
type PutCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutCollectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCollectionsOK creates a PutCollectionsOK with default headers values
func NewPutCollectionsOK() *PutCollectionsOK {
	return &PutCollectionsOK{}
}

/*
PutCollectionsOK describes a response with status code 200, with default header values.

Collection by given name.
*/
type PutCollectionsOK struct {
	Payload *models.Collection
}

// IsSuccess returns true when this put collections o k response has a 2xx status code
func (o *PutCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put collections o k response has a 3xx status code
func (o *PutCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put collections o k response has a 4xx status code
func (o *PutCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put collections o k response has a 5xx status code
func (o *PutCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put collections o k response a status code equal to that given
func (o *PutCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put collections o k response
func (o *PutCollectionsOK) Code() int {
	return 200
}

func (o *PutCollectionsOK) Error() string {
	return fmt.Sprintf("[PUT /collections][%d] putCollectionsOK  %+v", 200, o.Payload)
}

func (o *PutCollectionsOK) String() string {
	return fmt.Sprintf("[PUT /collections][%d] putCollectionsOK  %+v", 200, o.Payload)
}

func (o *PutCollectionsOK) GetPayload() *models.Collection {
	return o.Payload
}

func (o *PutCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCollectionsNotFound creates a PutCollectionsNotFound with default headers values
func NewPutCollectionsNotFound() *PutCollectionsNotFound {
	return &PutCollectionsNotFound{}
}

/*
PutCollectionsNotFound describes a response with status code 404, with default header values.

Collection not found.
*/
type PutCollectionsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this put collections not found response has a 2xx status code
func (o *PutCollectionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put collections not found response has a 3xx status code
func (o *PutCollectionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put collections not found response has a 4xx status code
func (o *PutCollectionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put collections not found response has a 5xx status code
func (o *PutCollectionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put collections not found response a status code equal to that given
func (o *PutCollectionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put collections not found response
func (o *PutCollectionsNotFound) Code() int {
	return 404
}

func (o *PutCollectionsNotFound) Error() string {
	return fmt.Sprintf("[PUT /collections][%d] putCollectionsNotFound  %+v", 404, o.Payload)
}

func (o *PutCollectionsNotFound) String() string {
	return fmt.Sprintf("[PUT /collections][%d] putCollectionsNotFound  %+v", 404, o.Payload)
}

func (o *PutCollectionsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCollectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCollectionsDefault creates a PutCollectionsDefault with default headers values
func NewPutCollectionsDefault(code int) *PutCollectionsDefault {
	return &PutCollectionsDefault{
		_statusCode: code,
	}
}

/*
PutCollectionsDefault describes a response with status code -1, with default header values.

bad input parameter/missing collection
*/
type PutCollectionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this put collections default response has a 2xx status code
func (o *PutCollectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put collections default response has a 3xx status code
func (o *PutCollectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put collections default response has a 4xx status code
func (o *PutCollectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put collections default response has a 5xx status code
func (o *PutCollectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put collections default response a status code equal to that given
func (o *PutCollectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put collections default response
func (o *PutCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *PutCollectionsDefault) Error() string {
	return fmt.Sprintf("[PUT /collections][%d] PutCollections default  %+v", o._statusCode, o.Payload)
}

func (o *PutCollectionsDefault) String() string {
	return fmt.Sprintf("[PUT /collections][%d] PutCollections default  %+v", o._statusCode, o.Payload)
}

func (o *PutCollectionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}