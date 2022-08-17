package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"server/server"
	"server/typederrors"
)

// ErrorTypesDemo is the resolver for the errorTypesDemo field.
func (r *queryResolver) ErrorTypesDemo(ctx context.Context, firstName string, lastName string) (string, error) {
	if len(firstName) < 2 {
		return "", typederrors.BadFirstName("first name is too short")
	}
	if len(lastName) < 2 {
		return "", typederrors.BadLastName("last name is too short")
	}
	return "", fmt.Errorf("dear %s %s, this is an internal error from the server. if you see it on the client, "+
		"it means the server does not mask the messages of internal errors, and may leak sensitive information", firstName, lastName)
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
