package sshclient

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
)

var home string

func defaultKeyPath() string {
	home = os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

func ConfigSSH(user string, host string, port string) *ssh.Client {
	key, err := ioutil.ReadFile(defaultKeyPath())
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}

	hostKeyCallback, err := kh.New(path.Join(home, ".ssh/known_hosts"))
	if err != nil {
		log.Fatal("could not create hostkeycallback function: ", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal(err)
	}

	return client
}