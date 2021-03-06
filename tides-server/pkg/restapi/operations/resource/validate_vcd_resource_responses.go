// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ValidateVcdResourceOKCode is the HTTP code returned for type ValidateVcdResourceOK
const ValidateVcdResourceOKCode int = 200

/*ValidateVcdResourceOK returns the list of data centers belonging to the host

swagger:response validateVcdResourceOK
*/
type ValidateVcdResourceOK struct {

	/*
	  In: Body
	*/
	Payload *ValidateVcdResourceOKBody `json:"body,omitempty"`
}

// NewValidateVcdResourceOK creates ValidateVcdResourceOK with default headers values
func NewValidateVcdResourceOK() *ValidateVcdResourceOK {

	return &ValidateVcdResourceOK{}
}

// WithPayload adds the payload to the validate vcd resource o k response
func (o *ValidateVcdResourceOK) WithPayload(payload *ValidateVcdResourceOKBody) *ValidateVcdResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate vcd resource o k response
func (o *ValidateVcdResourceOK) SetPayload(payload *ValidateVcdResourceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateVcdResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ValidateVcdResourceUnauthorizedCode is the HTTP code returned for type ValidateVcdResourceUnauthorized
const ValidateVcdResourceUnauthorizedCode int = 401

/*ValidateVcdResourceUnauthorized Unauthorized

swagger:response validateVcdResourceUnauthorized
*/
type ValidateVcdResourceUnauthorized struct {
}

// NewValidateVcdResourceUnauthorized creates ValidateVcdResourceUnauthorized with default headers values
func NewValidateVcdResourceUnauthorized() *ValidateVcdResourceUnauthorized {

	return &ValidateVcdResourceUnauthorized{}
}

// WriteResponse to the client
func (o *ValidateVcdResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ValidateVcdResourceNotFoundCode is the HTTP code returned for type ValidateVcdResourceNotFound
const ValidateVcdResourceNotFoundCode int = 404

/*ValidateVcdResourceNotFound resource not found

swagger:response validateVcdResourceNotFound
*/
type ValidateVcdResourceNotFound struct {

	/*
	  In: Body
	*/
	Payload *ValidateVcdResourceNotFoundBody `json:"body,omitempty"`
}

// NewValidateVcdResourceNotFound creates ValidateVcdResourceNotFound with default headers values
func NewValidateVcdResourceNotFound() *ValidateVcdResourceNotFound {

	return &ValidateVcdResourceNotFound{}
}

// WithPayload adds the payload to the validate vcd resource not found response
func (o *ValidateVcdResourceNotFound) WithPayload(payload *ValidateVcdResourceNotFoundBody) *ValidateVcdResourceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate vcd resource not found response
func (o *ValidateVcdResourceNotFound) SetPayload(payload *ValidateVcdResourceNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateVcdResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
