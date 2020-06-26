package sshkeymanager

import "github.com/rs-pro/sshkeymanager/passwd"

type User struct {
	UID   string
	GID   string
	Name  string
	Home  string
	Shell string
}

func (c *Client) GetUsers() ([]User, error) {
	raw, err := c.Execute("cat /etc/passwd")
	if err != nil {
		return nil, err
	}
	return passwd.Parse(raw)
}
