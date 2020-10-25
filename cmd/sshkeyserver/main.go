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
	log.Println("keyserver starting, listen on:", listen)

	r := gin.Default()
	r.Use(api.CheckApiKey())
	r.Use(api.SetClient(api.DefaultGetClient))
	r = api.AddRoutes(r)

	log.Fatal(r.Run(listen))
}
