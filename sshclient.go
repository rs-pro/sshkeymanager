package sshkeymanager

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
	"io/ioutil"
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

func ConfigSSH(user string, host string, port string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(DefaultKeyPath())
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	HostKeyCallback, err = kh.New(path.Join(Home, ".ssh/known_hosts"))
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return client, nil
}
