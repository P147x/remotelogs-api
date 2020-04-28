package main

import (
	"net/http"

	"github.com/P147x/remotelogs-api/internal/controller"
	"github.com/P147x/remotelogs-api/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.String(http.StatusAccepted, "Pong")
}

func configureV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("ping", ping)
	v1.Use(middleware.JWTMiddleware().MiddlewareFunc())
	controller.AuthentificationRoutes(r.Group("/api/v1/auth"))
	controller.SystemRoutes(v1.Group("/system"))
}

func InitRouter(r *gin.Engine) {
	r.Use(cors.Default())
	configureV1(r)
}
