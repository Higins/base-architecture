package di

import (
	userRepo "github.com/Higins/blocks/entity/user/repository"
	userUse "github.com/Higins/blocks/entity/user/usecase"
	"github.com/Higins/blocks/graph"
	"github.com/Higins/blocks/handlers"
	"github.com/Higins/blocks/router"
	"github.com/Higins/blocks/router/middlewares"
	"github.com/Higins/blocks/setting"
	"gorm.io/gorm"
)

func InitializeRouter(cfg setting.AppConfig, db *gorm.DB) *router.Router {
	userRepository := userRepo.NewUserRepository()
	authHandler := handlers.NewAuthHandler(cfg, db, userRepository)
	userUsecase := userUse.NewUserUseCase(db, cfg, userRepository, authHandler)
	resolver := graph.New(userUsecase)
	middleware := middlewares.NewMiddleware(db, authHandler, cfg)
	router := router.NewRouter(resolver, authHandler, middleware)
	return router
}
