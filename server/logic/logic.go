package logic

import (
	"errors"
	"fmt"
	"server/generated/typederrors"
)

func ErrorTypesDemo(firstName string, lastName string) (string, error) {
	if len(firstName) < 2 {
		return "", typederrors.BadFirstName("first name is too short")
	}
	if len(lastName) < 2 {
		return "", typederrors.BadLastName("last name is too short")
	}
	return fmt.Sprintf("Well done %s %s", firstName, lastName), nil
}

func ErrorMaskingDemo() (*bool, error) {
	return nil, errors.New("if you see this error message on the client, there is an information leak")
}
