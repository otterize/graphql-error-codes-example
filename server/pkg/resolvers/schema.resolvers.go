package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"server/pkg/model"
	"server/pkg/server"
	"server/pkg/typederrors"
	"strings"
)

// AddDog is the resolver for the addDog field.
func (r *mutationResolver) AddDog(ctx context.Context, dogInput model.DogInput) (*bool, error) {
	if dogInput.Breed == "" {
		return nil, typederrors.BadRequest("dog breed cannot be empty")
	}
	if _, ok := dogsDB[dogInput.Name]; ok {
		return nil, typederrors.ConflictError("dog named %s already exists", dogInput.Name)
	}
	dogsDB[dogInput.Name] = dogInput
	return nil, nil
}

// UpdateDog is the resolver for the updateDog field.
func (r *mutationResolver) UpdateDog(ctx context.Context, dogInput model.DogInput) (*bool, error) {
	if _, ok := dogsDB[dogInput.Name]; !ok {
		return nil, typederrors.NotFound("Dog %s not found", dogInput.Name)
	}
	if dogInput.Breed == "" {
		return nil, typederrors.BadRequest("Dog breed cannot be empty")
	}
	dogsDB[dogInput.Name] = dogInput
	return nil, nil
}

// DeleteDog is the resolver for the deleteDog field.
func (r *mutationResolver) DeleteDog(ctx context.Context, name string) (*bool, error) {
	if strings.ToLower(name) == "tobi" {
		return nil, typederrors.ForbiddenError("No one can delete tobi")
	}
	if _, ok := dogsDB[name]; ok {
		delete(dogsDB, name)
	} else {
		return nil, typederrors.NotFound("Dog %s not found", name)
	}
	return nil, nil
}

// Dog is the resolver for the dog field.
func (r *queryResolver) Dog(ctx context.Context, name string, password string) (*model.DogInfo, error) {
	if password != readPassword {
		return nil, typederrors.ForbiddenError("Wrong password")
	}
	if _, ok := dogsDB[name]; !ok {
		return nil, typederrors.NotFound("Dog %s not found", name)
	}
	age, err := calculateAge(dogsDB[name].Birthday)
	if err != nil {
		return nil, err
	}
	return &model.DogInfo{Name: dogsDB[name].Name,
		Breed:    dogsDB[name].Breed,
		Birthday: dogsDB[name].Birthday,
		Age:      age,
	}, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
