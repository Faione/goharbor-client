// Code generated by go-swagger; DO NOT EDIT.

package purge

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

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// NewUpdatePurgeScheduleParams creates a new UpdatePurgeScheduleParams object
// with the default values initialized.
func NewUpdatePurgeScheduleParams() *UpdatePurgeScheduleParams {
	var ()
	return &UpdatePurgeScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePurgeScheduleParamsWithTimeout creates a new UpdatePurgeScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePurgeScheduleParamsWithTimeout(timeout time.Duration) *UpdatePurgeScheduleParams {
	var ()
	return &UpdatePurgeScheduleParams{

		timeout: timeout,
	}
}

// NewUpdatePurgeScheduleParamsWithContext creates a new UpdatePurgeScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePurgeScheduleParamsWithContext(ctx context.Context) *UpdatePurgeScheduleParams {
	var ()
	return &UpdatePurgeScheduleParams{

		Context: ctx,
	}
}

// NewUpdatePurgeScheduleParamsWithHTTPClient creates a new UpdatePurgeScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePurgeScheduleParamsWithHTTPClient(client *http.Client) *UpdatePurgeScheduleParams {
	var ()
	return &UpdatePurgeScheduleParams{
		HTTPClient: client,
	}
}

/*UpdatePurgeScheduleParams contains all the parameters to send to the API endpoint
for the update purge schedule operation typically these are written to a http.Request
*/
type UpdatePurgeScheduleParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Schedule
	  The purge job's schedule, it is a json object. ｜
	The sample format is ｜
	{"parameters":{"audit_retention_hour":168,"dry_run":true, "include_operations":"create,delete,pull"},"schedule":{"type":"Hourly","cron":"0 0 * * * *"}} ｜
	the include_operation should be a comma separated string, e.g. create,delete,pull, if it is empty, no operation will be purged.


	*/
	Schedule *model.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update purge schedule params
func (o *UpdatePurgeScheduleParams) WithTimeout(timeout time.Duration) *UpdatePurgeScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update purge schedule params
func (o *UpdatePurgeScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update purge schedule params
func (o *UpdatePurgeScheduleParams) WithContext(ctx context.Context) *UpdatePurgeScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update purge schedule params
func (o *UpdatePurgeScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update purge schedule params
func (o *UpdatePurgeScheduleParams) WithHTTPClient(client *http.Client) *UpdatePurgeScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update purge schedule params
func (o *UpdatePurgeScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the update purge schedule params
func (o *UpdatePurgeScheduleParams) WithXRequestID(xRequestID *string) *UpdatePurgeScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update purge schedule params
func (o *UpdatePurgeScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the update purge schedule params
func (o *UpdatePurgeScheduleParams) WithSchedule(schedule *model.Schedule) *UpdatePurgeScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the update purge schedule params
func (o *UpdatePurgeScheduleParams) SetSchedule(schedule *model.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePurgeScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
