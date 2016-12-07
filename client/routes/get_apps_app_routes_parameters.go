package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppsAppRoutesParams creates a new GetAppsAppRoutesParams object
// with the default values initialized.
func NewGetAppsAppRoutesParams() *GetAppsAppRoutesParams {
	var ()
	return &GetAppsAppRoutesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsAppRoutesParamsWithTimeout creates a new GetAppsAppRoutesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsAppRoutesParamsWithTimeout(timeout time.Duration) *GetAppsAppRoutesParams {
	var ()
	return &GetAppsAppRoutesParams{

		timeout: timeout,
	}
}

// NewGetAppsAppRoutesParamsWithContext creates a new GetAppsAppRoutesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsAppRoutesParamsWithContext(ctx context.Context) *GetAppsAppRoutesParams {
	var ()
	return &GetAppsAppRoutesParams{

		Context: ctx,
	}
}

/*GetAppsAppRoutesParams contains all the parameters to send to the API endpoint
for the get apps app routes operation typically these are written to a http.Request
*/
type GetAppsAppRoutesParams struct {

	/*App
	  Name of app for this set of routes.

	*/
	App string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get apps app routes params
func (o *GetAppsAppRoutesParams) WithTimeout(timeout time.Duration) *GetAppsAppRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps app routes params
func (o *GetAppsAppRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps app routes params
func (o *GetAppsAppRoutesParams) WithContext(ctx context.Context) *GetAppsAppRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps app routes params
func (o *GetAppsAppRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithApp adds the app to the get apps app routes params
func (o *GetAppsAppRoutesParams) WithApp(app string) *GetAppsAppRoutesParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the get apps app routes params
func (o *GetAppsAppRoutesParams) SetApp(app string) {
	o.App = app
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsAppRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
