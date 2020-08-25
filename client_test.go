package sshkeymanager

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/rs-pro/sshkeymanager/testserver"
	"github.com/stretchr/testify/assert"
)

const host = "localhost"
const port = "2222"

func ExecCommand(command string) {
	cmd := exec.Command("ssh", "-o StrictHostKeyChecking=no", "-itestdata/id_rsa", "-p "+port, "test@"+host, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("error in exec command", command, err)
		// It's ok to have bad exit code - we are restarting SSH server
		//panic(err)
	}
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	log.Println("test server: starting")
	server := testserver.Start()
	defer server.Stop()
	log.Println("test server: waiting to be up")
	WaitForSshServer(host, port)

	log.Println("test server: enabling root login")
	command := `sudo mkdir /root/.ssh`
	ExecCommand(command)

	command = `sudo cp -R /config/.ssh/* /root/.ssh/`
	ExecCommand(command)

	command = `sudo chmod -R 0700 /root/.ssh/`
	ExecCommand(command)

	command = `grep -qxF "PermitRootLogin without-password" /etc/ssh/sshd_config || echo "PermitRootLogin without-password" | sudo tee -a /etc/ssh/sshd_config`
	ExecCommand(command)

	command = `sudo killall sshd`
	ExecCommand(command)

	log.Println("test server: restarted")
	time.Sleep(1 * time.Second)
	WaitForSshServer(host, port)

	log.Println("test server: started")
	result := m.Run()
	log.Println("test server: stopping")
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

func TestListUsers(t *testing.T) {
	client, err := NewClient(host, port, MakeConfig([]string{"./testdata/id_rsa"}))
	assert.NoError(t, err)

	users, err := client.GetUsers()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}
