package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/casualjim/patmosdb/models"
)

/*PutEntryCreated entry was created

swagger:response putEntryCreated
*/
type PutEntryCreated struct {
	/*The version of this entry
	  Required: true
	*/
	Etag string `json:"Etag"`
	/*the location to get the newly created entry
	  Required: true
	*/
	Location strfmt.URI `json:"Location"`
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`
}

// NewPutEntryCreated creates PutEntryCreated with default headers values
func NewPutEntryCreated() *PutEntryCreated {
	return &PutEntryCreated{}
}

// WithEtag adds the etag to the put entry created response
func (o *PutEntryCreated) WithEtag(etag string) *PutEntryCreated {
	o.Etag = etag
	return o
}

// SetEtag sets the etag to the put entry created response
func (o *PutEntryCreated) SetEtag(etag string) {
	o.Etag = etag
}

// WithLocation adds the location to the put entry created response
func (o *PutEntryCreated) WithLocation(location strfmt.URI) *PutEntryCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the put entry created response
func (o *PutEntryCreated) SetLocation(location strfmt.URI) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the put entry created response
func (o *PutEntryCreated) WithXRequestID(xRequestID string) *PutEntryCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put entry created response
func (o *PutEntryCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *PutEntryCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Etag

	etag := o.Etag
	if etag != "" {
		rw.Header().Set("Etag", etag)
	}

	// response header Location

	location := o.Location.String()
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(201)
}

/*PutEntryNoContent entry was updated

swagger:response putEntryNoContent
*/
type PutEntryNoContent struct {
	/*The version of this entry
	  Required: true
	*/
	ETag string `json:"ETag"`
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`
}

// NewPutEntryNoContent creates PutEntryNoContent with default headers values
func NewPutEntryNoContent() *PutEntryNoContent {
	return &PutEntryNoContent{}
}

// WithETag adds the eTag to the put entry no content response
func (o *PutEntryNoContent) WithETag(eTag string) *PutEntryNoContent {
	o.ETag = eTag
	return o
}

// SetETag sets the eTag to the put entry no content response
func (o *PutEntryNoContent) SetETag(eTag string) {
	o.ETag = eTag
}

// WithXRequestID adds the xRequestId to the put entry no content response
func (o *PutEntryNoContent) WithXRequestID(xRequestID string) *PutEntryNoContent {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put entry no content response
func (o *PutEntryNoContent) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *PutEntryNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header ETag

	eTag := o.ETag
	if eTag != "" {
		rw.Header().Set("ETag", eTag)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(204)
}

/*PutEntryNotFound The entry was not found

swagger:response putEntryNotFound
*/
type PutEntryNotFound struct {
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutEntryNotFound creates PutEntryNotFound with default headers values
func NewPutEntryNotFound() *PutEntryNotFound {
	return &PutEntryNotFound{}
}

// WithXRequestID adds the xRequestId to the put entry not found response
func (o *PutEntryNotFound) WithXRequestID(xRequestID string) *PutEntryNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put entry not found response
func (o *PutEntryNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put entry not found response
func (o *PutEntryNotFound) WithPayload(payload *models.Error) *PutEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put entry not found response
func (o *PutEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutEntryConflict there is a version mismatch for the entry

swagger:response putEntryConflict
*/
type PutEntryConflict struct {
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutEntryConflict creates PutEntryConflict with default headers values
func NewPutEntryConflict() *PutEntryConflict {
	return &PutEntryConflict{}
}

// WithXRequestID adds the xRequestId to the put entry conflict response
func (o *PutEntryConflict) WithXRequestID(xRequestID string) *PutEntryConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put entry conflict response
func (o *PutEntryConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put entry conflict response
func (o *PutEntryConflict) WithPayload(payload *models.Error) *PutEntryConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put entry conflict response
func (o *PutEntryConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEntryConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutEntryDefault Error

swagger:response putEntryDefault
*/
type PutEntryDefault struct {
	_statusCode int
	/*The request id this is a response to
	  Required: true
	*/
	XRequestID string `json:"X-Request-Id"`

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutEntryDefault creates PutEntryDefault with default headers values
func NewPutEntryDefault(code int) *PutEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &PutEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put entry default response
func (o *PutEntryDefault) WithStatusCode(code int) *PutEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put entry default response
func (o *PutEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the put entry default response
func (o *PutEntryDefault) WithXRequestID(xRequestID string) *PutEntryDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the put entry default response
func (o *PutEntryDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the put entry default response
func (o *PutEntryDefault) WithPayload(payload *models.Error) *PutEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put entry default response
func (o *PutEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
