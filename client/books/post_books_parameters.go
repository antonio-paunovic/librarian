// Code generated by go-swagger; DO NOT EDIT.

package books

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

// NewPostBooksParams creates a new PostBooksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostBooksParams() *PostBooksParams {
	return &PostBooksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostBooksParamsWithTimeout creates a new PostBooksParams object
// with the ability to set a timeout on a request.
func NewPostBooksParamsWithTimeout(timeout time.Duration) *PostBooksParams {
	return &PostBooksParams{
		timeout: timeout,
	}
}

// NewPostBooksParamsWithContext creates a new PostBooksParams object
// with the ability to set a context for a request.
func NewPostBooksParamsWithContext(ctx context.Context) *PostBooksParams {
	return &PostBooksParams{
		Context: ctx,
	}
}

// NewPostBooksParamsWithHTTPClient creates a new PostBooksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostBooksParamsWithHTTPClient(client *http.Client) *PostBooksParams {
	return &PostBooksParams{
		HTTPClient: client,
	}
}

/*
PostBooksParams contains all the parameters to send to the API endpoint

	for the post books operation.

	Typically these are written to a http.Request.
*/
type PostBooksParams struct {

	// Body.
	Body *models.Book

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post books params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBooksParams) WithDefaults() *PostBooksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post books params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostBooksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post books params
func (o *PostBooksParams) WithTimeout(timeout time.Duration) *PostBooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post books params
func (o *PostBooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post books params
func (o *PostBooksParams) WithContext(ctx context.Context) *PostBooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post books params
func (o *PostBooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post books params
func (o *PostBooksParams) WithHTTPClient(client *http.Client) *PostBooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post books params
func (o *PostBooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post books params
func (o *PostBooksParams) WithBody(body *models.Book) *PostBooksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post books params
func (o *PostBooksParams) SetBody(body *models.Book) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostBooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
