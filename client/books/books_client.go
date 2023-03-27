// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new books API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for books API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteBooksIsbn(params *DeleteBooksIsbnParams, opts ...ClientOption) (*DeleteBooksIsbnNoContent, error)

	GetBooks(params *GetBooksParams, opts ...ClientOption) (*GetBooksOK, error)

	GetBooksIsbn(params *GetBooksIsbnParams, opts ...ClientOption) (*GetBooksIsbnOK, error)

	PostBooks(params *PostBooksParams, opts ...ClientOption) (*PostBooksOK, error)

	PutBooksIsbn(params *PutBooksIsbnParams, opts ...ClientOption) (*PutBooksIsbnOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteBooksIsbn deletes book by isbn
*/
func (a *Client) DeleteBooksIsbn(params *DeleteBooksIsbnParams, opts ...ClientOption) (*DeleteBooksIsbnNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBooksIsbnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBooksIsbn",
		Method:             "DELETE",
		PathPattern:        "/books/{isbn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBooksIsbnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBooksIsbnNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteBooksIsbnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBooks finds book by isbn
*/
func (a *Client) GetBooks(params *GetBooksParams, opts ...ClientOption) (*GetBooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBooks",
		Method:             "GET",
		PathPattern:        "/books",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBooksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBooksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetBooksIsbn finds book by isbn
*/
func (a *Client) GetBooksIsbn(params *GetBooksIsbnParams, opts ...ClientOption) (*GetBooksIsbnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBooksIsbnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBooksIsbn",
		Method:             "GET",
		PathPattern:        "/books/{isbn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBooksIsbnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBooksIsbnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBooksIsbnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostBooks creates a book
*/
func (a *Client) PostBooks(params *PostBooksParams, opts ...ClientOption) (*PostBooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostBooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostBooks",
		Method:             "POST",
		PathPattern:        "/books",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostBooksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostBooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostBooksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutBooksIsbn updates book by isbn
*/
func (a *Client) PutBooksIsbn(params *PutBooksIsbnParams, opts ...ClientOption) (*PutBooksIsbnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutBooksIsbnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutBooksIsbn",
		Method:             "PUT",
		PathPattern:        "/books/{isbn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutBooksIsbnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutBooksIsbnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutBooksIsbnDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}