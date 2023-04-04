package router

import (
	"context"
	"net/http"

	"github.com/Higins/blocks/constants"
	"github.com/Higins/blocks/graph"
	"github.com/Higins/blocks/handlers"
	"github.com/Higins/blocks/router/middlewares"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Higins/blocks/graph/generated"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{"type": "router"})

type Router struct {
	resolversGraph *graph.Resolver
	authHandler    *handlers.AuthHandler
	middlewares    *middlewares.Middleware
}

func NewRouter(
	resolversGraph *graph.Resolver,
	authHandler *handlers.AuthHandler,
	middlewares *middlewares.Middleware,

) *Router {
	return &Router{
		resolversGraph: resolversGraph,
		authHandler:    authHandler,
		middlewares:    middlewares,
	}
}

func (r *Router) InitGraphqlServer(prodMode bool) *gin.Engine {
	/////////////
	// GIN INIT
	///////////
	gqlConfig := generated.Config{Resolvers: r.resolversGraph}

	gqlConfig.Directives.IsAuth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		userId, err := r.authHandler.GetUserFromContext(ctx)
		if err != nil || userId <= 0 {
			return nil, constants.ErrInternalError
		}
		return next(ctx)
	}

	grapql := gin.New()

	srv := handler.New(generated.NewExecutableSchema(gqlConfig))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	grapql.Use(AllowCrossOrigin)
	if !prodMode {
		grapql.Any("/play", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	}

	grapql.Any("/query", r.middlewares.SessionAuthenticateWithoutErrorAdmin, middlewares.GinContextToContextMiddleware(), gin.WrapH(srv))

	return grapql
}

/////////
// CORS
///////

func AllowCrossOrigin(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")

	origin := c.GetHeader("Origin")
	if origin != "" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	}

	// https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request
	if c.Request.Method == http.MethodOptions {
		c.Status(200)
		return
	}
}
