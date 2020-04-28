package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"os"
)

func SystemRoutes(r *gin.RouterGroup) {
	r.GET("/logs", logs)
}

//GET /api/v1/sys/logs
func logs(c *gin.Context) {
	var logs []string
	err := filepath.Walk("/var/log",
        func(path string, info os.FileInfo, err error) error {
        	if err != nil {
                return err
            }
			logs = append(logs, path)
            return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "status": "error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "200", "status": "ok", "logs": logs})
	}
}