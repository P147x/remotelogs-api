package main

import (
	"net/http"
	"time"

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
	r.Use(cors.New(cors.Config{
		AllowMethods:  []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-*"},
		ExposeHeaders: []string{"Content-Length"},
		AllowOrigins:  []string{"http://localhost", "*"},
		AllowWildcard: true,

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	configureV1(r)
}
