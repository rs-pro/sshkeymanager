package sshkeymanager

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/rs-pro/sshkeymanager/passwd"
	"github.com/rs-pro/sshkeymanager/testserver"
	"github.com/stretchr/testify/assert"
)

const host = "localhost"
const port = "2222"

func ExecCommand(user, command string) string {
	cmd := exec.Command("ssh", "-o StrictHostKeyChecking=no", "-itestdata/id_rsa", "-p "+port, user+"@"+host, command)
	var buf *bytes.Buffer
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("error in exec command", command, err)
	}
	return buf.String()
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	log.Println("test server: starting")
	server := testserver.Start()
	log.Println("test server: waiting to be up")

	WaitForSshServer(host, port)

	log.Println("test server: started")
	result := m.Run()
	log.Println("test server: stopping")
	server.Stop()
	os.Exit(result)
}

func WaitForSshServer(host, port string) {
	attempt := 0
	for {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
			// Give it one more second to finish starting up
			time.Sleep(1 * time.Second)
			return
		} else {
			log.Println("wait for ssh on", host, port, "attempt number:", attempt)
			attempt += 1
			if attempt > 30 {
				panic("failed to start ssh docker image")
			}
			time.Sleep(1 * time.Second)
		}

	}
}

func TestSudo(t *testing.T) {
	client, err := NewClient(host, port, "test", MakeConfig([]string{"./testdata/id_rsa"}))
	assert.NoError(t, err)
	assert.Equal(t, client.useSudo, true)

	users, err := client.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 25) // value from a clean image
}

func GetClient(t *testing.T) *Client {
	client, err := NewClient(host, port, "root", MakeConfig([]string{"./testdata/id_rsa"}))
	assert.NoError(t, err)
	return client
}

func TestListGroups(t *testing.T) {
	client := GetClient(t)

	users, err := client.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 25) // value from a clean image
}

func TestListUsers(t *testing.T) {
	client := GetClient(t)

	users, err := client.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 25) // value from a clean image
}

func TestAddUser(t *testing.T) {
	client := GetClient(t)

	u, err := client.AddUser(&passwd.User{
		Name:  "user",
		GID:   "1000",
		Home:  "/data/user",
		Shell: "/bin/bash",
	})
	assert.NoError(t, err)
	assert.Equal(t, "user", u.Name)
	assert.Equal(t, "1000", u.GID)
	assert.Equal(t, "/data/user", u.Home)
	assert.Equal(t, "/bin/bash", u.Shell)

	users, err := client.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 26)

}
