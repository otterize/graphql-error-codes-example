# GraphQL Error Codes Example for GoLang

This repo is an example of how to add error types (a.k.a error codes) to your [gqlgen](https://github.com/99designs/gqlgen) GraphQL server for Go, and identify them on the client side.
Have a look at the source code to understand our implementation, and run the project to see how it works in run time.

For a further explanation, check out my blog post:
Who Moved My Error Codes? Adding error types to your GoLang GraphQL Server #TODO: add hyperlink

## How to run?
Run the server:
```shell
cd server/cmd
go run .
```

Then, run the client:
```shell
cd client/cmd
go run .
```

Add dogs to the pool, and then try to see their age. hint: password is `ilovedogs`, but try typing an incorrect password to see how the client handles the error.

Note that the server code is buggy on purpose, to demonstrate `Internal Server Error`. When receiving this error, have a look at the server log to see the original internal error.
