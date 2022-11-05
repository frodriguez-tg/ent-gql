package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"freg/graph"
	"freg/graph/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
)

// Permissions is the resolver for the permissions field.
func (r *userWhereInputResolver) Permissions(ctx context.Context, obj *graph.UserWhereInput, data []string) error {
	if len(data) == 0 {
		return nil
	}
	for _, permission := range data {
		obj.AddPredicates(func(s *sql.Selector) {
			s.Where(sqljson.ValueContains(user.FieldPermissions, permission))
		})
	}
	return nil
}
