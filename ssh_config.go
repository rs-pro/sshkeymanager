package sshkeymanager

import (
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

var DefaultConfig *ssh.ClientConfig

func init() {
	DefaultConfig = makeConfig()
}

func makeConfig() *ssh.ClientConfig {
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
			var signer ssh.Signer
			if os.Getenv("KEY_PASS") != "" {
				signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(os.Getenv("KEY_PASS")))
			} else {
				signer, err = ssh.ParsePrivateKey(key)
			}
			if err != nil {
				panic(err)
			}
			config.Auth = append(config.Auth, ssh.PublicKeys(signer))
		}
	}

	return config
}
