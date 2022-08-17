// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlerrors

import (
	"fmt"
	"io"
	"strconv"
)

type ErrorType string

const (
	ErrorTypeInternalServerError ErrorType = "InternalServerError"
	ErrorTypeBadFirstName        ErrorType = "BadFirstName"
	ErrorTypeBadLastName         ErrorType = "BadLastName"
)

var AllErrorType = []ErrorType{
	ErrorTypeInternalServerError,
	ErrorTypeBadFirstName,
	ErrorTypeBadLastName,
}

func (e ErrorType) IsValid() bool {
	switch e {
	case ErrorTypeInternalServerError, ErrorTypeBadFirstName, ErrorTypeBadLastName:
		return true
	}
	return false
}

func (e ErrorType) String() string {
	return string(e)
}

func (e *ErrorType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorType", str)
	}
	return nil
}

func (e ErrorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}