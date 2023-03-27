// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"librarian/models"
)

// NewPutCollectionsParams creates a new PutCollectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCollectionsParams() *PutCollectionsParams {
	return &PutCollectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCollectionsParamsWithTimeout creates a new PutCollectionsParams object
// with the ability to set a timeout on a request.
func NewPutCollectionsParamsWithTimeout(timeout time.Duration) *PutCollectionsParams {
	return &PutCollectionsParams{
		timeout: timeout,
	}
}

// NewPutCollectionsParamsWithContext creates a new PutCollectionsParams object
// with the ability to set a context for a request.
func NewPutCollectionsParamsWithContext(ctx context.Context) *PutCollectionsParams {
	return &PutCollectionsParams{
		Context: ctx,
	}
}

// NewPutCollectionsParamsWithHTTPClient creates a new PutCollectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCollectionsParamsWithHTTPClient(client *http.Client) *PutCollectionsParams {
	return &PutCollectionsParams{
		HTTPClient: client,
	}
}

/*
PutCollectionsParams contains all the parameters to send to the API endpoint

	for the put collections operation.

	Typically these are written to a http.Request.
*/
type PutCollectionsParams struct {

	// Body.
	Body *models.Collection

	/* Name.

	   Name of collection
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCollectionsParams) WithDefaults() *PutCollectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCollectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put collections params
func (o *PutCollectionsParams) WithTimeout(timeout time.Duration) *PutCollectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put collections params
func (o *PutCollectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put collections params
func (o *PutCollectionsParams) WithContext(ctx context.Context) *PutCollectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put collections params
func (o *PutCollectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put collections params
func (o *PutCollectionsParams) WithHTTPClient(client *http.Client) *PutCollectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put collections params
func (o *PutCollectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put collections params
func (o *PutCollectionsParams) WithBody(body *models.Collection) *PutCollectionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put collections params
func (o *PutCollectionsParams) SetBody(body *models.Collection) {
	o.Body = body
}

// WithName adds the name to the put collections params
func (o *PutCollectionsParams) WithName(name string) *PutCollectionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put collections params
func (o *PutCollectionsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutCollectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}