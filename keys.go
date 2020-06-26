package sshkeymanager

import (
	"errors"
	"strings"
)

func (c *Client) GetKeys(user User) ([]SSHKey, error) {
	raw, err := c.Execute("cat " + user.AuthorizedKeys())
	if err != nil {
		return nil, errors.Wrap(err, "failed to read "+user.AuthorizedKeys())
	}
	return authorized_keys.Parse(raw)
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

func (c *Client) StartSCP() error {
	if err := session.Run("/usr/bin/scp -tr /"); err != nil {
		panic("Failed to run: " + err.Error())
	}
}

func (c *Client) WriteFile(path string, content []byte) error {
	session, err := client.NewSession()
	if err != nil {
		return "", errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()
	go func() {
		w, _ := session.StdinPipe()
		defer w.Close()
		fmt.Fprintln(w, "D0700", 0, filepath.Dir(path))
		fmt.Fprintln(w, "C0600", len(content), filepath.Base(path))
		fmt.Fprint(w, content)
		fmt.Fprint(w, "\x00")
	}()
}

func (c *Client) WriteKeys(user User, keys []SSHKey) {
	keyFile := authorized_keys.Generate(keys)
	c.WriteFile(user.AuthorizedKeys(), keyFile)
}
