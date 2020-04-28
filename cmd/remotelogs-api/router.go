package main

import (
	"github.com/gin-gonic/gin"
	"github.com/P147x/remotelogs-api/internal/controller"
	"github.com/P147x/remotelogs-api/internal/middleware"
	"net/http"
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
	configureV1(r)
}