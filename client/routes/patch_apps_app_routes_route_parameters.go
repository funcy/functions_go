package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// NewPatchAppsAppRoutesRouteParams creates a new PatchAppsAppRoutesRouteParams object
// with the default values initialized.
func NewPatchAppsAppRoutesRouteParams() *PatchAppsAppRoutesRouteParams {
	var ()
	return &PatchAppsAppRoutesRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppsAppRoutesRouteParamsWithTimeout creates a new PatchAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAppsAppRoutesRouteParamsWithTimeout(timeout time.Duration) *PatchAppsAppRoutesRouteParams {
	var ()
	return &PatchAppsAppRoutesRouteParams{

		timeout: timeout,
	}
}

// NewPatchAppsAppRoutesRouteParamsWithContext creates a new PatchAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAppsAppRoutesRouteParamsWithContext(ctx context.Context) *PatchAppsAppRoutesRouteParams {
	var ()
	return &PatchAppsAppRoutesRouteParams{

		Context: ctx,
	}
}

// NewPatchAppsAppRoutesRouteParamsWithHTTPClient creates a new PatchAppsAppRoutesRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAppsAppRoutesRouteParamsWithHTTPClient(client *http.Client) *PatchAppsAppRoutesRouteParams {
	var ()
	return &PatchAppsAppRoutesRouteParams{
		HTTPClient: client,
	}
}

/*PatchAppsAppRoutesRouteParams contains all the parameters to send to the API endpoint
for the patch apps app routes route operation typically these are written to a http.Request
*/
type PatchAppsAppRoutesRouteParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithTimeout(timeout time.Duration) *PatchAppsAppRoutesRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithContext(ctx context.Context) *PatchAppsAppRoutesRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithHTTPClient(client *http.Client) *PatchAppsAppRoutesRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithApp(app string) *PatchAppsAppRoutesRouteParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetApp(app string) {
	o.App = app
}

// WithBody adds the body to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithBody(body *models.RouteWrapper) *PatchAppsAppRoutesRouteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetBody(body *models.RouteWrapper) {
	o.Body = body
}

// WithRoute adds the route to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) WithRoute(route string) *PatchAppsAppRoutesRouteParams {
	o.SetRoute(route)
	return o
}

// SetRoute adds the route to the patch apps app routes route params
func (o *PatchAppsAppRoutesRouteParams) SetRoute(route string) {
	o.Route = route
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppsAppRoutesRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
