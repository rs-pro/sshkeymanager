package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListGroups(c *gin.Context) {
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
}
