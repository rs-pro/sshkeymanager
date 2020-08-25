package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

var r *gin.Engine

// GetClient is designed to be overriden for custom API server settings
func GetClient(c *gin.Context) *sshkeymanager.Client {
	host := c.Param("host")
	port := c.Param("port")
	user := c.Param("user")
	client, err := sshkeymanager.NewClient(host, port, user, sshkeymanager.DefaultConfig())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return nil
	}
	return client
}

// GetRouter
func GetRouter(GetClient func(*gin.Context) *sshkeymanager.Client) *gin.Engine {
	if r == nil {
		r := gin.Default()
		r.Use(CheckApiKey())

		r.GET("/robots.txt", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte("User-agent: *\nDisallow: /"))
		})

		r.POST("/:host/:port/users", func(c *gin.Context) {
			client := GetClient(c)
			if client == nil {
				return
			}
			users, err := client.GetUsers()
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, users)
		})
	}

	return r
}
