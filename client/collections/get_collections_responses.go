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

// GetCollectionsReader is a Reader for the GetCollections structure.
type GetCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCollectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCollectionsOK creates a GetCollectionsOK with default headers values
func NewGetCollectionsOK() *GetCollectionsOK {
	return &GetCollectionsOK{}
}

/*
GetCollectionsOK describes a response with status code 200, with default header values.

All collections from the library.
*/
type GetCollectionsOK struct {
	Payload []*models.Collection
}

// IsSuccess returns true when this get collections o k response has a 2xx status code
func (o *GetCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get collections o k response has a 3xx status code
func (o *GetCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collections o k response has a 4xx status code
func (o *GetCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get collections o k response has a 5xx status code
func (o *GetCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get collections o k response a status code equal to that given
func (o *GetCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get collections o k response
func (o *GetCollectionsOK) Code() int {
	return 200
}

func (o *GetCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /collections][%d] getCollectionsOK  %+v", 200, o.Payload)
}

func (o *GetCollectionsOK) String() string {
	return fmt.Sprintf("[GET /collections][%d] getCollectionsOK  %+v", 200, o.Payload)
}

func (o *GetCollectionsOK) GetPayload() []*models.Collection {
	return o.Payload
}

func (o *GetCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectionsNotFound creates a GetCollectionsNotFound with default headers values
func NewGetCollectionsNotFound() *GetCollectionsNotFound {
	return &GetCollectionsNotFound{}
}

/*
GetCollectionsNotFound describes a response with status code 404, with default header values.

Collection not found.
*/
type GetCollectionsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get collections not found response has a 2xx status code
func (o *GetCollectionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get collections not found response has a 3xx status code
func (o *GetCollectionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get collections not found response has a 4xx status code
func (o *GetCollectionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get collections not found response has a 5xx status code
func (o *GetCollectionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get collections not found response a status code equal to that given
func (o *GetCollectionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get collections not found response
func (o *GetCollectionsNotFound) Code() int {
	return 404
}

func (o *GetCollectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /collections][%d] getCollectionsNotFound  %+v", 404, o.Payload)
}

func (o *GetCollectionsNotFound) String() string {
	return fmt.Sprintf("[GET /collections][%d] getCollectionsNotFound  %+v", 404, o.Payload)
}

func (o *GetCollectionsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCollectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCollectionsDefault creates a GetCollectionsDefault with default headers values
func NewGetCollectionsDefault(code int) *GetCollectionsDefault {
	return &GetCollectionsDefault{
		_statusCode: code,
	}
}

/*
GetCollectionsDefault describes a response with status code -1, with default header values.

bad input parameter/missing collection
*/
type GetCollectionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get collections default response has a 2xx status code
func (o *GetCollectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get collections default response has a 3xx status code
func (o *GetCollectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get collections default response has a 4xx status code
func (o *GetCollectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get collections default response has a 5xx status code
func (o *GetCollectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get collections default response a status code equal to that given
func (o *GetCollectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get collections default response
func (o *GetCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *GetCollectionsDefault) Error() string {
	return fmt.Sprintf("[GET /collections][%d] GetCollections default  %+v", o._statusCode, o.Payload)
}

func (o *GetCollectionsDefault) String() string {
	return fmt.Sprintf("[GET /collections][%d] GetCollections default  %+v", o._statusCode, o.Payload)
}

func (o *GetCollectionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
