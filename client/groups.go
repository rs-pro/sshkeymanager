package client

import (
	"github.com/rs-pro/sshkeymanager/api"
	"github.com/rs-pro/sshkeymanager/group"
)

func (c *Client) GetGroups() ([]group.Group, error) {
	r, err := c.Execute("get-groups", api.EmptyRequest{}, &api.GetGroupsResponse{})

	return r.(api.GetGroupsResponse).Groups, err
}

func (c *Client) ClearGroupCache() error {
	_, err := c.Execute("clear-group-cache", api.EmptyRequest{}, &api.EmptyResponse{})
	return err
}

func (c *Client) FindGroup(g *group.Group) (*group.Group, error) {
	r, err := c.Execute("find-group", g, &group.Group{})
	return r.(*group.Group), err
}

func (c *Client) AddGroup(g *group.Group) (*group.Group, error) {
	r, err := c.Execute("add-group", api.AddGroupRequest{Group: g}, api.AddGroupResponse{})
	if err != nil {
		return nil, err
	}
	return r.Group, r.Err
}

func (c *Client) DeleteGroup(g *group.Group) (*group.Group, error) {
	r, err := c.Execute("delete-group", api.DeleteGroupRequest{Group: g}, api.DeleteGroupResponse{})
	if err != nil {
		return nil, err
	}
	return r.Group, r.Err
}
