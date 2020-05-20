package sshkeymanager

import (
	"errors"
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
	allUsers, err := c.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, u := range allUsers {
		if u.UID == uid {
			user.Home = u.Home
		}
	}
	//client, err := ConfigSSH(rootUser, host, port)
	//if err != nil {
	//	return nil, err
	//}
	//defer client.Close()
	//session, err := client.NewSession()
	//if err != nil {
	//	return nil, err
	//}
	//defer session.Close()
	raw, err := c.Ses.CombinedOutput("cat " + user.Home + "/.ssh/authorized_keys")
	//raw, err := session.CombinedOutput("cat " + user.Home + "/.ssh/authorized_keys")
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

func (c *IClient) DeleteKey(key string, uid string) error{
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
	//err = sync(newKeys, uid, c)
	if err != nil {
		return err
	}
	return nil
}

func (c *IClient) AddKey(key string, uid string) error{

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
	//err = sync(keys, uid, c)
	if err != nil {
		return err
	}
	return nil
}
// Подумать о передаче сессии
//func sync(keys []SSHKey, uid string, c *IClient) error {
//
//	tmpAuthorizedKeys, err := ioutil.TempFile("", "authorizedKeys")
//	if err != nil {
//		return err
//	}
//
//	for _, k := range keys {
//		fmt.Fprintln(tmpAuthorizedKeys, k.Key+" "+k.Email)
//	}
//	err = tmpAuthorizedKeys.Close()
//	if err != nil {
//		return err
//	}
//
//	clientConfig, _ := auth.PrivateKey(rootUser, path.Join(Home, ".ssh/id_rsa"), HostKeyCallback)
//
//	client := scp.NewClient(host+":"+port, &clientConfig)
//
//	err = client.Connect()
//	if err != nil {
//		return err
//	}
//
//	f, err := os.Open(tmpAuthorizedKeys.Name())
//	if err != nil {
//		return err
//	}
//
//	defer client.Close()
//
//	var homeDir string
//
//	for _, h := range allUsers {
//		if h.UID == uid {
//			homeDir = h.Home
//		}
//	}
//
//	err = client.CopyFile(f, path.Join(homeDir, "/.ssh/authorized_keys"), "0600")
//
//	if err != nil {
//		return err
//	}
//	if err := os.Remove(tmpAuthorizedKeys.Name()); err != nil {
//		return errors.New("Cannot delete file, not exist")
//	}
//	err = f.Close()
//	if err != nil {
//		return err
//	}
//	return nil
//}
