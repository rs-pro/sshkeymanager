package passwd

import (
	"log"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/davecgh/go-spew/spew"

	"github.com/GehirnInc/crypt"
	_ "github.com/GehirnInc/crypt/sha256_crypt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	UID      string `json:"uid"`
	GID      string `json:"gid"`
	Desc     string `json:"desc"`
	Home     string `json:"home"`
	Shell    string `json:"shell"`
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

func (u *User) SSHDir() string {
	return u.Home + "/.ssh"
}

func (u *User) AuthorizedKeys() string {
	return u.SSHDir() + "/authorized_keys"
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

func (u *User) SetPassword(password string) {
	crypt := crypt.SHA256.New()
	ret, _ := crypt.Generate([]byte(password), []byte("$5$zz"))
	u.Password = ret
}

func (u *User) UserAdd() string {
	command := []string{
		"useradd",
		"-m",
	}
	if u.UID != "" {
		command = append(command, "-u "+shellescape.Quote(u.UID))
	}
	command = append(command, "-g "+shellescape.Quote(u.GID))
	if u.Password != "" {
		command = append(command, "-p "+shellescape.Quote(u.Password))
	} else {
		command = append(command, "-p x")
	}
	if u.Home != "" {
		command = append(command, "-d "+shellescape.Quote(u.Home))
	}
	if u.Shell != "" {
		command = append(command, "-s "+shellescape.Quote(u.Shell))
	}
	command = append(command, u.Name)
	return strings.Join(command, " ")
}

func (u *User) UserDelete(removeHome bool) string {
	command := []string{
		"userdel",
	}
	if removeHome {
		command = append(command, "-r")
	}
	command = append(command, u.Name)
	return strings.Join(command, " ")
}
