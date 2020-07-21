package sshkeymanager

import (
	"fmt"
	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/passwd"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func (c *Client) GetKeys(user passwd.User) ([]authorized_keys.SSHKey, error) {
	raw, err := c.Execute("cat " + user.AuthorizedKeys())
	if err != nil {
		return nil, errors.Wrap(err, "failed to read "+user.AuthorizedKeys())
	}
	return authorized_keys.Parse(raw)
}

func (c *Client) DeleteKey(user passwd.User, key string) error {
	var (
		newKeys []authorized_keys.SSHKey
		newKey  authorized_keys.SSHKey
	)

	fields := strings.Fields(key)
	newKey.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		newKey.Email = fields[2]
	}

	keys, err := c.GetKeys(user)
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
		return errors.New("Key to delete not found in authorized_keys")
	}

	return c.WriteKeys(user, newKeys)
}

func (c *Client) AddKey(user passwd.User, key string) error {
	var k authorized_keys.SSHKey

	fields := strings.Fields(key)
	k.Key = fields[0] + " " + fields[1]
	if len(fields) > 2 {
		k.Email = fields[2]
	}
	keys, err := c.GetKeys(user)

	if err != nil {
		return err
	}

	for _, ck := range keys {
		if k.Key == ck.Key {
			return errors.New("Key already present in authorized_keys")
		}
	}

	return c.WriteKeys(user, append(keys, k))
}

func (c *Client) StartSCP(session *ssh.Session) error {
	err := session.Run("/usr/bin/scp -tr /")
	if err != nil {
		return errors.Wrap(err, "Failed to run")
	}
	return nil
}

func (c *Client) WriteFile(path string, content []byte) error {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		fmt.Fprintln(w, "D0700", 0, filepath.Dir(path))
		fmt.Fprintln(w, "C0600", len(content), filepath.Base(path))
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")

		r, _ := session.StdoutPipe()
		data, err := ioutil.ReadAll(r)
		log.Println("scp response", data, err)
		defer w.Close()
		session.Close()
	}()

	c.StartSCP(session)
	return nil
}

func (c *Client) WriteKeys(user passwd.User, keys []authorized_keys.SSHKey) error {
	keyFile := authorized_keys.Generate(keys)
	return c.WriteFile(user.AuthorizedKeys(), keyFile)
}
