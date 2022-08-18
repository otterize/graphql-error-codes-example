package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"server/generated/server"
	"server/logic"
)

type Resolver struct{}

// ErrorTypesDemo is the resolver for the errorTypesDemo field.
func (r *queryResolver) ErrorTypesDemo(ctx context.Context, firstName string, lastName string) (string, error) {
	return logic.ErrorTypesDemo(firstName, lastName)
}

// ErrorMaskingDemo is the resolver for the errorMaskingDemo field.
func (r *queryResolver) ErrorMaskingDemo(ctx context.Context) (*bool, error) {
	return logic.ErrorMaskingDemo()
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
