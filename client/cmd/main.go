package main

import (
	"client/pkg/client"
	"client/pkg/gqlerrors"
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

func dogInput() client.DogInput {
	fmt.Println("Enter dog info")
	return client.DogInput{Name: stringInput("Name: "),
		Breed:    stringInput("Breed: "),
		Birthday: stringInput("Birthday: "),
	}
}

func addDog(c *client.Client) {
	dogInfo := dogInput()
	err := c.AddDog(context.Background(), dogInfo)
	for err != nil {
		if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeConflict {
			dogInfo.Name = stringInput("Name is already taken, enter another name: ")
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeBadRequest {
			fmt.Printf("%s. lets try again:", err)
			dogInfo = dogInput()
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeInternalServerError {
			fmt.Println("Oops, something is wrong with the server, please try again later")
			os.Exit(1)
		} else {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
		err = c.AddDog(context.Background(), dogInfo)
	}
	fmt.Printf("Dog %s added successfully\n", dogInfo.Name)
}

func printDogAge(c *client.Client) {
	dogName := stringInput("Enter dog name: ")
	password := stringInput("Enter password: ")
	info, err := c.GetDog(context.Background(), dogName, password)
	for err != nil {
		if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeForbidden {
			password = stringInput("Wrong password, enter again: ")
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeNotFound {
			fmt.Printf("Not dog named %s\n", dogName)
			return
		} else if gqlerrors.GetGQLErrorType(err) == gqlerrors.ErrorTypeInternalServerError {
			fmt.Println("Oops, something is wrong with the server, please try again later")
			os.Exit(1)
		} else {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
		info, err = c.GetDog(context.Background(), dogName, password)
	}
	fmt.Printf("%s is %d years old!\n", info.Name, info.Age)
}

func main() {
	c := client.NewClient("http://localhost:8080/query")
	println("Welcome to DogShell")
	for true {
		action := stringInput(`Enter 1 to add a dog, or 2 to see the age of a dog: `)
		if action == "1" {
			addDog(c)
		} else if action == "2" {
			printDogAge(c)
		} else {
			println("invalid input")
		}
	}
}
