package typederrors

import (
	"fmt"
	"server/model"
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

func InternalServerError(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeInternalServerError}
}

func BadFirstName(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeBadFirstName}
}

func BadLastName(messageToUserFormat string, args ...any) error {
	return &typedError{err: fmt.Errorf(messageToUserFormat, args...), errorType: model.ErrorTypeBadLastName}
}
