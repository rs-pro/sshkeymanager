package sshkeymanager

import (
	"log"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/pkg/errors"

	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
)

func (c *Client) GetKeys(user passwd.User) ([]authorized_keys.SSHKey, error) {
	raw, se, err := c.Execute("cat " + user.AuthorizedKeys())
	if err != nil {
		return []authorized_keys.SSHKey{}, errors.Wrap(err, raw+se+": failed to read "+user.AuthorizedKeys())
	}
	return authorized_keys.Parse(raw)
}

type KeyDoesNotExistError struct{}

func (e *KeyDoesNotExistError) Error() string {
	return "Key to delete not found in authorized_keys"
}

type KeyExistsError struct{}

func (e *KeyExistsError) Error() string {
	return "Key to add already present in authorized_keys"
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
		return &KeyDoesNotExistError{}
	}

	return c.WriteKeys(&user, newKeys)
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
		log.Println("WARN failed to read keys:", err)
	}

	for _, ck := range keys {
		if k.Key == ck.Key {
			return &KeyExistsError{}
		}
	}

	return c.WriteKeys(&user, append(keys, k))
}

func (c *Client) WriteKeys(user *passwd.User, keys []authorized_keys.SSHKey) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	err := c.CreateSSHDir(user)
	if err != nil {
		return err
	}

	keyFile := authorized_keys.Generate(keys)
	err = c.WriteFile(user.AuthorizedKeys(), keyFile)
	if err != nil {
		return err
	}

	err = c.ChownSSH(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) CreateSSHDir(user *passwd.User) error {
	so, se, err := c.Execute("mkdir -p " + shellescape.Quote(user.Home+"/.ssh"))
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}

func (c *Client) ChownHomedir(user *passwd.User) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	g := &group.Group{}
	g.GID = user.GID
	g = c.FindGroup(g)
	if g.Name == "" {
		return errors.New("failed to find group")
	}

	so, se, err := c.Execute("chown -R " + shellescape.Quote(user.Name+":"+g.Name) + " " + shellescape.Quote(user.Home))
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}

func (c *Client) ChownSSH(user *passwd.User) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	g := &group.Group{}
	g.GID = user.GID
	g = c.FindGroup(g)
	if g.Name == "" {
		return errors.New("failed to find group")
	}

	so, se, err := c.Execute("chown -R " + shellescape.Quote(user.Name+":"+g.Name) + " " + shellescape.Quote(user.Home) + "/.ssh")
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}
