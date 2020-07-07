package api

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

var r *gin.Engine

func GetRouter() *gin.Engine {
	if r == nil {
		r := gin.Default()
		r.Use(CheckApiKey())

		r.GET("/robots.txt", func(c *gin.Context) {
			c.Writer.WriteHeader(http.StatusOK)
			c.Writer.Write([]byte("User-agent: *\nDisallow: /"))
		}

		r.POST("/:host/:port/users", func(c *gin.Context) {
			host := c.Param("host")
			port := c.Param("port")
			client := sshkeymanager.NewClient(host, port, sshkeymanager.DefaultConfig)
			users, err := client.GetUsers()
			if err != nil {
				c.JSON(status, map[string]string{
					"error": err.Error() + ": " + message,
				})
				return
			}
			spew.Dump(client)
		})
	}

	return r
}
