package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/group"
)

func (c *Client) GetGroups() ([]group.Group, error) {
	r, err := c.Execute("get-groups", &api.EmptyRequest{}, &api.GroupsResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.GroupsResponse)
	return gr.Groups, gr.Err.Err()
}

func (c *Client) ClearGroupCache() error {
	// NO OP - not needed in client-server mode
	// method is present for interface compatibility
	return nil
}

func (c *Client) FindGroup(g group.Group) (*group.Group, error) {
	r, err := c.Execute("find-group", &api.GroupRequest{Group: &g}, &api.GroupResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.GroupResponse)
	return gr.Group, gr.Err.Err()
}

func (c *Client) AddGroup(g group.Group) (*group.Group, error) {
	r, err := c.Execute("add-group", &api.GroupRequest{Group: &g}, &api.GroupResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.GroupResponse)
	return gr.Group, gr.Err.Err()
}

func (c *Client) DeleteGroup(g group.Group) (*group.Group, error) {
	r, err := c.Execute("delete-group", &api.GroupRequest{Group: &g}, &api.GroupResponse{})
	if err != nil {
		return nil, err
	}

	gr := r.(*api.GroupResponse)
	return gr.Group, gr.Err.Err()
}
