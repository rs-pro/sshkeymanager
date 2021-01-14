package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/config"
)

// Designed to be replaced with your custom implementation
// You can use this file as a starting point and add security, gin middleware etc.

func main() {
	listen := config.Config.Listen
	if listen == "" {
		listen = "127.0.0.1:12020"
	}

	r := gin.Default()
	r.Use(api.CheckApiKey())
	r.Use(api.SetClient(api.DefaultGetClient))
	if config.Config.Log {
		log.Println("enable full request log")
		r.Use(api.RequestLoggerMiddleware())
	}

	r = api.AddRoutes(r)

	log.Println("keyserver starting, listen on:", listen)
	log.Fatal(r.Run(listen))
}
