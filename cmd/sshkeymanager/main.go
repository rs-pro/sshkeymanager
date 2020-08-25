package main

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/rs-pro/sshkeymanager"
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

				u, err = client.AddUser(u)
				if err != nil {
					return err
				}
				log.Println("added user:", u.Name, u.UID, u.GID)
				return nil
			},
		},
		{
			Name:  "list-keys",
			Usage: "list keys",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "add-key",
			Usage: "add key",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
