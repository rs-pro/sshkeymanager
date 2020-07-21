package sshkeymanager

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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
	var counter int
	THIS:
	client, err := NewClient(host, port, clientCfg)
	if err != nil {
		time.Sleep(time.Second)
		counter += 1
		if counter >= 10 {
			log.Fatalf("Failed connect to %s:%s", host, port)
		}
		goto THIS
	}

	users, err := client.GetUsers()
	if err != nil {
		log.Println(err)
	}
	for _, u := range users {
		if u.UID == "1000" {
			fmt.Println(u)
		}
	}
}
