# golang ssh key manager

Includes a golang library, API server and client library, and a CLI tool with support for working locally or via a remote server.

## Project status:

Beta. Key management and most extra features (api, client-server mode) work.

#### Known issues:

- no locking is present (when attempting to add multiple keys from multiple goroutines to one user at the same time, chances are only one will be added)

## Security contact

Please be responsible and contact me directly at glebtv [at] gmail.com in case you find security issues.

### Not intended for web-facing applications.

#### No safeguards or security measures included besides API Key.

#### You should implement your own.

```
$ go get github.com/rs-pro/sshkeymanager
```

## Examples:

##### Get users:
```go
package main

import (
	"fmt"
	"github.com/rs-pro/sshkeymanager"
	"log"
	)

func main() {
    	host := "host.name"
    	port := "22"

			client := sshkeymanager.NewClient(host, port, sshkeymanager.DefaultConfig)
			users, err := client.GetUsers()

    	users, err := c.GetUsers()
    	if err != nil {
    		log.Println(err)
    	}

    	for _, u := range users {
    		fmt.Printf("UID: %v\nUsername: %v\nHome dir: %v\nShell: %v\n\n", u.UID, u.Name, u.Home, u.Shell)
    	}

```
##### Get user keys:
```go
    	uid := "3104"
    	keys, err := c.GetKeys(uid)
    	if err != nil {
    		log.Println(err)
    	}
    	for _, k := range keys {
    		fmt.Printf("String num: %d\nKey: %s\nEmail: %v\n\n", k.Num, k.Key, k.Email)
    	}
```

##### Add key
```go
    	key1 := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqF4hRYsFzO3ylja7wPxut+vu6y2VhYmfOz5wMHuP7XpUvoK/O6Red4bOUAPgexHzRw5kRAKYnaIoMPjNQYCSIhr5xNLwkZTWBxKQ48pLkuBC0yrm+ePXe8sjdFq/0ctPMYX2ZAKYUledoAeb/JbE+zPCEnzhUUqq9pkqGkJJ7I3Fp6uaRx+DELYggIHs6gqWgXLHGdaGkGPNs1xoG4EFwHOx51Jlp1IKAktRjooM9rqPV/TUkM02CoR0VncWbkgDja2lSywdFb8e8keFvbBSPYsB40VMSpXroRJjQ5eQyJlaVyuodXkKGuJmd/5lEZrtQQLISspAjYF2cFgJSsvzz mail1@example.com"
    	err = c.AddKey(key1, uid)
    	if err != nil {
    		fmt.Println(err)
    	}
```
##### Delete key
```go
        err = c.DeleteKey(key1, uid)
           	if err != nil {
           		fmt.Println(err)
           	}
        // Closing connection
    	err = c.CloseConnection()
}
```

## Client-server mode

Create config.yml (see example)

```
KEY_PASS=your-ssh-key-passphrase go run cmd/sshkeyserver/main.go
go run cmd/sshkeymanager/main.go --keyserver http://localhost:12010 --apikey your-key --host r8s.ru list-users
```

## Running specs

```
env DEBUG=YES INSECURE_IGNORE_HOST_KEY=YES go test -v ./...
```

## License

MIT License
