package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

var r *gin.Engine

func SetClient(GetClient func(*gin.Context) *sshkeymanager.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := GetClient(c)
		if client == nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

// GetRouter
func GetRouter(GetClient func(*gin.Context) *sshkeymanager.Client) *gin.Engine {
	if r == nil {
		r := gin.Default()
		r.Use(CheckApiKey())
		r.Use(SetClient(GetClient))

		r.GET("/robots.txt", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte("User-agent: *\nDisallow: /"))
		})

		r.POST("/get-groups", GetGroups)
		r.POST("/add-group", AddGroup)
		r.POST("/delete-group", DeleteGroup)

		r.POST("/get-users", GetUsers)
		r.POST("/add-user", AddUser)
	}

	return r
}
