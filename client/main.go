package main

import (
	"client/client"
	"client/gqlerrors"
	"context"
	"fmt"
	"os"
)

func stringInput(message string) string {
	fmt.Print(message)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}

func main() {
	c := client.NewClient("http://localhost:8080/query")
	println("Welcome to the error types demo")
	firstName := stringInput("Enter first name: ")
	lastName := stringInput("Enter last name: ")
	_, err := c.ErrorTypesDemo(context.Background(), firstName, lastName)
	for err != nil {
		fmt.Printf("Received error from server: %s", err.Error())
		if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeBadFirstName {
			firstName = stringInput("Please enter another first name: ")
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeBadLastName {
			lastName = stringInput("Please enter another last name: ")
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeInternalServerError {
			fmt.Println("Yay! the server masks error messages successfully. You may check the server log to see the original error.")
			os.Exit(0)
		} else {
			// We can't identify the error in code, so we just exit
			os.Exit(1)
		}
		_, err = c.ErrorTypesDemo(context.Background(), firstName, lastName)
	}
}
