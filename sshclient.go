package sshkeymanager

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

var (
	Home            string
	HostKeyCallback ssh.HostKeyCallback
)

func DefaultKeyPath() string {
	Home = os.Getenv("HOME")
	if len(Home) > 0 {
		return path.Join(Home, ".ssh/id_rsa")
	}
	return ""
}

func ConfigSSH(user string, host string, port string) *ssh.Client {
	key, err := ioutil.ReadFile(DefaultKeyPath())
	if err != nil {
		log.Fatal("Error reading private key", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("Error parsing private key", err)
	}

	HostKeyCallback, err = kh.New(path.Join(Home, ".ssh/known_hosts"))
	if err != nil {
		log.Fatal("Could not create hostkeycallback function: ", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: HostKeyCallback,
		Timeout:         10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("Could not dial to server", err)
	}

	return client
}
