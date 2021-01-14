package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetUsers() ([]passwd.User, error) {
	r, err := c.Execute("get-users", &api.EmptyRequest{}, &api.UsersResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.UsersResponse)

	return gr.Users, gr.Err.Err()
}

func (c *Client) ClearUserCache() error {
	// NO OP - not needed in client-server mode
	// method is present for interface compatibility
	return nil
}

func (c *Client) GetUserByUid(uid string) (*passwd.User, error) {
	return c.FindUser(passwd.User{UID: uid})
}

func (c *Client) GetUserByName(name string) (*passwd.User, error) {
	return c.FindUser(passwd.User{Name: name})
}

func (c *Client) FindUser(user passwd.User) (*passwd.User, error) {
	r, err := c.Execute("find-user", &api.UserRequest{User: &user}, &api.UsersResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.UserResponse)

	return gr.User, gr.Err.Err()
}

func (c *Client) CreateHome(user passwd.User) (*passwd.User, error) {
	r, err := c.Execute("create-home", &api.UserRequest{User: &user}, &api.UserResponse{})
	if err != nil {
		return nil, err
	}

	ur := r.(*api.UserResponse)
	return ur.User, ur.Err.Err()
}

func (c *Client) AddUser(user passwd.User, createHome bool) (*passwd.User, error) {
	r, err := c.Execute("add-user", &api.UserRequest{User: &user, CreateHome: &createHome}, &api.UserResponse{})
	if err != nil {
		return nil, err
	}

	ur := r.(*api.UserResponse)
	return ur.User, ur.Err.Err()
}

func (c *Client) DeleteUser(user passwd.User, removeHome bool) (*passwd.User, error) {
	r, err := c.Execute("delete-user", &api.UserRequest{User: &user, RemoveHome: &removeHome}, &api.UserResponse{})
	if err != nil {
		return nil, err
	}

	ur := r.(*api.UserResponse)
	return ur.User, ur.Err.Err()
}
