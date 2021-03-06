package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/P147x/remotelogs-api/pkg/utils"
	"net/http"
	"path/filepath"
	"os"
)

func SystemRoutes(r *gin.RouterGroup) {
	r.GET("/logs", logs)
	r.GET("/name", osname)
}

//GET /api/v1/system/logs
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

//GET /api/v1/system/name
func osname(c *gin.Context) {
	name, err := utils.UtsString()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": "200", "status": "ok", "name": name})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "500", "status": "error"})
	}
}