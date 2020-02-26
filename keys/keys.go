package keys

import (
	"github.com/ssh-key-manager/sshclient"
	"github.com/ssh-key-manager/users"
	"log"
	"strings"
)

type SSHKey struct {
	Num   int
	Key   string
	Email string
}

var sshKeys []SSHKey

func GetKeys(uid string, rootUser string, host string, port string) []SSHKey {

	usrs := users.GetUsers(rootUser, host, port)

	for _, u := range usrs {
		if u.UID == uid {
			client := sshclient.ConfigSSH(rootUser, host, port)
			defer client.Close()
			session, err := client.NewSession()
			if err != nil {
				log.Fatal("Unable to create session ", err)
			}
			defer session.Close()
			raw, err := session.CombinedOutput("cat " + u.Home + "/.ssh/authorized_keys")
			if err != nil {
				log.Fatal("Unable to run command ", err)
			}
			rawToString := string(raw)

			strs := strings.Split(rawToString, "\n")
			for i, s := range strs {
				k := strings.Fields(s)
				if len(k) > 1 {
					var sshKey SSHKey
					sshKey.Num = i
					sshKey.Key = k[1]
					sshKey.Email = k[2]
					sshKeys = append(sshKeys, sshKey)
				}
			}

		}
	}
	return sshKeys
}
