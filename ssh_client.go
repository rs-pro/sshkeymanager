package sshkeymanager

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func NewClient(host, port, user string, config *ssh.ClientConfig) (*Client, error) {
	client := Client{host: host, port: port, SSHConfig: config}
	if user != "root" {
		client.useSudo = true
		client.user = user
		config.User = user
	}
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *Client) Prefix() string {
	if c.useSudo {
		return "sudo "
	} else {
		return ""
	}
}

func (c *Client) Connect() error {
	client, err := ssh.Dial("tcp", c.host+":"+c.port, c.SSHConfig)
	if err != nil {
		return err
	}
	c.SSHClient = client
	return nil
}

func (c *Client) Execute(command string) (string, string, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return "", "", errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()

	var so, se bytes.Buffer
	session.Stdout = &so
	session.Stderr = &se
	err = session.Run(c.Prefix() + command)
	if err != nil {
		log.Println("execute:", command, "result:", so.String(), se.String(), "error:", err)
		return strings.TrimSuffix(so.String(), "\n"), strings.TrimSuffix(se.String(), "\n"), err
	}

	if os.Getenv("DEBUG") == "YES" {
		log.Println("execute:", command)
		log.Println("result:")
		log.Println(so.String())
		log.Println(se.String())
	}

	return strings.TrimSuffix(so.String(), "\n"), strings.TrimSuffix(se.String(), "\n"), nil
}
