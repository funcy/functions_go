package version

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

// NewGetVersionParams creates a new GetVersionParams object
// with the default values initialized.
func NewGetVersionParams() *GetVersionParams {

	return &GetVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionParamsWithTimeout creates a new GetVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionParamsWithTimeout(timeout time.Duration) *GetVersionParams {

	return &GetVersionParams{

		timeout: timeout,
	}
}

// NewGetVersionParamsWithContext creates a new GetVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionParamsWithContext(ctx context.Context) *GetVersionParams {

	return &GetVersionParams{

		Context: ctx,
	}
}

/*GetVersionParams contains all the parameters to send to the API endpoint
for the get version operation typically these are written to a http.Request
*/
type GetVersionParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get version params
func (o *GetVersionParams) WithTimeout(timeout time.Duration) *GetVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version params
func (o *GetVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version params
func (o *GetVersionParams) WithContext(ctx context.Context) *GetVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version params
func (o *GetVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
