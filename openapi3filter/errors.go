package openapi3filter

import (
	"fmt"

	"github.com/wicoady1/kin-openapi/openapi3"
)

var _ error = &RequestError{}

// RequestError is returned by ValidateRequest when request does not match OpenAPI spec
type RequestError struct {
	Input       *RequestValidationInput
	Parameter   *openapi3.Parameter
	RequestBody *openapi3.RequestBody
	Reason      string
	Err         error
}

func (err *RequestError) Error() string {
	reason := err.Reason
	if e := err.Err; e != nil {
		if len(reason) == 0 {
			reason = e.Error()
		} else {
			reason += ": " + e.Error()
		}
	}
	if v := err.Parameter; v != nil {
		return fmt.Sprintf("parameter %q in %s has an error: %s", v.Name, v.In, reason)
	} else if v := err.RequestBody; v != nil {
		return fmt.Sprintf("request body has an error: %s", reason)
	} else {
		return reason
	}
}

var _ error = &ResponseError{}

// ResponseError is returned by ValidateResponse when response does not match OpenAPI spec
type ResponseError struct {
	Input  *ResponseValidationInput
	Reason string
	Err    error
}

func (err *ResponseError) Error() string {
	reason := err.Reason
	if e := err.Err; e != nil {
		if len(reason) == 0 {
			reason = e.Error()
		} else {
			reason += ": " + e.Error()
		}
	}
	return reason
}

var _ error = &SecurityRequirementsError{}

// SecurityRequirementsError is returned by ValidateSecurityRequirements
// when no requirement is met.
type SecurityRequirementsError struct {
	SecurityRequirements openapi3.SecurityRequirements
	Errors               []error
}

func (err *SecurityRequirementsError) Error() string {
	return "Security requirements failed"
}
