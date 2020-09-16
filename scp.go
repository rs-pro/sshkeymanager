package sshkeymanager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func (c *Client) StartSCP(session *ssh.Session, path string) error {
	err := session.Run(c.Prefix() + "/usr/bin/scp -tr " + path)
	if err != nil {
		return errors.Wrap(err, "Failed to run")
	}
	return nil
}

func (c *Client) WriteFile(path string, content []byte) error {
	if os.Getenv("DEBUG") == "YES" {
		log.Println("write file:", path, "content:", string(content))
	}

	session, err := c.SSHClient.NewSession()
	if err != nil {
		return errors.Wrap(err, "ssh NewSession")
	}
	defer session.Close()
	go func() {
		r, err := session.StdoutPipe()
		if err != nil {
			log.Println("failed to open stdout", err)
		}
		e, err := session.StderrPipe()
		if err != nil {
			log.Println("failed to open stderr", err)
		}
		w, err := session.StdinPipe()
		if err != nil {
			log.Println("failed to open stdin", err)
		}

		//fmt.Println("D0700", 0, filepath.Dir(path))
		//fmt.Fprintln(w, "D0700", 0, filepath.Dir(path))
		//fmt.Println("C0600", len(content), filepath.Base(path))
		fmt.Fprintln(w, "C0600", len(content), filepath.Base(path))
		//fmt.Print(string(content))
		fmt.Fprint(w, string(content))
		fmt.Fprint(w, "\x00")
		w.Close()

		log.Println("reading stdout")
		data, err := ioutil.ReadAll(r)
		if err != nil {
			log.Println("failed to read stdout", err)
		}

		log.Println("reading stderr")
		edata, err := ioutil.ReadAll(e)
		if err != nil {
			log.Println("failed to read stderr", err)
		}

		log.Println("scp response", string(data), string(edata), err)
		session.Close()
	}()

	c.StartSCP(session, filepath.Dir(path))
	return nil
}
