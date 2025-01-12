// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chimesdkmeetings

import (
	"github.com/sotowang/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input parameters don't match the service's restrictions.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The client is permanently forbidden from making the request.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request exceeds the resource limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// One or more of the resources in the request does not exist in the system.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is currently unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// The user isn't authorized to request a resource.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeUnprocessableEntityException for service response error code
	// "UnprocessableEntityException".
	//
	// The request was well-formed but was unable to be followed due to semantic
	// errors.
	ErrCodeUnprocessableEntityException = "UnprocessableEntityException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ForbiddenException":           newErrorForbiddenException,
	"LimitExceededException":       newErrorLimitExceededException,
	"NotFoundException":            newErrorNotFoundException,
	"ServiceUnavailableException":  newErrorServiceUnavailableException,
	"UnauthorizedException":        newErrorUnauthorizedException,
	"UnprocessableEntityException": newErrorUnprocessableEntityException,
}
