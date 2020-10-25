package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetUsers() ([]passwd.User, error) {
	r, err := c.Execute("get-users", &api.EmptyRequest{}, &api.GetUsersResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.Result().(*api.GetUsersResponse)

	return gr.Users, err
}

func (c *Client) ClearUserCache() error {
	_, err := c.Execute("clear-group-cache", &api.EmptyRequest{}, &api.EmptyResponse{})
	return err
}

func (c *Client) GetUserByUid(uid string) (*passwd.User, error) {
	return nil, nil
}

func (c *Client) GetUserByName(name string) (*passwd.User, error) {
	return nil, nil
}

func (c *Client) FindUser(user *passwd.User) (*passwd.User, error) {
	return nil, nil
}

func (c *Client) CreateHome(u *passwd.User) (*passwd.User, error) {
	return nil, nil
}

func (c *Client) AddUser(user *passwd.User, createHome bool) (*passwd.User, error) {
	return nil, nil
}

func (c *Client) DeleteUser(user *passwd.User, removeHome bool) (*passwd.User, error) {
	return nil, nil
}
