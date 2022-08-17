# GraphQL Error Codes Example for GoLang

This repo is an example of how to add error types (a.k.a error codes) to your [gqlgen](https://github.com/99designs/gqlgen) GraphQL server for Go, and identify them on the client side.
Run the project to see how it works, then have a look at the implementation to understand how it uses error types for its functionality.

For a further explanation, check out my blog post:
Who Moved My Error Codes? Adding error types to your GoLang GraphQL Server #TODO: add hyperlink

## What's inside?
The server exposes a graphql Query called `errorTypesDemo()`, that receives first and last name.
Each name is validated to be longer than 2 characters, and in case it doesn't, an error with the matching error type is returned.

Assuming the validation passed, an internal error is intentionally thrown in the server, to demonstrate the masking of unexpected errors with "InternalServerError".
For comparison, try going to the server's `main.go` file and comment out the `srv.SetErrorPresenter(GqlErrorPresenter)` line, to see the leak of the internal error message to the client.

## How to run?
Run the server:
```shell
cd server/cmd
go run .
```

You can work against the server using the GraphQL playground at http://localhost:8080/playground 
Or by running the client:
```shell
cd client/cmd
go run .
```

## Notes
- The GraphQL schema is in the `graphql/` dir. This is the only thing that is shared between the client and the server.
- Files that are named `*_gen.go` or `generated.go` are auto-generated. You can see the command that generated them in `generate.go`
