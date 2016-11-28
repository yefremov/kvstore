package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindKeysHandlerFunc turns a function with the right signature into a find keys handler
type FindKeysHandlerFunc func(FindKeysParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindKeysHandlerFunc) Handle(params FindKeysParams) middleware.Responder {
	return fn(params)
}

// FindKeysHandler interface for that can handle valid find keys params
type FindKeysHandler interface {
	Handle(FindKeysParams) middleware.Responder
}

// NewFindKeys creates a new http.Handler for the find keys operation
func NewFindKeys(ctx *middleware.Context, handler FindKeysHandler) *FindKeys {
	return &FindKeys{Context: ctx, Handler: handler}
}

/*FindKeys swagger:route GET /kv kv findKeys

lists all the keys

*/
type FindKeys struct {
	Context *middleware.Context
	Handler FindKeysHandler
}

func (o *FindKeys) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewFindKeysParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}