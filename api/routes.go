package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddRoutes
func AddRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/robots.txt", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("User-agent: *\nDisallow: /"))
	})

	// Do NOT use GET - less secure

	r.POST("/get-groups", GetGroups)
	r.POST("/add-group", AddGroup)
	r.POST("/delete-group", DeleteGroup)

	r.POST("/get-users", GetUsers)
	r.POST("/add-user", AddUser)

	return r
}
