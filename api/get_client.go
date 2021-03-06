package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager"
)

const CONTEXT_KEY = "ssh-key-manager-client"

func SetClient(GetClient func(*gin.Context) *sshkeymanager.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := GetClient(c)
		if client == nil {
			c.Abort()
			return
		}
		c.Set(CONTEXT_KEY, client)
		c.Next()
	}
}

func GetClient(c *gin.Context) *sshkeymanager.Client {
	client, exists := c.Get(CONTEXT_KEY)
	if !exists {
		return nil
	}
	return client.(*sshkeymanager.Client)
}

// DefaultGetClient is designed to be overriden for custom API server settings
func DefaultGetClient(c *gin.Context) *sshkeymanager.Client {
	host := c.Query("host")
	port := c.Query("port")
	user := c.Query("user")

	client, err := sshkeymanager.NewClient(host, port, user, sshkeymanager.DefaultConfig())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return nil
	}

	return client
}
