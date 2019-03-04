// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/cha87de/tshub/models"
)

// GetDomainOKCode is the HTTP code returned for type GetDomainOK
const GetDomainOKCode int = 200

/*GetDomainOK An array of domains

swagger:response getDomainOK
*/
type GetDomainOK struct {

	/*
	  In: Body
	*/
	Payload *models.DomainDetails `json:"body,omitempty"`
}

// NewGetDomainOK creates GetDomainOK with default headers values
func NewGetDomainOK() *GetDomainOK {

	return &GetDomainOK{}
}

// WithPayload adds the payload to the get domain o k response
func (o *GetDomainOK) WithPayload(payload *models.DomainDetails) *GetDomainOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get domain o k response
func (o *GetDomainOK) SetPayload(payload *models.DomainDetails) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDomainOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDomainNotFoundCode is the HTTP code returned for type GetDomainNotFound
const GetDomainNotFoundCode int = 404

/*GetDomainNotFound Element not found

swagger:response getDomainNotFound
*/
type GetDomainNotFound struct {
}

// NewGetDomainNotFound creates GetDomainNotFound with default headers values
func NewGetDomainNotFound() *GetDomainNotFound {

	return &GetDomainNotFound{}
}

// WriteResponse to the client
func (o *GetDomainNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetDomainInternalServerErrorCode is the HTTP code returned for type GetDomainInternalServerError
const GetDomainInternalServerErrorCode int = 500

/*GetDomainInternalServerError Internal Server Error

swagger:response getDomainInternalServerError
*/
type GetDomainInternalServerError struct {
}

// NewGetDomainInternalServerError creates GetDomainInternalServerError with default headers values
func NewGetDomainInternalServerError() *GetDomainInternalServerError {

	return &GetDomainInternalServerError{}
}

// WriteResponse to the client
func (o *GetDomainInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}