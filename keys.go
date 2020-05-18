package sshkeymanager

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
)

type SSHKey struct {
	Num   int
	Key   string
	Email string
}

var allUsers []User

func (c *Client) GetKeys(uid string, rootUser string, host string, port string) ([]SSHKey, error) {
	var (
		sshKeys []SSHKey
		user    User
	)
	allUsers = c.GetUsers(rootUser, host, port)
	for _, u := range allUsers {
		if u.UID == uid {
			user.Home = u.Home
		}
	}

	raw, err := c.Session.CombinedOutput("cat " + user.Home + "/.ssh/authorized_keys")
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

func (c *Client) DeleteKey(key string, uid string, rootUser string, host string, port string) error {
	var (
		newKeys []SSHKey
		newKey  SSHKey
	)

	fields := strings.Fields(key)
	newKey.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		newKey.Email = fields[2]
	}
	keys, err := c.GetKeys(uid, rootUser, host, port)

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
	return c.sync(newKeys, uid, rootUser, host, port)
}

func (c *Client) AddKey(key string, uid string, rootUser string, host string, port string) error {

	var k SSHKey

	fields := strings.Fields(key)
	k.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		k.Email = fields[2]
	}
	keys, err := c.GetKeys(uid, rootUser, host, port)

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
	return c.sync(keys, uid, rootUser, host, port)
}

func (c *Client) sync(keys []SSHKey, uid string, rootUser string, host string, port string) error {

	f, err := os.Create("authorized_keys")
	if err != nil {
		log.Fatal("Cannot create file ", err)
		f.Close()
		return err
	}

	for _, k := range keys {
		fmt.Fprintln(f, k.Key+" "+k.Email)
	}
	err = f.Close()
	if err != nil {
		log.Fatal("Cannot write to file", err)
		return err
	}

	clientConfig, _ := auth.PrivateKey(rootUser, path.Join(Home, ".ssh/id_rsa"), HostKeyCallback)

	client := scp.NewClient(host+":"+port, &clientConfig)

	err = client.Connect()
	if err != nil {
		log.Fatal("Couldn't establish a connection to the remote server ", err)
		return err
	}

	f, err := os.Open("authorized_keys")
	if err != nil {
		log.Fatal("Couldn't open file ", err)
		return err
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
		return err
	}
	if err := os.Remove("authorized_keys"); err != nil {
		log.Println("Cannot delete file, not exist")
		return err
	}
}
