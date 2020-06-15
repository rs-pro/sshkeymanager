package sshkeymanager

import (
	"strings"
)

type User struct {
	Name  string
	UID   string
	Home  string
	Shell string
}

var users []User

func (c *IClient) GetUsers() ([]User, error) {
	if err := c.NewSession(); err != nil {
		return nil, err
	}
	defer c.CloseSession()
	raw, err := c.Ses.CombinedOutput("cat /etc/passwd")
	if err != nil {
		return nil, err
	}
	rawToString := string(raw)

	strs := strings.Split(rawToString, "\n")

	for _, s := range strs {
		u := strings.Split(s, ":")
		if len(s) == 0 {
			continue
		}
		var user User
		user.Name = u[0]
		user.UID = u[2]
		user.Home = u[5]
		user.Shell = u[6]
		users = append(users, user)
	}

	return users, nil
}
