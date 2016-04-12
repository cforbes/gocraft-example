package gocraftsplit

import (
	"github.com/gocraft/web"
)

func GetRoutes() *web.Router {
	router := web.New(StatusHandler{})
	router.Get(`/status`, (*StatusHandler).Get)

	subRouter := router.Subrouter(RedisContext{}, `/api`)
	subRouter.Get(`/redis/ping`, (*RedisContext).Ping)

	return router
}