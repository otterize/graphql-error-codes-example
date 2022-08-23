package main

import (
	"client/client"
	"client/gqlerrors"
	"context"
	"fmt"
	"github.com/samber/lo"
)

func main() {
	c := client.NewClient("http://localhost:8080/query")

	// Error types demo
	println(ErrorTypesDemoDescription)
	firstName := "Ms"
	lastName := "G"
	retry := false

	fmt.Printf("Calling server with name \"%s %s\"...", firstName, lastName)
	res, err := c.ErrorTypesDemo(context.Background(), firstName, lastName)
	if err != nil {
		println("\nServer returned an error:")
		fmt.Printf("  - error message: %s", err)
		errType := gqlerrors.GetGQLErrorType(err)
		fmt.Printf("  - error type: %s\n\n", lo.Ternary(len(errType) > 0, errType, "unspecified"))
		switch errType {
		case gqlerrors.ErrorTypeBadFirstName:
			fmt.Printf("First name %s is bad, will try again with \"Ariana\" instead\n", firstName)
			firstName = "Ariana"
			retry = true
		case gqlerrors.ErrorTypeBadLastName:
			fmt.Printf("Last name %s is bad, will try again with \"Grande\" instead\n", lastName)
			lastName = "Grande"
			retry = true
		default:
			println("Server did not specify error type, so the client code can't handle the error.")
			break
		}
	} else {
		fmt.Printf("\nServer returned: %s", res)
	}

	if retry {
		fmt.Printf("Retrying to call server with new name \"%s %s\"...", firstName, lastName)
		res, err := c.ErrorTypesDemo(context.Background(), firstName, lastName)
		if err != nil {
			println("\nServer returned an error:")
			fmt.Printf("  - error message: %s", err)
			errType := gqlerrors.GetGQLErrorType(err)
			fmt.Printf("  - error type: %s\n\n", lo.Ternary(len(errType) > 0, errType, "unspecified"))
		} else {
			fmt.Printf("\nServer returned: %s", res)
		}
	}
	println("\n\nError types demo completed\n")

	// Error masking Demo
	println(ErrorMaskingDemoDescription)
	err = c.ErrorMaskingDemo(context.Background())
	errType := gqlerrors.GetGQLErrorType(err)
	if errType == gqlerrors.ErrorTypeInternalServerError {
		println("Success - the server masked the internal error. server returned:")
	} else {
		println("Oops - the server did not mask the internal error. server returned:")
	}
	fmt.Printf("  - error message: %s", err)
	fmt.Printf("  - error type: %s", lo.Ternary(len(errType) > 0, errType, "unspecified"))

	println("\n\nError masking demo completed")
}

const (
	ErrorTypesDemoDescription = `### Error types demo ###

In this demo the client send first and last names to errorTypesDemo() query.
The server is going to validate the parameters lengths, and if one of them is too short, it will return an error with the appropriate message and type.
The client will check the returned error type, and replace the bad name, and retry.
Starting demo...
`
	ErrorMaskingDemoDescription = `### Error masking demo ###

In this demo the client is going to call the errorMaskingDemo() query, that causes an unexpected error in the server.
The server should mask the internal error, and we should receive a generic internal error message, with the "InternalServerError" error type.
Starting demo...
`
)
