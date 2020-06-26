package sshkeymanager

import (
	"log"

	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetUsers() ([]passwd.User, error) {
	if c.UsersCache == nil {
		raw, err := c.Execute("cat /etc/passwd")
		if err != nil {
			return nil, err
		}
		users, err := passwd.Parse(raw)
		if err != nil {
			return nil, err
		}
		c.UsersCache = &users
	}

	return *c.UsersCache, nil
}

func (c *Client) ClearUserCache() {
	c.UsersCache = nil
}

func (c *Client) GetUserByUid(uid string) *passwd.User {
	users, err := c.GetUsers()
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, u := range users {
		if u.UID == uid {
			return &u
		}
	}
	return nil
}

func (c *Client) GetUserByName(name string) *passwd.User {
	users, err := c.GetUsers()
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, u := range users {
		if u.Name == name {
			return &u
		}
	}
	return nil
}
