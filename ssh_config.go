package sshkeymanager

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	kh "golang.org/x/crypto/ssh/knownhosts"
)

func DefaultConfig() *ssh.ClientConfig {
	var keys []string

	if os.Getenv("SSH_KEY") != "" {
		keys = []string{os.Getenv("SSH_KEY")}
	} else {
		keys = []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}
	}

	return MakeConfig(keys)
}

func MakeConfig(keys []string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{},
	}

	if os.Getenv("SSH_HOST_KEY") != "" {
		key, err := ssh.ParsePublicKey([]byte(os.Getenv("SSH_HOST_KEY")))
		if err != nil {
			log.Fatal("failed to parse public key: ", os.Getenv("SSH_HOST_KEY"))
		}
		config.HostKeyCallback = ssh.FixedHostKey(key)
	} else if os.Getenv("INSECURE_IGNORE_HOST_KEY") == "YES" {
		if os.Getenv("INSECURE_IGNORE_HOST_KEY") != "YES" {
			log.Fatal("INSECURE_IGNORE_HOST_KEY: only possible value is YES in all caps")
		}
		config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	} else {
		hostKeyCallback, err := kh.New(os.Getenv("HOME") + "/.ssh/known_hosts")
		if err != nil {
			log.Fatal("could not create hostkeycallback function: ", err)
		}
		config.HostKeyCallback = hostKeyCallback
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
