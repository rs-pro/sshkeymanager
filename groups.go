package sshkeymanager

import (
	"log"

	"github.com/pkg/errors"
	"github.com/rs-pro/sshkeymanager/group"
)

func (c *Client) GetGroups() ([]group.Group, error) {
	if c == nil {
		return nil, errors.New("client not initialized")
	}
	if c.GroupsCache == nil {
		raw, err := c.Execute("cat /etc/group")
		if err != nil {
			return nil, err
		}
		groups, err := group.Parse(raw)
		if err != nil {
			return nil, err
		}
		c.GroupsCache = &groups
	}

	return *c.GroupsCache, nil
}

func (c *Client) ClearGroupCache() {
	c.GroupsCache = nil
}

func (c *Client) FindGroup(group *group.Group) *group.Group {
	groups, err := c.GetGroups()
	if err != nil {
		log.Println("failed to get groups", err)
		return nil
	}
	for _, g := range groups {
		if g.GID == group.GID || g.Name == group.Name {
			return &g
		}
	}
	return nil
}

func (c *Client) AddGroup(group *group.Group) (*group.Group, error) {
	if group.Name == "" {
		return nil, errors.New("'group name cannot be empty'")
	}
	g := c.FindGroup(group)
	if g != nil {
		return g, nil
	}

	_, err := c.Execute(group.GroupAdd())
	if err != nil {
		return nil, err
	}

	c.ClearGroupCache()

	g = c.FindGroup(group)
	if g == nil {
		return nil, errors.New("failed to add group")
	}
	return g, nil
}
