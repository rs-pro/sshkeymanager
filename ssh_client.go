package sshkeymanager

import (
	"bytes"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func (c *Client) Connect() error {
	client, err := ssh.Dial("tcp", c.host+":"+c.port, c.SSHConfig)
	if err != nil {
		return err
	}
	c.SSHClient = client
	return nil
}

func NewClient(host, port string, config *ssh.ClientConfig) (*Client, error) {
	client := Client{host: host, port: port, SSHConfig: config}
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Execute(command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)

	return stdoutBuf.String(), nil
}
