package sshkeymanager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
)

type SSHKey struct {
	Num   int
	Key   string
	Email string
}

func (c *Client) GetKeys(uid string) ([]SSHKey, error) {
	var (
		sshKeys []SSHKey
		user    User
	)
	var err error
	allUsers, err = c.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range allUsers {
		if u.UID == uid {
			user.Home = u.Home
		}
	}

	if err := c.NewSession(); err != nil {
		return nil, err
	}
	defer c.CloseSession()
	raw, err := c.Ses.CombinedOutput("cat " + user.Home + "/.ssh/authorized_keys")
	if err != nil {
		return nil, errors.New("Read error. Maybe \"~/.ssh/authorized_keys\" not exist. " + err.Error())
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

func (c *Client) DeleteKey(key string, uid string) error {
	var (
		newKeys []SSHKey
		newKey  SSHKey
	)

	fields := strings.Fields(key)
	newKey.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		newKey.Email = fields[2]
	}
	keys, err := c.GetKeys(uid)

	if err != nil {
		return err
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
	err = sync(newKeys, uid, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddKey(key string, uid string) error {

	var k SSHKey

	fields := strings.Fields(key)
	k.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		k.Email = fields[2]
	}
	keys, err := c.GetKeys(uid)

	if err != nil {
		return err
	}

	k.Num = len(keys) + 1
	for _, ck := range keys {
		if k.Key == ck.Key {
			return errors.New("Key exist!")

		}
	}

	keys = append(keys, k)
	err = sync(keys, uid, c)
	if err != nil {
		return err
	}
	return nil
}



func (c *Client) WriteFile(path string, content []byte) error {

}

func (c *Client) WriteKeys(keys []SSHKey) {
	keyFile := c.GenerateAuthorizedKeys(keys)
	c.WriteFile()
}

func (c *Client) sync(keys []SSHKey, uid string, c *IClient) {

	clientConfig, _ := auth.PrivateKey(c.User, path.Join(Home, ".ssh/id_rsa"), HostKeyCallback)

	client := scp.NewClient(c.Host+":"+c.Port, &clientConfig)

	err = client.Connect()
	if err != nil {
		return err
	}

	f, err := os.Open(tmpAuthorizedKeys.Name())
	if err != nil {
		return err
	}

	defer client.Close()

	var homeDir string

	for _, h := range allUsers {
		if h.UID == uid {
			homeDir = h.Home
		}
	}

	err = client.CopyFile(f, path.Join(homeDir, "/.ssh/authorized_keys"), "0600")

	if err != nil {
		return err
	}
	if err := os.Remove(tmpAuthorizedKeys.Name()); err != nil {
		return errors.New("Cannot delete file, not exist")
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
