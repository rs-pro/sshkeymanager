package sshkeymanager

import (
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
	client := sshkeymanager.NewClient(host, port, sshkeymanager.DefaultConfig)
	users, err := client.GetUsers()
	if err != nil {
		log.Println(err)
	}
}
