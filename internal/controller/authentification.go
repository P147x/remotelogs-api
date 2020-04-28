package controller

import (
	"github.com/P147x/remotelogs-api/internal/database"
	"github.com/P147x/remotelogs-api/internal/middleware"
	"github.com/P147x/remotelogs-api/internal/model"
	"github.com/gin-gonic/gin"

	"net/http"
)

//AuthentificationRoutes is used to store Authentification routes
func AuthentificationRoutes(r *gin.RouterGroup) {
	r.POST("/login", middleware.JWTMiddleware().LoginHandler)
	r.POST("/register", register)
}

//POST /api/v1/auth/register
func register(c *gin.Context) {
	var l model.User
	err := c.ShouldBind(&l)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
	} else if database.InsertUser(l) == false {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "ok"})
	}
}