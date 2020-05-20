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

func GetUsers(user string, host string, port string) ([]User, error) {
	client, err := ConfigSSH(user, host, port)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	raw, err := session.CombinedOutput("cat /etc/passwd")
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
