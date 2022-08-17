package gqlerrors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetGQLErrorType(err error) ErrorType {
	if errList, ok := err.(gqlerror.List); ok {
		gqlerr := &gqlerror.Error{}
		if errList.As(&gqlerr) && gqlerr.Extensions != nil {
			errorTypeString, isString := gqlerr.Extensions["errorType"].(string)
			if isString {
				return ErrorType(errorTypeString)
			}
		}
	}
	return ""
}
