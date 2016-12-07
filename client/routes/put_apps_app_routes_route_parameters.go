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

	"github.com/iron-io/functions_go/models"
)

// NewPutAppsAppRoutesRouteParams creates a new PutAppsAppRoutesRouteParams object
// with the default values initialized.
func NewPutAppsAppRoutesRouteParams() *PutAppsAppRoutesRouteParams {
	var ()
	return &PutAppsAppRoutesRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAppsAppRoutesRouteParamsWithTimeout creates a new PutAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAppsAppRoutesRouteParamsWithTimeout(timeout time.Duration) *PutAppsAppRoutesRouteParams {
	var ()
	return &PutAppsAppRoutesRouteParams{

		timeout: timeout,
	}
}

// NewPutAppsAppRoutesRouteParamsWithContext creates a new PutAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAppsAppRoutesRouteParamsWithContext(ctx context.Context) *PutAppsAppRoutesRouteParams {
	var ()
	return &PutAppsAppRoutesRouteParams{

		Context: ctx,
	}
}

/*PutAppsAppRoutesRouteParams contains all the parameters to send to the API endpoint
for the put apps app routes route operation typically these are written to a http.Request
*/
type PutAppsAppRoutesRouteParams struct {

	/*App
	  name of the app.

	*/
	App string
	/*Body
	  One route to post.

	*/
	Body *models.RouteWrapper
	/*Route
	  route path.

	*/
	Route string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) WithTimeout(timeout time.Duration) *PutAppsAppRoutesRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) WithContext(ctx context.Context) *PutAppsAppRoutesRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithApp adds the app to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) WithApp(app string) *PutAppsAppRoutesRouteParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) SetApp(app string) {
	o.App = app
}

// WithBody adds the body to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) WithBody(body *models.RouteWrapper) *PutAppsAppRoutesRouteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) SetBody(body *models.RouteWrapper) {
	o.Body = body
}

// WithRoute adds the route to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) WithRoute(route string) *PutAppsAppRoutesRouteParams {
	o.SetRoute(route)
	return o
}

// SetRoute adds the route to the put apps app routes route params
func (o *PutAppsAppRoutesRouteParams) SetRoute(route string) {
	o.Route = route
}

// WriteToRequest writes these params to a swagger request
func (o *PutAppsAppRoutesRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.RouteWrapper)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param route
	if err := r.SetPathParam("route", o.Route); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
