package sshkeymanager

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/rs-pro/sshkeymanager/testserver"

)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	log.Println("start test ssh server")
	server := testserver.Start()
	log.Println("started test ssh server")
	result := m.Run()
	server.Stop()
	os.Exit(result)
}

func TestListUsers(t *testing.T) {
	host := "localhost"
	port := "2222"
	clientCfg, err := makeTestConfig()
	if err != nil {
		log.Fatal(err)
	}
	client, err := NewClient(host, port, clientCfg)
	if err != nil {
		fmt.Println(err)
	}

	users, err := client.GetUsers()
	if err != nil {
		log.Println(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
}
