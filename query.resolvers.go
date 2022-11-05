package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"freg/graph"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (graph.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]graph.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context, after *graph.Cursor, first *int, before *graph.Cursor, last *int, where *graph.CarWhereInput) (*graph.CarConnection, error) {
	fmt.Println(where)
	return r.client.Car.Query().Paginate(ctx, after, first, before, last, graph.WithCarFilter(where.Filter))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *graph.Cursor, first *int, before *graph.Cursor, last *int, where *graph.GroupWhereInput) (*graph.GroupConnection, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	d, err := getAllChildren(ctx, tx, "2")
	if err != nil {
		return nil, err
	}

	fmt.Println(d)
	tx.Commit()
	return r.client.Group.Query().Paginate(ctx, after, first, before, last, graph.WithGroupFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *graph.Cursor, first *int, before *graph.Cursor, last *int, where *graph.UserWhereInput) (*graph.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, graph.WithUserFilter(where.Filter))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// UserWhereInput returns UserWhereInputResolver implementation.
func (r *Resolver) UserWhereInput() UserWhereInputResolver { return &userWhereInputResolver{r} }

type queryResolver struct{ *Resolver }
type userWhereInputResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func getAllChildren(ctx context.Context, tx *graph.Tx, id string) ([]string, error) {
	g, err := tx.Group.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	children, err := tx.Group.QueryChildren(g).All(ctx)
	if err != nil {
		return nil, err
	}
	toReturn := make([]string, 0, len(children))
	for _, c := range children {
		toReturn = append(toReturn, c.ID)
		child, err := getAllChildren(ctx, tx, c.ID)
		if err != nil {
			return nil, err
		}
		if len(child) > 0 {
			toReturn = append(toReturn, child...)
		}
	}

	return toReturn, nil
}
