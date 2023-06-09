// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"librarian/models"
)

// DeleteBooksIsbnNoContentCode is the HTTP code returned for type DeleteBooksIsbnNoContent
const DeleteBooksIsbnNoContentCode int = 204

/*
DeleteBooksIsbnNoContent No content on successful delete

swagger:response deleteBooksIsbnNoContent
*/
type DeleteBooksIsbnNoContent struct {

	/*
	  In: Body
	*/
	Payload *models.Book `json:"body,omitempty"`
}

// NewDeleteBooksIsbnNoContent creates DeleteBooksIsbnNoContent with default headers values
func NewDeleteBooksIsbnNoContent() *DeleteBooksIsbnNoContent {

	return &DeleteBooksIsbnNoContent{}
}

// WithPayload adds the payload to the delete books isbn no content response
func (o *DeleteBooksIsbnNoContent) WithPayload(payload *models.Book) *DeleteBooksIsbnNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete books isbn no content response
func (o *DeleteBooksIsbnNoContent) SetPayload(payload *models.Book) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBooksIsbnNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBooksIsbnNotFoundCode is the HTTP code returned for type DeleteBooksIsbnNotFound
const DeleteBooksIsbnNotFoundCode int = 404

/*
DeleteBooksIsbnNotFound Book not found.

swagger:response deleteBooksIsbnNotFound
*/
type DeleteBooksIsbnNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBooksIsbnNotFound creates DeleteBooksIsbnNotFound with default headers values
func NewDeleteBooksIsbnNotFound() *DeleteBooksIsbnNotFound {

	return &DeleteBooksIsbnNotFound{}
}

// WithPayload adds the payload to the delete books isbn not found response
func (o *DeleteBooksIsbnNotFound) WithPayload(payload *models.Error) *DeleteBooksIsbnNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete books isbn not found response
func (o *DeleteBooksIsbnNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBooksIsbnNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteBooksIsbnDefault bad input parametr/missing book

swagger:response deleteBooksIsbnDefault
*/
type DeleteBooksIsbnDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBooksIsbnDefault creates DeleteBooksIsbnDefault with default headers values
func NewDeleteBooksIsbnDefault(code int) *DeleteBooksIsbnDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteBooksIsbnDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete books isbn default response
func (o *DeleteBooksIsbnDefault) WithStatusCode(code int) *DeleteBooksIsbnDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete books isbn default response
func (o *DeleteBooksIsbnDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete books isbn default response
func (o *DeleteBooksIsbnDefault) WithPayload(payload *models.Error) *DeleteBooksIsbnDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete books isbn default response
func (o *DeleteBooksIsbnDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBooksIsbnDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
