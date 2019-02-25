// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetHostPlotdataHandlerFunc turns a function with the right signature into a get host plotdata handler
type GetHostPlotdataHandlerFunc func(GetHostPlotdataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHostPlotdataHandlerFunc) Handle(params GetHostPlotdataParams) middleware.Responder {
	return fn(params)
}

// GetHostPlotdataHandler interface for that can handle valid get host plotdata params
type GetHostPlotdataHandler interface {
	Handle(GetHostPlotdataParams) middleware.Responder
}

// NewGetHostPlotdata creates a new http.Handler for the get host plotdata operation
func NewGetHostPlotdata(ctx *middleware.Context, handler GetHostPlotdataHandler) *GetHostPlotdata {
	return &GetHostPlotdata{Context: ctx, Handler: handler}
}

/*GetHostPlotdata swagger:route GET /host/{hostname}/plotdata getHostPlotdata

Returns the plotdata (past and prediction) of given host and metric

*/
type GetHostPlotdata struct {
	Context *middleware.Context
	Handler GetHostPlotdataHandler
}

func (o *GetHostPlotdata) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetHostPlotdataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
