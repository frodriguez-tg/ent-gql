package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"freg/ent"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input UserCreateInput) (*ent.User, error) {
	return r.client.User.Create().
		SetAge(input.Age).
		SetNillableName(input.Name).
		Save(ctx)
}

// CreateCar is the resolver for the createCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input CarCreateInput) (*ent.Car, error) {
	return r.client.Car.Create().
		SetModel(input.Model).
		Save(ctx)
}

// AddCarsToUser is the resolver for the addCarsToUser field.
func (r *mutationResolver) AddCarsToUser(ctx context.Context, input CarsToUserInput) (*GenericResponse, error) {
	if err := r.client.User.UpdateOneID(input.UserID).AddCarIDs(input.CarIDs...).Exec(ctx); err != nil {
		return nil, err
	}
	return &GenericResponse{Status: "ok"}, nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input GroupCreateInput) (*ent.Group, error) {
	return r.client.Group.Create().
		SetName(input.Name).
		Save(ctx)
}

// AddUsersToGroup is the resolver for the addUsersToGroup field.
func (r *mutationResolver) AddUsersToGroup(ctx context.Context, input UsersToGroup) (*GenericResponse, error) {
	if err := r.client.Group.UpdateOneID(input.GroupID).AddUserIDs(input.UserIDs...).Exec(ctx); err != nil {
		return nil, err
	}

	return &GenericResponse{
		Status: "ok",
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
