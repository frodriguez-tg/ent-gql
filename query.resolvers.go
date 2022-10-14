package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"freg/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.CarWhereInput) (*ent.CarConnection, error) {
	fmt.Println(where)
	return r.client.Car.Query().Paginate(ctx, after, first, before, last, ent.WithCarFilter(where.Filter))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().Paginate(ctx, after, first, before, last, ent.WithGroupFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
