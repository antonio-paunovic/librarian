// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"librarian/models"
)

// PostBooksOKCode is the HTTP code returned for type PostBooksOK
const PostBooksOKCode int = 200

/*
PostBooksOK Created a book

swagger:response postBooksOK
*/
type PostBooksOK struct {

	/*
	  In: Body
	*/
	Payload *models.Book `json:"body,omitempty"`
}

// NewPostBooksOK creates PostBooksOK with default headers values
func NewPostBooksOK() *PostBooksOK {

	return &PostBooksOK{}
}

// WithPayload adds the payload to the post books o k response
func (o *PostBooksOK) WithPayload(payload *models.Book) *PostBooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post books o k response
func (o *PostBooksOK) SetPayload(payload *models.Book) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostBooksDefault bad input for book

swagger:response postBooksDefault
*/
type PostBooksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostBooksDefault creates PostBooksDefault with default headers values
func NewPostBooksDefault(code int) *PostBooksDefault {
	if code <= 0 {
		code = 500
	}

	return &PostBooksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post books default response
func (o *PostBooksDefault) WithStatusCode(code int) *PostBooksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post books default response
func (o *PostBooksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post books default response
func (o *PostBooksDefault) WithPayload(payload *models.Error) *PostBooksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post books default response
func (o *PostBooksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBooksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
