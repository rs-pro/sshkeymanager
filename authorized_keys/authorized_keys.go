package authorized_keys

import (
	"bytes"
	"fmt"
	"strings"
)

type SSHKey struct {
	Key   string `json:"key"`
	Email string `json:"email"`
}

func Parse(data string) ([]SSHKey, error) {
	keysStrings := strings.Split(data, "\n")
	sshKeys := make([]SSHKey, 0)
	for _, s := range keysStrings {
		if len(s) == 0 {
			continue
		}
		k := strings.Fields(s)
		var sshKey SSHKey
		sshKey.Key = k[0] + " " + k[1]
		if len(k) > 2 {
			sshKey.Email = strings.Join(k[2:], " ")
		}
		sshKeys = append(sshKeys, sshKey)
	}
	return sshKeys, nil
}

func Generate(keys []SSHKey) []byte {
	out := &bytes.Buffer{}
	for _, k := range keys {
		fmt.Fprintln(out, k.Key+" "+k.Email)
	}
	return out.Bytes()
}
