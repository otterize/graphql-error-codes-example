package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
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
	return fmt.Sprintf("Well done %s %s", firstName, lastName), nil
}

// ErrorMaskingDemo is the resolver for the errorMaskingDemo field.
func (r *queryResolver) ErrorMaskingDemo(ctx context.Context) (*bool, error) {
	return nil, errors.New("if you see this error message on the client, there is an information leak")
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
