package passwd

import (
	"log"
	"strings"
)

type User struct {
	UID   string
	GID   string
	Name  string
	Home  string
	Shell string
}

func Parse(raw string) ([]User, error) {
	strs := strings.Split(raw, "\n")

	users := make([]User, 0)
	for _, s := range strs {
		u := strings.Split(s, ":")
		if len(s) < 6 {
			log.Println("bad /etc/passwd entry:", s)
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

func (u User) AuthorizedKeys() string {
	return u.Home + "/.ssh/authorized_keys"
}
