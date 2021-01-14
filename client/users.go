package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetUsers() ([]passwd.User, error) {
	r, err := c.Execute("get-users", &api.EmptyRequest{}, &api.UsersResponse{})
	//spew.Dump(r, err)
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
	return c.FindUser(&passwd.User{UID: uid})
}

func (c *Client) GetUserByName(name string) (*passwd.User, error) {
	return c.FindUser(&passwd.User{Name: name})
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
