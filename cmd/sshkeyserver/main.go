package main

import (
	"log"

	"github.com/rs-pro/sshkeymanager/api"
)

func main() {
	log.Fatal(api.GetRouter().Run())
}
