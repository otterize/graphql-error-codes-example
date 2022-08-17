package gqlerrors

//go:generate go run github.com/99designs/gqlgen@v0.17.13
// we only need models_gen for the enum, so we delete the server code
//go:generate rm generated.go
