package passwd

import (
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type User struct {
	Name     string
	Password string
	UID      string
	GID      string
	Desc     string
	Home     string
	Shell    string
}

func Parse(raw string) ([]User, error) {
	strs := strings.Split(raw, "\n")

	users := make([]User, 0)
	for _, s := range strs {
		// skip empty lines
		if s == "" {
			continue
		}
		u := strings.Split(s, ":")
		if len(u) < 6 {
			log.Println("bad /etc/passwd entry:")
			spew.Dump(s)
			continue
		}
		var user User
		user.Name = u[0]
		user.Password = u[1]
		user.UID = u[2]
		user.GID = u[3]
		user.Desc = u[4]
		user.Home = u[5]
		user.Shell = u[6]
		users = append(users, user)
	}

	return users, nil
}

func (u *User) AuthorizedKeys() string {
	return u.Home + "/.ssh/authorized_keys"
}

func (u *User) Serialize() string {
	return strings.Join([]string{
		u.Name,
		u.Password,
		u.UID,
		u.GID,
		u.Desc,
		u.Home,
		u.Shell,
	}, ":")
}

func (u *User) UserAdd() string {
	command := []string{
		"useradd",
		"-m",
	}
	if u.UID != "" {
		command = append(command, "-u "+u.UID)
	}
	command = append(command, "-g "+u.GID)
	if u.Password != "" {
		command = append(command, "-p "+u.Password)
	} else {
		command = append(command, "-p x")
	}
	if u.Home != "" {
		command = append(command, "-d "+u.Home)
	}
	if u.Shell != "" {
		command = append(command, "-s "+u.Shell)
	}
	command = append(command, u.Name)
	return strings.Join(command, " ")
}
