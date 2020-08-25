package sshkeymanager

import (
	"bytes"
	"log"
	"os"

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

func (c *Client) Execute(command string) (string, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	session.Run(c.Prefix() + command)

	if os.Getenv("DEBUG") == "YES" {
		log.Println("execute:", command, "result:", stdoutBuf.String(), "errors:", stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}
