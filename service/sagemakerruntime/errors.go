// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"github.com/sotowang/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalDependencyException for service response error code
	// "InternalDependencyException".
	//
	// Your request caused an exception with an internal dependency. Contact customer
	// support.
	ErrCodeInternalDependencyException = "InternalDependencyException"

	// ErrCodeInternalFailure for service response error code
	// "InternalFailure".
	//
	// An internal failure occurred.
	ErrCodeInternalFailure = "InternalFailure"

	// ErrCodeModelError for service response error code
	// "ModelError".
	//
	// Model (owned by the customer in the container) returned 4xx or 5xx error
	// code.
	ErrCodeModelError = "ModelError"

	// ErrCodeModelNotReadyException for service response error code
	// "ModelNotReadyException".
	//
	// Either a serverless endpoint variant's resources are still being provisioned,
	// or a multi-model endpoint is still downloading or loading the target model.
	// Wait and try your request again.
	ErrCodeModelNotReadyException = "ModelNotReadyException"

	// ErrCodeServiceUnavailable for service response error code
	// "ServiceUnavailable".
	//
	// The service is unavailable. Try your call again.
	ErrCodeServiceUnavailable = "ServiceUnavailable"

	// ErrCodeValidationError for service response error code
	// "ValidationError".
	//
	// Inspect your request and try again.
	ErrCodeValidationError = "ValidationError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalDependencyException": newErrorInternalDependencyException,
	"InternalFailure":             newErrorInternalFailure,
	"ModelError":                  newErrorModelError,
	"ModelNotReadyException":      newErrorModelNotReadyException,
	"ServiceUnavailable":          newErrorServiceUnavailable,
	"ValidationError":             newErrorValidationError,
}
