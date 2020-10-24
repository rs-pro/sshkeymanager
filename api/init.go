package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager/config"
)

type EmptyRequest struct{}
type EmptyResponse struct{}

func init() {
	if config.Config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}
