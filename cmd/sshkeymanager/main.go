package main

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/rs-pro/sshkeymanager"
	"github.com/rs-pro/sshkeymanager/authorized_keys"
	"github.com/rs-pro/sshkeymanager/client"
	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
	"github.com/urfave/cli/v2"
)

var Host = "localhost"
var Port = "22"
var User = "root"
var KeyServer = ""
var ApiKey = ""

func getClient() (sshkeymanager.ClientInterface, error) {
	if KeyServer != "" || ApiKey != "" {
		if KeyServer == "" {
			return nil, errors.New("no key server provided (or remove api key)")
		}
		if ApiKey == "" {
			return nil, errors.New("no api key provided (or remove key server)")
		}
		client := client.NewClient(KeyServer, ApiKey).WithConfig(Host, Port, User)
		client.ApiComment = "sshkeymanager-cli"
		return client, nil
	} else {
		return sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
	}
}

func main() {
	App := cli.NewApp()
	App.EnableBashCompletion = true
	App.Name = "sshkeymanager"
	App.Usage = "sshkeymanager command"
	App.Version = "0.0.1"

	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version, v",
		Usage: "print only the version",
	}

	App.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Usage:       "ssh host",
			Value:       "localhost",
			Destination: &Host,
			EnvVars:     []string{"SSH_HOST"},
		},
		&cli.StringFlag{
			Name:        "port",
			Usage:       "ssh port",
			Value:       "22",
			Destination: &Port,
			EnvVars:     []string{"SSH_PORT"},
		},
		&cli.StringFlag{
			Name:        "user",
			Usage:       "ssh user",
			Value:       "root",
			Destination: &User,
			EnvVars:     []string{"SSH_USER"},
		},
		&cli.StringFlag{
			Name:        "keyserver",
			Usage:       "keymanager server",
			Value:       "",
			Destination: &KeyServer,
			EnvVars:     []string{"KEY_SERVER"},
		},
		&cli.StringFlag{
			Name:        "apikey",
			Usage:       "keymanager server api key",
			Value:       "",
			Destination: &ApiKey,
			EnvVars:     []string{"API_KEY"},
		},
	}

	App.Commands = []*cli.Command{
		{
			Name:  "get-groups",
			Usage: "get groups",
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				groups, err := client.GetGroups()
				if err != nil {
					return err
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"GID", "Name", "Password", "Members"})
				for _, group := range groups {
					table.Append([]string{
						group.GID,
						group.Name,
						group.Password,
						group.Members,
					})
				}
				table.Render()
				return nil
			},
		},
		{
			Name:  "find-group",
			Usage: "find group",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "group name",
				},
				&cli.StringFlag{
					Name:  "gid",
					Usage: "gid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				g := group.Group{}
				g.GID = c.String("gid")
				g.Name = c.String("name")
				gr, err := client.FindGroup(g)
				if err != nil {
					return err
				}
				if gr == nil {
					log.Println("group not found")
				} else {
					log.Println("group found:", gr.GID, gr.Name)
				}

				return nil
			},
		},
		{
			Name:  "add-group",
			Usage: "add group",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "group name",
				},
				&cli.StringFlag{
					Name:  "gid",
					Usage: "gid",
				},
				&cli.StringFlag{
					Name:  "members",
					Usage: "members",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				g := group.Group{}
				g.GID = c.String("gid")
				g.Name = c.String("name")
				g.Members = c.String("members")

				gr, err := client.AddGroup(g)
				if err != nil {
					return err
				}
				log.Println("added group:", gr.GID, gr.Name, gr.Members)
				return nil
			},
		},
		{
			Name:  "delete-group",
			Usage: "delete group",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "group name",
				},
				&cli.StringFlag{
					Name:  "gid",
					Usage: "gid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				g := group.Group{}
				g.GID = c.String("gid")
				g.Name = c.String("name")

				gr, err := client.DeleteGroup(g)
				if err != nil {
					return err
				}
				log.Println("deleted group:", gr.GID, gr.Name)
				return nil
			},
		},
		{
			Name:  "get-users",
			Usage: "get users",
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				users, err := client.GetUsers()
				if err != nil {
					return err
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"UID", "GID", "Name", "Home", "Shell"})
				for _, user := range users {
					table.Append([]string{
						user.UID,
						user.GID,
						user.Name,
						user.Home,
						user.Shell,
					})
				}
				table.Render()
				return nil
			},
		},
		{
			Name:  "find-user",
			Usage: "find user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				user, err := client.FindUser(u)
				if err != nil {
					return err
				}
				if user == nil {
					log.Println("user not found")
				} else {
					log.Println("user found:", u.Name, u.UID, u.GID, u.Home)
				}
				return nil
			},
		},
		{
			Name:  "add-user",
			Usage: "add user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
				&cli.StringFlag{
					Name:  "gid",
					Usage: "user gid",
				},
				&cli.StringFlag{
					Name:  "home",
					Usage: "home dir",
				},
				&cli.StringFlag{
					Name:  "password",
					Usage: "user password",
				},
				&cli.StringFlag{
					Name:  "shell",
					Usage: "user shell",
					Value: "/bin/bash",
				},
				&cli.BoolFlag{
					Name:  "create-home",
					Usage: "create user's home dir",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				u.GID = c.String("gid")
				u.Home = c.String("home")
				u.Shell = c.String("shell")
				if c.String("password") != "" {
					pw := c.String("password")
					u.SetPassword(pw)
				}
				createHome := c.Bool("create-home")

				us, err := client.AddUser(u, createHome)
				if err != nil {
					return err
				}

				log.Println("added user:", us.Name, us.UID, us.GID, us.Home)
				return nil
			},
		},
		{
			Name:  "delete-user",
			Usage: "delete user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
				&cli.BoolFlag{
					Name:  "delete-home",
					Usage: "delete user's home dir",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				deleteHome := c.Bool("delete-home")

				us, err := client.DeleteUser(u, deleteHome)
				if err != nil {
					return err
				}
				log.Println("deleted user:", us.Name, us.UID, us.GID)
				return nil
			},
		},
		{
			Name:  "get-keys",
			Usage: "list keys",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}
				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				us, err := client.FindUser(u)
				if err != nil {
					return err
				}
				if us == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", us.Name, us.UID, us.GID)
				keys, err := client.GetKeys(*us)
				if err != nil {
					return err
				}
				for _, key := range keys {
					log.Println(key.Email, key.Key)
				}

				return nil
			},
		},
		{
			Name:  "add-key",
			Usage: "add key",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}

				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				us, err := client.FindUser(u)
				if err != nil {
					return err
				}
				if us == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", us.Name, us.UID, us.GID)

				args := c.Args().Slice()
				keys, err := authorized_keys.Parse(args[0])
				if us == nil {
					return errors.Wrap(err, "failed to parse key")
				}
				err = client.AddKey(*us, keys[0])
				if err != nil {
					return err
				}

				return nil
			},
		},

		{
			Name:  "delete-key",
			Usage: "delete key",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Usage: "user name",
				},
				&cli.StringFlag{
					Name:  "uid",
					Usage: "user uid",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := getClient()
				if err != nil {
					return err
				}

				u := passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				us, err := client.FindUser(u)
				if err != nil {
					return err
				}
				if us == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", us.Name, us.UID, us.GID)

				args := c.Args().Slice()
				keys, err := authorized_keys.Parse(args[0])
				if us == nil {
					return errors.Wrap(err, "failed to parse key")
				}
				err = client.DeleteKey(*us, keys[0])
				if err != nil {
					return err
				}

				return nil
			},
		},
	}

	err := App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
