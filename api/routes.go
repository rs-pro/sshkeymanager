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
	r.POST("/find-group", FindGroup)
	r.POST("/add-group", AddGroup)
	r.POST("/delete-group", DeleteGroup)

	r.POST("/get-users", GetUsers)
	r.POST("/find-user", FindUser)
	r.POST("/create-home", CreateHome)
	r.POST("/add-user", AddUser)
	r.POST("/delete-user", DeleteUser)

	r.POST("/get-keys", GetKeys)
	r.POST("/delete-key", DeleteKey)
	r.POST("/add-key", AddKey)
	r.POST("/write-keys", WriteKeys)

	return r
}
