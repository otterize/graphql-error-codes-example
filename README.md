# GraphQL Error Codes Example for GoLang

This repo is an example of how to add error types (a.k.a error codes) to your [gqlgen](https://github.com/99designs/gqlgen) GraphQL server for Go.

The project demonstrates two features:
1. Intentional errors are returned with an error type, used by the client code to identify the error and behave accordingly.
2. Unexpected errors in the server are masked and replaced with "Internal Server Error". This prevents leak of sensitive information inside error messages.

For a further explanation, check out my blog post:
Who Moved My Error Codes? Adding error types to your GoLang GraphQL Server #TODO: add hyperlink

## How to use?

Run the project to see how it works, then have a look at the implementation to understand how it uses error types for its functionality.

Then, try disabling the error type and error masking features by editing the `server/main.go` file and comment out the `srv.SetErrorPresenter(GqlErrorPresenter)` line, and rerun the server. Repeat the demo to see how it is affected.

## How to run?
Run the server:
```shell
cd server
go run .
```

Then run the server by running the client:
```shell
cd client
go run .
```
You can also work against the server using the GraphQL playground at http://localhost:8080/playground

## Notes
- The GraphQL schema is in the `graphql/` dir. This is the only thing that is shared between the client and the server.
- Files that are named `*_gen.go` or `generated.go` are auto-generated. You can see the command that generated them in `generate.go`
