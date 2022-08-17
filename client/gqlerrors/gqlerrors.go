package gqlerrors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetGQLErrorType(err error) ErrorType {
	if errList, ok := err.(gqlerror.List); ok {
		gqlerr := &gqlerror.Error{}
		if errList.As(&gqlerr) && gqlerr.Extensions != nil {
			errorTypeStr, exists := gqlerr.Extensions["errorType"]
			if !exists {
				return ""
			}
			return ErrorType(errorTypeStr.(string))
		}
	}
	return ""
}
