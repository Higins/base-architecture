package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Higins/blocks/graph/generated"
)

// DummyMutation is the resolver for the dummyMutation field.
func (r *mutationResolver) DummyMutation(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// DummyMutation is the resolver for the dummyMutation field.
func (r *queryResolver) DummyMutation(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
