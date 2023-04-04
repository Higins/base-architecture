package domain

import (
	"context"

	appModel "github.com/Higins/blocks/model"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Get(ctx context.Context, id int) (*appModel.User, error)
	Logout(ctx context.Context) (bool, error)
	Login(ctx context.Context, email string, password string) (*appModel.User, error)
	Me(ctx context.Context) (*appModel.User, error)
	Register(ctx context.Context, name, password, email string) (*appModel.User, error)
}

type UserRepository interface {
	Get(tx *gorm.DB, userId int) (*appModel.User, error)
	GetUserByEmail(tx *gorm.DB, email string) (*appModel.User, error)
}
