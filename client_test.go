package sshkeymanager

import (
	"github.com/stretchr/testify/assert"
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
	CONNECT:
	client, err := NewClient(host, port, clientCfg)
	if err != nil {
		time.Sleep(time.Second)
		counter += 1
		if counter >= 10 {
			log.Fatalf("Failed connect to %s:%s", host, port)
		}
		goto CONNECT
	}

	users, err := client.GetUsers()
	if err != nil {
		log.Println(err)
	}
	assert.NotEqual(t, len(users), 0)
		u := users[len(users)-1]
		assert.Equal(t, "1000", u.UID, "UID should be 1000")
		assert.Equal(t, "1000", u.GID, "UID should be 1000")
		assert.Equal(t, "test", u.Name, "Name should be test")
		assert.Equal(t, "/config", u.Home, "Home should be /config")
		assert.Equal(t, "/bin/bash", u.Shell, "Shell should be /bin/bash")
}


