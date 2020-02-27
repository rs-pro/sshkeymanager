package sshkeymanager

import (
	"log"
	"strings"
)

type User struct {
	Name  string
	UID   string
	Home  string
	Shell string
}

var users []User

func GetUsers(user string, host string, port string) []User {
	client := ConfigSSH(user, host, port)
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Unable to create session", err)
	}
	defer session.Close()
	raw, err := session.CombinedOutput("cat /etc/passwd")
	if err != nil {
		log.Fatal("Unable to run command", err)
	}
	rawToString := string(raw)

	strs := strings.Split(rawToString, "\n")

	for _, s := range strs {
		u := strings.Split(s, ":")
		if len(u) > 1 {
			var user User
			user.Name = u[0]
			user.UID = u[2]
			user.Home = u[5]
			user.Shell = u[6]
			users = append(users, user)
		}
	}
	return users
}
