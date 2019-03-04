// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetProfileNamesOKCode is the HTTP code returned for type GetProfileNamesOK
const GetProfileNamesOKCode int = 200

/*GetProfileNamesOK An array profile names

swagger:response getProfileNamesOK
*/
type GetProfileNamesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetProfileNamesOK creates GetProfileNamesOK with default headers values
func NewGetProfileNamesOK() *GetProfileNamesOK {

	return &GetProfileNamesOK{}
}

// WithPayload adds the payload to the get profile names o k response
func (o *GetProfileNamesOK) WithPayload(payload []string) *GetProfileNamesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile names o k response
func (o *GetProfileNamesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileNamesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetProfileNamesNotFoundCode is the HTTP code returned for type GetProfileNamesNotFound
const GetProfileNamesNotFoundCode int = 404

/*GetProfileNamesNotFound Element not found

swagger:response getProfileNamesNotFound
*/
type GetProfileNamesNotFound struct {
}

// NewGetProfileNamesNotFound creates GetProfileNamesNotFound with default headers values
func NewGetProfileNamesNotFound() *GetProfileNamesNotFound {

	return &GetProfileNamesNotFound{}
}

// WriteResponse to the client
func (o *GetProfileNamesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetProfileNamesInternalServerErrorCode is the HTTP code returned for type GetProfileNamesInternalServerError
const GetProfileNamesInternalServerErrorCode int = 500

/*GetProfileNamesInternalServerError Internal Server Error

swagger:response getProfileNamesInternalServerError
*/
type GetProfileNamesInternalServerError struct {
}

// NewGetProfileNamesInternalServerError creates GetProfileNamesInternalServerError with default headers values
func NewGetProfileNamesInternalServerError() *GetProfileNamesInternalServerError {

	return &GetProfileNamesInternalServerError{}
}

// WriteResponse to the client
func (o *GetProfileNamesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}