package middlewares

import (
	"context"

	"github.com/Higins/blocks/constants"
	"github.com/Higins/blocks/handlers"
	"github.com/Higins/blocks/setting"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Middleware struct {
	cfg         setting.AppConfig
	db          *gorm.DB
	authHandler *handlers.AuthHandler
}

func NewMiddleware(
	db *gorm.DB,
	authHandler *handlers.AuthHandler,
	cfg setting.AppConfig,
) *Middleware {
	return &Middleware{
		db:          db,
		authHandler: authHandler,
		cfg:         cfg,
	}
}

var log = logrus.WithFields(logrus.Fields{"type": "middleware"})

func (m *Middleware) SessionAuthenticateWithoutErrorAdmin(c *gin.Context) {
	session, err := c.Cookie("session")
	if session != "" {
		if err == nil {
			value := make(map[string]int)
			if err = m.authHandler.SecureCookie().Decode("session", session, &value); err == nil {
				c.Set(constants.CtxUserID, value["id"])
				c.Next()
				return
			}
		}
		_ = c.Error(err)
		log.Info("User error: ", err)
		return
	}
	c.Next()

}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), constants.CtxGin, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetGinContextFromContext(c context.Context) (res *gin.Context) {
	res, _ = c.Value(constants.CtxGin).(*gin.Context)
	return
}
