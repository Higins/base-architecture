package usecase

import (
	"context"
	"errors"

	"github.com/Higins/blocks/constants"
	"github.com/Higins/blocks/domain"
	"github.com/Higins/blocks/handlers"
	appModel "github.com/Higins/blocks/model"
	"github.com/Higins/blocks/setting"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var log = logrus.WithFields(logrus.Fields{"type": "user_usecase"})

type userUseCase struct {
	db          *gorm.DB
	cfg         setting.AppConfig
	userRepo    domain.UserRepository
	authHandler *handlers.AuthHandler
}

func NewUserUseCase(
	db *gorm.DB,
	cfg setting.AppConfig,
	userRepo domain.UserRepository,
	authHandler *handlers.AuthHandler,

) domain.UserUsecase {
	return &userUseCase{
		db:          db,
		userRepo:    userRepo,
		authHandler: authHandler,
		cfg:         cfg,
	}
}

func (u *userUseCase) Get(ctx context.Context, id int) (*appModel.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()
	user, err := u.userRepo.Get(tx, id)
	if err != nil {
		return nil, constants.ErrInternalError
	}
	tx.Commit()
	return user, nil
}
func (u *userUseCase) Login(ctx context.Context, email string, password string) (*appModel.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()
	user, err := u.userRepo.GetUserByEmail(tx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrInvalidEmailOrPassword
		}
		return nil, constants.ErrDatabaseError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, constants.ErrInvalidEmailOrPassword
	}

	value := map[string]int{
		"id": user.ID,
	}
	ginContext := ctx.Value(constants.CtxGin).(*gin.Context)

	encoded, err := u.authHandler.SecureCookie().Encode("session", value)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	tx.Commit()
	ginContext.SetCookie("session", encoded, 15811200, "/", "", false, true)

	return user, nil
}

func (u *userUseCase) Logout(ctx context.Context) (bool, error) {
	gin := ctx.Value(constants.CtxGinContext).(*gin.Context)
	gin.SetCookie("session", "", -1, "/", u.cfg.Domain, true, true)

	return true, nil
}
func (u *userUseCase) Me(ctx context.Context) (*appModel.User, error) {
	usr, err := u.authHandler.GetUserFromContext(ctx)
	if err != nil {
		return nil, constants.ErrInternalError
	}
	tx := u.db.Begin()
	defer tx.Rollback()
	user, err := u.userRepo.Get(tx, usr)
	if err != nil {
		return nil, constants.ErrInternalError
	}
	tx.Commit()
	return user, nil
}
func (u *userUseCase) Register(ctx context.Context, name, password, email string) (*appModel.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()
	var user appModel.User
	user.Email = email
	user.Name = name
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return nil, constants.ErrInternalError
	}
	user.Password = string(bytes)

	err = tx.Save(&user).Error
	if err != nil {
		log.Error(err)
		return nil, constants.ErrInternalError
	}
	tx.Commit()
	return &user, nil
}
