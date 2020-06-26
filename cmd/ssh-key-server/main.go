package main

import (
	"log"

	"github.com/rs-pro/ssh-key-manager/pkg/api"
)

func main() {
	log.Fatal(api.GetRouter().Run())
}
