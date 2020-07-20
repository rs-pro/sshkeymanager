package sshkeymanager

import (
	"errors"
	"github.com/pkg/sftp"
	"path"
	"strconv"
	"strings"
)

type SSHKey struct {
	Num   int
	Key   string
	Email string
}

var allUsers []User

func (c *IClient) GetKeys(uid string) ([]SSHKey, error) {
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

func (c *IClient) DeleteKey(key string, uid string) error {
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

func (c *IClient) AddKey(key string, uid string) error {

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

func sync(keys []SSHKey, uid string, c *IClient) error {

	var homeDir string

	for _, h := range allUsers {
		if h.UID == uid {
			homeDir = h.Home
		}
	}

	client, err := sftp.NewClient(c.Cl)
	if err != nil {
		return err
	}
	defer client.Close()
	authorized_keys, err := client.Create(path.Join(homeDir, "/.ssh/authorized_keys"))
	if err != nil {
		return err
	}
	defer authorized_keys.Close()
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	err = authorized_keys.Chown(uidInt, uidInt)
	if err != nil {
		return err
	}
	err = authorized_keys.Chmod(0600)
	for _, k := range keys {
		if _, err := authorized_keys.Write([]byte(k.Key + " " + k.Email + "\n")); err != nil {
			return err
		}
	}

	return nil
}
