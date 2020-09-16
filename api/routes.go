package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

var r *gin.Engine

// GetRouter
func GetRouter(GetClient func(*gin.Context) *sshkeymanager.Client) *gin.Engine {
	if r == nil {
		r := gin.Default()
		r.Use(CheckApiKey())

		r.GET("/robots.txt", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte("User-agent: *\nDisallow: /"))
		})

		r.POST("/list-users", GetUsers)
		r.POST("/list-groups", GetGroups)
	}

	return r
}
