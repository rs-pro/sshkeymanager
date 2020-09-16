package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

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
