package handlers

import (
	"context"

	"github.com/Higins/blocks/constants"
	"github.com/Higins/blocks/domain"
	"github.com/Higins/blocks/setting"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log = logrus.WithFields(logrus.Fields{"type": "auth_handler"})

type AuthHandler struct {
	db             *gorm.DB
	secureCookie   *securecookie.SecureCookie
	userRepository domain.UserRepository
	cfg            setting.AppConfig
}

func NewAuthHandler(
	cfg setting.AppConfig,
	db *gorm.DB,
	userRepository domain.UserRepository,
) *AuthHandler {
	return &AuthHandler{
		db:             db,
		secureCookie:   securecookie.New([]byte("ZlLoBmFTMa4a2k2LcvvQNSXWn2Hw6ci5"), []byte("OMN3OUeIO1P1uvXXNOdjPVwCpVMMg5Ip")),
		userRepository: userRepository,
		cfg:            cfg,
	}
}

func (a *AuthHandler) SecureCookie() *securecookie.SecureCookie {
	return a.secureCookie
}
func (a *AuthHandler) GenerateTokenForUser(userId int) (string, error) {
	return NewSignedTokenWithClaims(a.cfg.JwtSecret, map[string]interface{}{
		"uid": userId,
	})
}
func (a *AuthHandler) GetUserFromContext(ctx context.Context) (int, error) {
	ginContext := ctx.Value(constants.CtxGin).(*gin.Context)
	value, exist := ginContext.Get(constants.CtxUserID)
	userId, ok := value.(int)

	if !exist || !ok {
		return 0, constants.ErrUnauthorized
	}

	return userId, nil
}
