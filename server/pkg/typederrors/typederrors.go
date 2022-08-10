package typederrors

import (
	"fmt"
	"server/pkg/model"
)

type typedError struct {
	err       error
	errorType model.ErrorType // error types are auto-generated from the schema
}

func (g typedError) Error() string {
	return g.err.Error()
}

func (g typedError) Unwrap() error {
	return g.err
}

func (g typedError) ErrorType() model.ErrorType {
	return g.errorType
}

func NotFound(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeNotFound}
}

func InternalServerError(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeInternalServerError}
}

func ForbiddenError(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeForbidden}
}

func BadRequest(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeBadRequest}
}

func ConflictError(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeConflict}
}
