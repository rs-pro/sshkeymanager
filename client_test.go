package sshkeymanager

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

func init() {
	cmd := exec.Command(
		"docker",
		"create",
		"--name=openssh-server",
		"--hostname=openssh-server",
		"-e PUID=1000",
		"-e PGID=1000",
		"-e TZ=Europe/Moscow",
		"-e PUBLIC_KEY=./test-data/id_rsa.pub",
		"-e SUDO_ACCESS=true",
		"-e PASSWORD_ACCESS=false",
		//"-p target=2222:published=12222",
		//"-v ./test-data/config:/config",
		"--restart=unless-stopped",
		"linuxserver/openssh-server",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestListUsers(t *testing.T) {
	log.Println("start")

	config := 
	s, err := Connect(config)
}
