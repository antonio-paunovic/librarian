// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCollectionsHandlerFunc turns a function with the right signature into a delete collections handler
type DeleteCollectionsHandlerFunc func(DeleteCollectionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCollectionsHandlerFunc) Handle(params DeleteCollectionsParams) middleware.Responder {
	return fn(params)
}

// DeleteCollectionsHandler interface for that can handle valid delete collections params
type DeleteCollectionsHandler interface {
	Handle(DeleteCollectionsParams) middleware.Responder
}

// NewDeleteCollections creates a new http.Handler for the delete collections operation
func NewDeleteCollections(ctx *middleware.Context, handler DeleteCollectionsHandler) *DeleteCollections {
	return &DeleteCollections{Context: ctx, Handler: handler}
}

/*
	DeleteCollections swagger:route DELETE /collections collections deleteCollections

Delete collection by name
*/
type DeleteCollections struct {
	Context *middleware.Context
	Handler DeleteCollectionsHandler
}

func (o *DeleteCollections) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteCollectionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
