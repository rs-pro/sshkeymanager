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

func defaultKeyPath() string {
	Home = os.Getenv("HOME")
	if len(Home) > 0 {
		return path.Join(Home, ".ssh/id_rsa")
	}
	return ""
}

func (c *Client) configSSH() error {
	key, err := ioutil.ReadFile(defaultKeyPath())
	if err != nil {
		return err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	HostKeyCallback, err = kh.New(path.Join(Home, ".ssh/known_hosts"))
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: c.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: HostKeyCallback,
		Timeout:         10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	c.Cl, err = ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}

	return nil
}
