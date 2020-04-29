package main

import (
	db "github.com/P147x/remotelogs-api/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	InitRouter(r)
	defer db.GetDatabase().Close()
	r.Run(":3000")
}
