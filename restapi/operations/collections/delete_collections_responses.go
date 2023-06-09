// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"librarian/models"
)

// DeleteCollectionsNoContentCode is the HTTP code returned for type DeleteCollectionsNoContent
const DeleteCollectionsNoContentCode int = 204

/*
DeleteCollectionsNoContent No content on successful delete

swagger:response deleteCollectionsNoContent
*/
type DeleteCollectionsNoContent struct {

	/*
	  In: Body
	*/
	Payload *models.Collection `json:"body,omitempty"`
}

// NewDeleteCollectionsNoContent creates DeleteCollectionsNoContent with default headers values
func NewDeleteCollectionsNoContent() *DeleteCollectionsNoContent {

	return &DeleteCollectionsNoContent{}
}

// WithPayload adds the payload to the delete collections no content response
func (o *DeleteCollectionsNoContent) WithPayload(payload *models.Collection) *DeleteCollectionsNoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete collections no content response
func (o *DeleteCollectionsNoContent) SetPayload(payload *models.Collection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCollectionsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteCollectionsDefault bad input parametr/missing collection

swagger:response deleteCollectionsDefault
*/
type DeleteCollectionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteCollectionsDefault creates DeleteCollectionsDefault with default headers values
func NewDeleteCollectionsDefault(code int) *DeleteCollectionsDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteCollectionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete collections default response
func (o *DeleteCollectionsDefault) WithStatusCode(code int) *DeleteCollectionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete collections default response
func (o *DeleteCollectionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete collections default response
func (o *DeleteCollectionsDefault) WithPayload(payload *models.Error) *DeleteCollectionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete collections default response
func (o *DeleteCollectionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCollectionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
