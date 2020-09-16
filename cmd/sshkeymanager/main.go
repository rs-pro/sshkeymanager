package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/rs-pro/sshkeymanager"
	"github.com/rs-pro/sshkeymanager/group"
	"github.com/rs-pro/sshkeymanager/passwd"
	"github.com/urfave/cli/v2"
)

var Host = "localhost"
var Port = "22"
var User = "root"

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
		},
		&cli.StringFlag{
			Name:        "port",
			Usage:       "ssh port",
			Value:       "22",
			Destination: &Port,
		},
		&cli.StringFlag{
			Name:        "user",
			Usage:       "ssh user",
			Value:       "root",
			Destination: &User,
		},
	}

	App.Commands = []*cli.Command{
		{
			Name:  "list-groups",
			Usage: "list groups",
			Action: func(c *cli.Context) error {
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}
				g := &group.Group{}
				g.GID = c.String("gid")
				g.Name = c.String("name")
				g.Members = c.String("members")

				g, err = client.AddGroup(g)
				if err != nil {
					return err
				}
				log.Println("added group:", g.GID, g.Name, g.Members)
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}
				g := &group.Group{}
				g.GID = c.String("gid")
				g.Name = c.String("name")
				if g.Name == "" && g.GID != "" {
					g = client.FindGroup(g)
				}

				_, err = client.DeleteGroup(g)
				if err != nil {
					return err
				}
				log.Println("deleted group:", g.GID, g.Name)
				return nil
			},
		},
		{
			Name:  "list-users",
			Usage: "list users",
			Action: func(c *cli.Context) error {
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}
				u := &passwd.User{}
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

				u, err = client.AddUser(u, createHome)
				if err != nil {
					return err
				}

				log.Println("added user:", u.Name, u.UID, u.GID, u.Home)
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}
				u := &passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				deleteHome := c.Bool("delete-home")

				u, err = client.DeleteUser(u, deleteHome)
				if err != nil {
					return err
				}
				log.Println("deleted user:", u.Name, u.UID, u.GID)
				return nil
			},
		},
		{
			Name:  "list-keys",
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}
				u := &passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				u = client.FindUser(u)
				if u == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", u.Name, u.UID, u.GID)
				keys, err := client.GetKeys(*u)
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}

				u := &passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				u = client.FindUser(u)
				if u == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", u.Name, u.UID, u.GID)
				args := c.Args().Slice()
				err = client.AddKey(*u, strings.Join(args, " "))
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
				client, err := sshkeymanager.NewClient(Host, Port, User, sshkeymanager.DefaultConfig())
				if err != nil {
					return err
				}

				u := &passwd.User{}
				u.Name = c.String("name")
				u.UID = c.String("uid")
				u = client.FindUser(u)
				if u == nil {
					return errors.New("user not found")
				}
				log.Println("found user:", u.Name, u.UID, u.GID)
				args := c.Args().Slice()
				err = client.DeleteKey(*u, strings.Join(args, " "))
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
