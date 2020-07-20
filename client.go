package sshkeymanager

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	Cl *ssh.Client
	Ses *ssh.Session
	User string
	Host string
	Port string
}

func (c *Client) NewConnection()  error {
	var err error
	err = c.configSSH()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) NewSession()  error {
	var err error
	c.Ses, err = c.Cl.NewSession()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CloseConnection() error {
	err := c.Cl.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CloseSession() error {
	err := c.Ses.Close()
	if err != nil {
		return err
	}
	return nil
}