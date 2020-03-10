package sshkeymanager

import (
	"errors"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"log"
	"os"
	"path"
	"strings"
)

type SSHKey struct {
	Num   int
	Key   string
	Email string
}

var allUsers []User

func GetKeys(uid string, rootUser string, host string, port string) ([]SSHKey, error) {
	var (
		sshKeys []SSHKey
		user    User
	)
	allUsers = GetUsers(rootUser, host, port)
	for _, u := range allUsers {
		if u.UID == uid {
			user.Home = u.Home
		}
	}
	client := ConfigSSH(rootUser, host, port)
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Unable to create session ", err)
	}
	defer session.Close()
	raw, err := session.CombinedOutput("cat " + user.Home + "/.ssh/authorized_keys")
	if err != nil {
		return nil, errors.New("Read error. Maybe \"~/.ssh/authorized_keys\" not exist.")
	}
	rawToString := string(raw)

	keysStrings := strings.Split(rawToString, "\n")
	for i, s := range keysStrings {
		if len(s) == 0 {
			continue
		}
		k := strings.Fields(s)
		var sshKey SSHKey
		sshKey.Num = i + 1
		sshKey.Key = k[0] + " " + k[1]
		if len(k) > 2 {
			sshKey.Email = k[2]
		}
		sshKeys = append(sshKeys, sshKey)
	}

	return sshKeys, nil
}

func DeleteKey(key string, uid string, rootUser string, host string, port string) error{
	var (
		newKeys []SSHKey
		newKey  SSHKey
	)

	fields := strings.Fields(key)
	newKey.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		newKey.Email = fields[2]
	}
	keys, err := GetKeys(uid, rootUser, host, port)

	if err != nil {
		log.Println(err)
	}

	var keyExist bool
	for _, k := range keys {
		if k.Key != newKey.Key {
			newKeys = append(newKeys, k)
		} else {
			keyExist = true
		}
	}
	if !keyExist {
		return errors.New("Key is not exist")
	}
	sync(newKeys, uid, rootUser, host, port)
	return nil
}

func AddKey(key string, uid string, rootUser string, host string, port string) error{

	var k SSHKey

	fields := strings.Fields(key)
	k.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		k.Email = fields[2]
	}
	keys, err := GetKeys(uid, rootUser, host, port)

	if err != nil {
		log.Println(err)
	}

	k.Num = len(keys) + 1
	for _, ck := range keys {
		if k.Key == ck.Key {
			return errors.New("Key exist!")

		}
	}

	keys = append(keys, k)
	sync(keys, uid, rootUser, host, port)
	return nil
}

func sync(keys []SSHKey, uid string, rootUser string, host string, port string) {

	f, err := os.Create("authorized_keys")
	if err != nil {
		log.Fatal("Cannot create file ", err)
		f.Close()
		return
	}

	for _, k := range keys {
		fmt.Fprintln(f, k.Key+" "+k.Email)
	}
	err = f.Close()
	if err != nil {
		log.Fatal("Cannot write to file", err)
		return
	}

	clientConfig, _ := auth.PrivateKey(rootUser, path.Join(Home, ".ssh/id_rsa"), HostKeyCallback)

	client := scp.NewClient(host+":"+port, &clientConfig)

	err = client.Connect()
	if err != nil {
		log.Fatal("Couldn't establish a connection to the remote server ", err)
		return
	}

	f, errFile := os.Open("authorized_keys")
	if errFile != nil {
		log.Fatal("Couldn't open file ", errFile)
	}

	defer client.Close()

	defer f.Close()

	var homeDir string

	for _, h := range allUsers {
		if h.UID == uid {
			homeDir = h.Home
		}
	}

	err = client.CopyFile(f, path.Join(homeDir, "/.ssh/authorized_keys"), "0600")

	if err != nil {
		log.Fatal("Error while copying file ", err)
	}
	if err := os.Remove("authorized_keys"); err != nil {
		log.Println("Cannot delete file, not exist")
	}
}
