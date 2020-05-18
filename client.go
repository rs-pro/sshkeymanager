package sshkeymanager

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	Client  *ssh.Client
	Session *ssh.Session
}

func (c *Client) StartSession() error {
	if cl.Session != nil {
		return nil
	}

	session, err := client.NewSession()
	cl.Session = session
	if err != nil {
		return errors.Wrap(err, "Unable to create session")
	}
	return nil
}

func (c *Client) CloseSession() error {
	if c.Session == nil {
		return errors.New("no active session")
	}

	err := c.Session.Close()
	if err != nil {
		return errors.Wrap(err, "Unable to close session")
	}

	c.Session = nil
	return nil
}

func NewWithClient(c *ssh.Client) (*Client, error) {
	cl := &Client{
		Client * ssh.Client,
	}
	var err error
	cl.Session, err = cl.StartSession()
	return cl, err
}

func NewDefault() (*Client, error) {
	return NewWithClient(ConfigSSH())
}

func (c *Client) Close() {
	c.EndSession()
	c.SSHClient.Close()
}
