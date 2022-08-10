package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg/resolvers"
	"server/pkg/server"
)

func main() {
	conf := server.Config{Resolvers: &resolvers.Resolver{}}
	srv := handler.NewDefaultServer(server.NewExecutableSchema(conf))
	srv.SetErrorPresenter(server.GqlErrorPresenter)
	http.Handle("/query", srv)
	logrus.Info("Starting server on port http://localhost:8080/query")
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
