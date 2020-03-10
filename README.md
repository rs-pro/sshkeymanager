# ssh-key-manager
```
$ go get github.com/rs-pro/ssh-key-manager
```

## example:
```go
package main

import (
	"fmt"
	sshkeymanager "github.com/rs-pro/ssh-key-manager"
)

func main() {

	rootUser := "root"
	host := "server.name"
	port := "22"
	
	users := sshkeymanager.GetUsers(rootUser, host, port)
	
	for _, u := range users {
		fmt.Printf("UID: %v\nUsername: %v\nHome dir: %v\nShell: %v\n\n", u.UID, u.Name, u.Home, u.Shell)
	}
	
	uid := "3104"
	
	keys := sshkeymanager.GetKeys(uid, rootUser, host, port)
	for _, k := range keys {
		fmt.Printf("String num: %s\nKey: %s\nEmail: %v\n\n", k.Num, k.Key, k.Email)
	}

	key1 := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqF4hRYsFzO3ylja7wPxut+vu6y2VhYmfOz5wMHuP7XpUvoK/O6Red4bOUAPgexHzRw5kRAKYnaIoMPjNQYCSIhr5xNLwkZTWBxKQ48pLkuBC0yrm+ePXe8sjdFq/0ctPMYX2ZAKYUledoAeb/JbE+zPCEnzhUUqq9pkqGkJJ7I3Fp6uaRx+DELYggIHs6gqWgXLHGdaGkGPNs1xoG4EFwHOx51Jlp1IKAktRjooM9rqPV/TUkM02CoR0VncWbkgDja2lSywdFb8e8keFvbBSPYsB40VMSpXroRJjQ5eQyJlaVyuodXkKGuJmd/5lEZrtQQLISspAjYF2cFgJSsvzz mail1@example.com"
        key2 := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCYI5i3jdeetGQ+qPSDUHGM8xt3hciswqARvquiWG9C6bFhdLHvhzTXB+qYmOrGDPZfd6cb8pT9AJ3G94+o1vTXCkbhOyT7I2DS5UfroQ1thgiFSv90jJNHWC2vhFVSdN1x14DpuCk1jlZTzeW0fZ2a6/vX3OUcLWmiGiT1AhDKgcvGH0j1NZYmYOZl+pd5WN7EAj/dZPjHQt72mUTPMfppKdl3yJS3WD2Lp0nMmL43buvMeoGRMZm8Fu8U36xuNX4GWf4dlTSh5nYs/A85mDGixrOvSu8F+vEv38A5Ua88mUxuAC9M102VxdgTN3exaUxTlz07JhZeCInxn+hQCkLj mail2@example.com"

	
	err := sshkeymanager.AddKey(key1, uid, rootUser, host, port)
    if err != nil {
        fmt.Println(err)
    }   
	err = sshkeymanager.DeleteKey(key2, uid, rootUser, host, port)
    if err != nil {
        fmt.Println(err)
    }   

}
```
