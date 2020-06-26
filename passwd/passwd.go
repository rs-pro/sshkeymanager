package passwd

import "strings"

type User struct {
	UID   string
	GID   string
	Name  string
	Home  string
	Shell string
}

func Parse(raw string) ([]User, error) {
	strs := strings.Split(raw, "\n")

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
