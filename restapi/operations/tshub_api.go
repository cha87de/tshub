// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTshubAPI creates a new Tshub instance
func NewTshubAPI(spec *loads.Document) *TshubAPI {
	return &TshubAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		GetDomainHandler: GetDomainHandlerFunc(func(params GetDomainParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDomain has not yet been implemented")
		}),
		GetDomainPlotdataHandler: GetDomainPlotdataHandlerFunc(func(params GetDomainPlotdataParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDomainPlotdata has not yet been implemented")
		}),
		GetDomainsHandler: GetDomainsHandlerFunc(func(params GetDomainsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetDomains has not yet been implemented")
		}),
		GetHostHandler: GetHostHandlerFunc(func(params GetHostParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHost has not yet been implemented")
		}),
		GetHostPlotdataHandler: GetHostPlotdataHandlerFunc(func(params GetHostPlotdataParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHostPlotdata has not yet been implemented")
		}),
		GetHostsHandler: GetHostsHandlerFunc(func(params GetHostsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHosts has not yet been implemented")
		}),
		GetProfileHandler: GetProfileHandlerFunc(func(params GetProfileParams) middleware.Responder {
			return middleware.NotImplemented("operation GetProfile has not yet been implemented")
		}),
		GetProfileNamesHandler: GetProfileNamesHandlerFunc(func(params GetProfileNamesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetProfileNames has not yet been implemented")
		}),
	}
}

/*TshubAPI The REST API for the TimeSeries-Profiler */
type TshubAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GetDomainHandler sets the operation handler for the get domain operation
	GetDomainHandler GetDomainHandler
	// GetDomainPlotdataHandler sets the operation handler for the get domain plotdata operation
	GetDomainPlotdataHandler GetDomainPlotdataHandler
	// GetDomainsHandler sets the operation handler for the get domains operation
	GetDomainsHandler GetDomainsHandler
	// GetHostHandler sets the operation handler for the get host operation
	GetHostHandler GetHostHandler
	// GetHostPlotdataHandler sets the operation handler for the get host plotdata operation
	GetHostPlotdataHandler GetHostPlotdataHandler
	// GetHostsHandler sets the operation handler for the get hosts operation
	GetHostsHandler GetHostsHandler
	// GetProfileHandler sets the operation handler for the get profile operation
	GetProfileHandler GetProfileHandler
	// GetProfileNamesHandler sets the operation handler for the get profile names operation
	GetProfileNamesHandler GetProfileNamesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TshubAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TshubAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TshubAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TshubAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TshubAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TshubAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TshubAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TshubAPI
func (o *TshubAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetDomainHandler == nil {
		unregistered = append(unregistered, "GetDomainHandler")
	}

	if o.GetDomainPlotdataHandler == nil {
		unregistered = append(unregistered, "GetDomainPlotdataHandler")
	}

	if o.GetDomainsHandler == nil {
		unregistered = append(unregistered, "GetDomainsHandler")
	}

	if o.GetHostHandler == nil {
		unregistered = append(unregistered, "GetHostHandler")
	}

	if o.GetHostPlotdataHandler == nil {
		unregistered = append(unregistered, "GetHostPlotdataHandler")
	}

	if o.GetHostsHandler == nil {
		unregistered = append(unregistered, "GetHostsHandler")
	}

	if o.GetProfileHandler == nil {
		unregistered = append(unregistered, "GetProfileHandler")
	}

	if o.GetProfileNamesHandler == nil {
		unregistered = append(unregistered, "GetProfileNamesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TshubAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TshubAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *TshubAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *TshubAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TshubAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TshubAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tshub API
func (o *TshubAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TshubAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/domain/{domainname}"] = NewGetDomain(o.context, o.GetDomainHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/domain/{domainname}/plotdata"] = NewGetDomainPlotdata(o.context, o.GetDomainPlotdataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/domains"] = NewGetDomains(o.context, o.GetDomainsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/host/{hostname}"] = NewGetHost(o.context, o.GetHostHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/host/{hostname}/plotdata"] = NewGetHostPlotdata(o.context, o.GetHostPlotdataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/hosts"] = NewGetHosts(o.context, o.GetHostsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profile/{profilename}"] = NewGetProfile(o.context, o.GetProfileHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/profiles"] = NewGetProfileNames(o.context, o.GetProfileNamesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TshubAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TshubAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TshubAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TshubAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
