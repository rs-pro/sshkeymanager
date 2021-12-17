package sshkeymanager

import (
	"log"

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

func (c *Client) FindKey(user passwd.User, key authorized_keys.SSHKey) (*authorized_keys.SSHKey, error) {
	keys, err := c.GetKeys(user)
	if err != nil {
		return nil, err
	}

	for _, k := range keys {
		if k.Key == key.Key || k.Email == key.Email {
			return &k, nil
		}
	}

	// Key not found
	return nil, nil
}

func (c *Client) DeleteKey(user passwd.User, key authorized_keys.SSHKey) error {
	var (
		newKeys []authorized_keys.SSHKey
	)

	keys, err := c.GetKeys(user)
	if err != nil {
		return err
	}

	var keyExist bool
	for _, k := range keys {
		if k.Key != key.Key {
			newKeys = append(newKeys, k)
		} else {
			keyExist = true
		}
	}

	if !keyExist {
		return &KeyDoesNotExistError{}
	}

	return c.WriteKeys(user, newKeys)
}

func (c *Client) AddKey(user passwd.User, key authorized_keys.SSHKey) error {
	keys, err := c.GetKeys(user)

	if err != nil {
		log.Println("WARN failed to read keys:", err)
	}

	for _, ck := range keys {
		if key.Key == ck.Key {
			return &KeyExistsError{}
		}
	}

	return c.WriteKeys(user, append(keys, key))
}

func (c *Client) WriteKeys(user passwd.User, keys []authorized_keys.SSHKey) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	err := c.createSSHDir(user)
	if err != nil {
		return err
	}

	keyFile := authorized_keys.Generate(keys)
	err = c.WriteFile(user.AuthorizedKeys(), keyFile)
	if err != nil {
		return err
	}

	err = c.chownSSH(user)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createSSHDir(user passwd.User) error {
	so, se, err := c.Execute("mkdir -p " + shellescape.Quote(user.Home+"/.ssh"))
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}

func (c *Client) chownHomedir(user passwd.User) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	g, err := c.FindGroup(group.Group{GID: user.GID})
	if err != nil {
		return err
	}
	if g == nil {
		return errors.New("group not found")
	}
	if g.Name == "" {
		return errors.New("failed to find group")
	}

	so, se, err := c.Execute("chown -R " + shellescape.Quote(user.Name+":"+g.Name) + " " + shellescape.Quote(user.Home))
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}

func (c *Client) chownSSH(user passwd.User) error {
	if user.Name == "" {
		return errors.New("no user name")
	}

	g, err := c.FindGroup(group.Group{GID: user.GID})
	if err != nil {
		return err
	}
	if g == nil {
		return errors.New("group not found")
	}
	if g.Name == "" {
		return errors.New("failed to find group")
	}

	so, se, err := c.Execute("chown -R " + shellescape.Quote(user.Name+":"+g.Name) + " " + shellescape.Quote(user.Home) + "/.ssh")
	if err != nil {
		return errors.Wrap(err, so+se)
	}
	return nil
}
