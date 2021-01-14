package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetKeys(user passwd.User) ([]authorized_keys.SSHKey, error) {
	r, err := c.Execute("get-keys", &api.UserRequest{User: &user}, &api.UsersResponse{})
	if err != nil {
		return nil, err
	}

	kr := r.(*api.KeysResponse)

	return kr.Keys, kr.Err.Err()
}

func (c *Client) FindKey(user passwd.User, key authorized_keys.SSHKey) (*authorized_keys.SSHKey, error) {
	r, err := c.Execute("find-key", &api.KeyRequest{User: user, Key: &key}, &api.KeyResponse{})
	if err != nil {
		return nil, err
	}

	kr := r.(*api.KeyResponse)

	return kr.Key, kr.Err.Err()
}

func (c *Client) DeleteKey(user passwd.User, key authorized_keys.SSHKey) error {
	r, err := c.Execute("delete-key", &api.KeyRequest{User: user, Key: &key}, &api.BasicResponse{})
	if err != nil {
		return err
	}

	kr := r.(*api.BasicResponse)

	return kr.Err.Err()
}

func (c *Client) AddKey(user passwd.User, key authorized_keys.SSHKey) error {
	r, err := c.Execute("add-key", &api.KeyRequest{User: user, Key: &key}, &api.BasicResponse{})
	if err != nil {
		return err
	}

	kr := r.(*api.BasicResponse)

	return kr.Err.Err()
}

func (c *Client) WriteKeys(user passwd.User, keys []authorized_keys.SSHKey) error {
	r, err := c.Execute("write-keys", &api.KeysRequest{User: user, Keys: keys}, &api.BasicResponse{})
	if err != nil {
		return err
	}

	kr := r.(*api.BasicResponse)

	return kr.Err.Err()
}
