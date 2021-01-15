package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/s4kibs4mi/newschain-cache/middlewares"
	"net/http"
)

var router = echo.New()

// GetRouter returns the api router
func GetRouter() http.Handler {
	router.Pre(middleware.AddTrailingSlash())

	router.Use(middleware.Logger())
	router.Use(middlewares.Recovery())
	router.Use(middleware.RequestID())

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	registerV1Routes()
	return router
}

func registerV1Routes() {
	v1 := router.Group("/v1")
}
