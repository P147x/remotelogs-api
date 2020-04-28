package main

import (
	"github.com/gin-gonic/gin"
	"github.com/P147x/remotelogs-api/internal/controller"
	"net/http"
)

func ping(c *gin.Context) {
	c.String(http.StatusAccepted, "Pong")
}

func configureV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("ping", ping)
	controller.AuthentificationRoutes(v1.Group("/auth"))
}

func InitRouter(r *gin.Engine) {
	configureV1(r)
}