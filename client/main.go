package main

import (
	"client/client"
	"client/gqlerrors"
	"context"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

func stringInput(message string) string {
	fmt.Print(message)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

func main() {
	c := client.NewClient("http://localhost:8080/query")

	demo := stringInput("Enter 1 for error types demo, or 2 for error masking demo: ")
	if demo == "1" {
		println(`Starting error types demo

In this demo the client is going ask for first and last name, send them as parameters to errorTypesDemo() query.
The server is going to validate the parameters lengths, and if one of them is too short, it will return an error with the appropriate message and type.
The client will check the returned error type, ask to type again only the bad parameter, and will try to call the server again.
`)
		_ = stringInput("Press enter to start\n")

		firstName := stringInput("Enter first name: ")
		lastName := stringInput("Enter last name: ")
		res, err := c.ErrorTypesDemo(context.Background(), firstName, lastName)
		for err != nil {
			println("\nServer returned an error:")
			fmt.Printf("  - error message: %s", err)
			errType := gqlerrors.GetGQLErrorType(err)
			fmt.Printf("  - error type: %s\n\n", lo.Ternary(len(errType) > 0, errType, "unspecified"))
			switch errType {
			case gqlerrors.ErrorTypeBadFirstName:
				firstName = stringInput("Please enter another first name: ")
			case gqlerrors.ErrorTypeBadLastName:
				lastName = stringInput("Please enter another last name: ")
			default:
				println("Server did not specify error type, so the client code can't handle the error.")
				break
			}
			res, err = c.ErrorTypesDemo(context.Background(), firstName, lastName)
		}
		fmt.Printf("\nServer returned: %s", res)
		println("\n\nDemo completed")

	} else if demo == "2" {
		println(`Starting error masking demo

In this demo the client is going to call the errorMaskingDemo() query, that causes an unexpected error in the server.
The server should mask the internal error, and we should receive a generic internal error message, with the "InternalServerError" error type.
`)
		_ = stringInput("Press enter to start\n")

		err := c.ErrorMaskingDemo(context.Background())
		errType := gqlerrors.GetGQLErrorType(err)
		if errType == gqlerrors.ErrorTypeInternalServerError {
			println("Success - the server masked the internal error. server returned:")
		} else {
			println("Oops - the server did not mask the internal error. server returned:")
		}
		fmt.Printf("  - error message: %s", err)
		fmt.Printf("  - error type: %s", lo.Ternary(len(errType) > 0, errType, "unspecified"))

		println("\n\nDemo completed")

	} else {
		println("invalid option, exiting")
		os.Exit(1)
	}
}
