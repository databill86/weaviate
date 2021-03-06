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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewThingsReferencesCreateParams creates a new ThingsReferencesCreateParams object
// with the default values initialized.
func NewThingsReferencesCreateParams() *ThingsReferencesCreateParams {
	var ()
	return &ThingsReferencesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsReferencesCreateParamsWithTimeout creates a new ThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsReferencesCreateParamsWithTimeout(timeout time.Duration) *ThingsReferencesCreateParams {
	var ()
	return &ThingsReferencesCreateParams{

		timeout: timeout,
	}
}

// NewThingsReferencesCreateParamsWithContext creates a new ThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsReferencesCreateParamsWithContext(ctx context.Context) *ThingsReferencesCreateParams {
	var ()
	return &ThingsReferencesCreateParams{

		Context: ctx,
	}
}

// NewThingsReferencesCreateParamsWithHTTPClient creates a new ThingsReferencesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsReferencesCreateParamsWithHTTPClient(client *http.Client) *ThingsReferencesCreateParams {
	var ()
	return &ThingsReferencesCreateParams{
		HTTPClient: client,
	}
}

/*ThingsReferencesCreateParams contains all the parameters to send to the API endpoint
for the things references create operation typically these are written to a http.Request
*/
type ThingsReferencesCreateParams struct {

	/*Body*/
	Body *models.SingleRef
	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID
	/*PropertyName
	  Unique name of the property related to the Thing.

	*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things references create params
func (o *ThingsReferencesCreateParams) WithTimeout(timeout time.Duration) *ThingsReferencesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things references create params
func (o *ThingsReferencesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things references create params
func (o *ThingsReferencesCreateParams) WithContext(ctx context.Context) *ThingsReferencesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things references create params
func (o *ThingsReferencesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things references create params
func (o *ThingsReferencesCreateParams) WithHTTPClient(client *http.Client) *ThingsReferencesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things references create params
func (o *ThingsReferencesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the things references create params
func (o *ThingsReferencesCreateParams) WithBody(body *models.SingleRef) *ThingsReferencesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the things references create params
func (o *ThingsReferencesCreateParams) SetBody(body *models.SingleRef) {
	o.Body = body
}

// WithID adds the id to the things references create params
func (o *ThingsReferencesCreateParams) WithID(id strfmt.UUID) *ThingsReferencesCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the things references create params
func (o *ThingsReferencesCreateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPropertyName adds the propertyName to the things references create params
func (o *ThingsReferencesCreateParams) WithPropertyName(propertyName string) *ThingsReferencesCreateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the things references create params
func (o *ThingsReferencesCreateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsReferencesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
