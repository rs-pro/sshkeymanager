package sshkeymanager

import (
	"github.com/rs-pro/sshkeymanager/passwd"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	host       string
	port       string
	SSHConfig  *ssh.ClientConfig
	SSHClient  *ssh.Client
	SSHSession *ssh.Session
	UsersCache *[]passwd.User
}

// GetPort returns client's ssh user
func (c *Client) GetUser() string {
	return c.SSHConfig.User
}

// GetPort returns client's ssh host
func (c *Client) GetHost() string {
	return c.host
}

// GetPort returns client's ssh port
func (c *Client) GetPort() string {
	return c.host
}
