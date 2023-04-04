package graph

import (
	"github.com/Higins/blocks/domain"
)

//go:generate go run ../scripts/gqlgen.go

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase domain.UserUsecase
}

func New(
	userUseCase domain.UserUsecase,

) *Resolver {
	return &Resolver{
		userUseCase: userUseCase,
	}
}
