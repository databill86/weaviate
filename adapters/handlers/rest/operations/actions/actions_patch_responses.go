//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsPatchNoContentCode is the HTTP code returned for type ActionsPatchNoContent
const ActionsPatchNoContentCode int = 204

/*ActionsPatchNoContent Successfully applied. No content provided.

swagger:response actionsPatchNoContent
*/
type ActionsPatchNoContent struct {
}

// NewActionsPatchNoContent creates ActionsPatchNoContent with default headers values
func NewActionsPatchNoContent() *ActionsPatchNoContent {

	return &ActionsPatchNoContent{}
}

// WriteResponse to the client
func (o *ActionsPatchNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ActionsPatchBadRequestCode is the HTTP code returned for type ActionsPatchBadRequest
const ActionsPatchBadRequestCode int = 400

/*ActionsPatchBadRequest The patch-JSON is malformed.

swagger:response actionsPatchBadRequest
*/
type ActionsPatchBadRequest struct {
}

// NewActionsPatchBadRequest creates ActionsPatchBadRequest with default headers values
func NewActionsPatchBadRequest() *ActionsPatchBadRequest {

	return &ActionsPatchBadRequest{}
}

// WriteResponse to the client
func (o *ActionsPatchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ActionsPatchUnauthorizedCode is the HTTP code returned for type ActionsPatchUnauthorized
const ActionsPatchUnauthorizedCode int = 401

/*ActionsPatchUnauthorized Unauthorized or invalid credentials.

swagger:response actionsPatchUnauthorized
*/
type ActionsPatchUnauthorized struct {
}

// NewActionsPatchUnauthorized creates ActionsPatchUnauthorized with default headers values
func NewActionsPatchUnauthorized() *ActionsPatchUnauthorized {

	return &ActionsPatchUnauthorized{}
}

// WriteResponse to the client
func (o *ActionsPatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActionsPatchForbiddenCode is the HTTP code returned for type ActionsPatchForbidden
const ActionsPatchForbiddenCode int = 403

/*ActionsPatchForbidden Forbidden

swagger:response actionsPatchForbidden
*/
type ActionsPatchForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsPatchForbidden creates ActionsPatchForbidden with default headers values
func NewActionsPatchForbidden() *ActionsPatchForbidden {

	return &ActionsPatchForbidden{}
}

// WithPayload adds the payload to the actions patch forbidden response
func (o *ActionsPatchForbidden) WithPayload(payload *models.ErrorResponse) *ActionsPatchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions patch forbidden response
func (o *ActionsPatchForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsPatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsPatchNotFoundCode is the HTTP code returned for type ActionsPatchNotFound
const ActionsPatchNotFoundCode int = 404

/*ActionsPatchNotFound Successful query result but no resource was found.

swagger:response actionsPatchNotFound
*/
type ActionsPatchNotFound struct {
}

// NewActionsPatchNotFound creates ActionsPatchNotFound with default headers values
func NewActionsPatchNotFound() *ActionsPatchNotFound {

	return &ActionsPatchNotFound{}
}

// WriteResponse to the client
func (o *ActionsPatchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ActionsPatchUnprocessableEntityCode is the HTTP code returned for type ActionsPatchUnprocessableEntity
const ActionsPatchUnprocessableEntityCode int = 422

/*ActionsPatchUnprocessableEntity The patch-JSON is valid but unprocessable.

swagger:response actionsPatchUnprocessableEntity
*/
type ActionsPatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsPatchUnprocessableEntity creates ActionsPatchUnprocessableEntity with default headers values
func NewActionsPatchUnprocessableEntity() *ActionsPatchUnprocessableEntity {

	return &ActionsPatchUnprocessableEntity{}
}

// WithPayload adds the payload to the actions patch unprocessable entity response
func (o *ActionsPatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ActionsPatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions patch unprocessable entity response
func (o *ActionsPatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsPatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActionsPatchInternalServerErrorCode is the HTTP code returned for type ActionsPatchInternalServerError
const ActionsPatchInternalServerErrorCode int = 500

/*ActionsPatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response actionsPatchInternalServerError
*/
type ActionsPatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActionsPatchInternalServerError creates ActionsPatchInternalServerError with default headers values
func NewActionsPatchInternalServerError() *ActionsPatchInternalServerError {

	return &ActionsPatchInternalServerError{}
}

// WithPayload adds the payload to the actions patch internal server error response
func (o *ActionsPatchInternalServerError) WithPayload(payload *models.ErrorResponse) *ActionsPatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the actions patch internal server error response
func (o *ActionsPatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActionsPatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
