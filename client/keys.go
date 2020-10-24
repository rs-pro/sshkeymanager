package client

import (
	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetKeys(user passwd.User) ([]authorized_keys.SSHKey, error) {
	return nil, nil
}

func (c *Client) DeleteKey(user passwd.User, key string) error {
	return nil
}

func (c *Client) AddKey(user passwd.User, key string) error {
	return nil
}

func (c *Client) WriteKeys(user *passwd.User, keys []authorized_keys.SSHKey) error {
	return nil
}

func (c *Client) CreateSSHDir(user *passwd.User) error {
	return nil
}

func (c *Client) ChownHomedir(user *passwd.User) error {
	return nil
}

func (c *Client) ChownSSH(user *passwd.User) error {
	return nil
}
