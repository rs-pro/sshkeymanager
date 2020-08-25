package main

import (
	"log"
	"os"

	"github.com/rs-pro/sshkeymanager/api"
)

func main() {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":12020"
	}
	log.Fatal(api.GetRouter(api.GetClient).Run(listen))
}
