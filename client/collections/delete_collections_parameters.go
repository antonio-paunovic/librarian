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
)

// NewDeleteCollectionsParams creates a new DeleteCollectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCollectionsParams() *DeleteCollectionsParams {
	return &DeleteCollectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCollectionsParamsWithTimeout creates a new DeleteCollectionsParams object
// with the ability to set a timeout on a request.
func NewDeleteCollectionsParamsWithTimeout(timeout time.Duration) *DeleteCollectionsParams {
	return &DeleteCollectionsParams{
		timeout: timeout,
	}
}

// NewDeleteCollectionsParamsWithContext creates a new DeleteCollectionsParams object
// with the ability to set a context for a request.
func NewDeleteCollectionsParamsWithContext(ctx context.Context) *DeleteCollectionsParams {
	return &DeleteCollectionsParams{
		Context: ctx,
	}
}

// NewDeleteCollectionsParamsWithHTTPClient creates a new DeleteCollectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCollectionsParamsWithHTTPClient(client *http.Client) *DeleteCollectionsParams {
	return &DeleteCollectionsParams{
		HTTPClient: client,
	}
}

/*
DeleteCollectionsParams contains all the parameters to send to the API endpoint

	for the delete collections operation.

	Typically these are written to a http.Request.
*/
type DeleteCollectionsParams struct {

	/* Name.

	   Name of collection
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectionsParams) WithDefaults() *DeleteCollectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCollectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete collections params
func (o *DeleteCollectionsParams) WithTimeout(timeout time.Duration) *DeleteCollectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete collections params
func (o *DeleteCollectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete collections params
func (o *DeleteCollectionsParams) WithContext(ctx context.Context) *DeleteCollectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete collections params
func (o *DeleteCollectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete collections params
func (o *DeleteCollectionsParams) WithHTTPClient(client *http.Client) *DeleteCollectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete collections params
func (o *DeleteCollectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete collections params
func (o *DeleteCollectionsParams) WithName(name string) *DeleteCollectionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete collections params
func (o *DeleteCollectionsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCollectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
