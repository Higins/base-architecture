package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	appModel "github.com/Higins/blocks/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*appModel.User, error) {
	return r.userUseCase.Login(ctx, email, password)
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	return r.userUseCase.Logout(ctx)
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, email string, password string, name string) (*appModel.User, error) {
	return r.userUseCase.Register(ctx, password, email, name)
}

// GetUser is the resolver for the get_user field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*appModel.User, error) {
	return r.userUseCase.Get(ctx, id)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*appModel.User, error) {
	return r.userUseCase.Me(ctx)
}
