package main

import (
	"log"

	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/config"
)

func main() {
	listen := config.Config.Listen
	if listen == "" {
		listen = "127.0.0.1:12020"
	}
	log.Fatal(api.GetRouter(api.DefaultGetClient).Run(listen))
}
