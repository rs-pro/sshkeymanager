package sshkeymanager

import (
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

var DefaultConfig *ssh.ClientConfig

func init() {
	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{},
	}

	//config.HostKeyCallback = ssh.FixedHostKey(hostKey)
	if os.Getenv("INSECURE_IGNORE_HOST_KEY") == "YES" {
		config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	for _, keyname := range keys {
		key, err := ioutil.ReadFile(keyname)
		if err == nil {
			//signer, err := ssh.ParsePrivateKey(key)
			signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(os.Getenv("KEY_PASS")))
			if err != nil {
				panic(err)
			}
			config.Auth = append(config.Auth, ssh.PublicKeys(signer))
		}
	}

	DefaultConfig = config
}

type Client struct {
	*ssh.Client
}

func makeConfig() *ssh.ClientConfig {
	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{},
	}

	//config.HostKeyCallback = ssh.FixedHostKey(hostKey)
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	for _, keyname := range keys {
		key, err := ioutil.ReadFile(keyname)
		if err == nil {
			//signer, err := ssh.ParsePrivateKey(key)
			signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(os.Getenv("KEY_PASS")))
			if err != nil {
				panic(err)
			}
			config.Auth = append(config.Auth, ssh.PublicKeys(signer))
		}
	}

	return config
}
