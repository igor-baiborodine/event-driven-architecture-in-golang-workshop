// Code generated by go-swagger; DO NOT EDIT.

package shopping_list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"eda-in-golang/depot/depotclient/models"
)

// NewAssignShoppingListParams creates a new AssignShoppingListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignShoppingListParams() *AssignShoppingListParams {
	return &AssignShoppingListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignShoppingListParamsWithTimeout creates a new AssignShoppingListParams object
// with the ability to set a timeout on a request.
func NewAssignShoppingListParamsWithTimeout(timeout time.Duration) *AssignShoppingListParams {
	return &AssignShoppingListParams{
		timeout: timeout,
	}
}

// NewAssignShoppingListParamsWithContext creates a new AssignShoppingListParams object
// with the ability to set a context for a request.
func NewAssignShoppingListParamsWithContext(ctx context.Context) *AssignShoppingListParams {
	return &AssignShoppingListParams{
		Context: ctx,
	}
}

// NewAssignShoppingListParamsWithHTTPClient creates a new AssignShoppingListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignShoppingListParamsWithHTTPClient(client *http.Client) *AssignShoppingListParams {
	return &AssignShoppingListParams{
		HTTPClient: client,
	}
}

/* AssignShoppingListParams contains all the parameters to send to the API endpoint
   for the assign shopping list operation.

   Typically these are written to a http.Request.
*/
type AssignShoppingListParams struct {

	// Body.
	Body *models.AssignShoppingListParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign shopping list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignShoppingListParams) WithDefaults() *AssignShoppingListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign shopping list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignShoppingListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign shopping list params
func (o *AssignShoppingListParams) WithTimeout(timeout time.Duration) *AssignShoppingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign shopping list params
func (o *AssignShoppingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign shopping list params
func (o *AssignShoppingListParams) WithContext(ctx context.Context) *AssignShoppingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign shopping list params
func (o *AssignShoppingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign shopping list params
func (o *AssignShoppingListParams) WithHTTPClient(client *http.Client) *AssignShoppingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign shopping list params
func (o *AssignShoppingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the assign shopping list params
func (o *AssignShoppingListParams) WithBody(body *models.AssignShoppingListParamsBody) *AssignShoppingListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the assign shopping list params
func (o *AssignShoppingListParams) SetBody(body *models.AssignShoppingListParamsBody) {
	o.Body = body
}

// WithID adds the id to the assign shopping list params
func (o *AssignShoppingListParams) WithID(id string) *AssignShoppingListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign shopping list params
func (o *AssignShoppingListParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignShoppingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
