package main

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/rs-pro/sshkeymanager"
	"github.com/urfave/cli/v2"
)

var Host = "localhost"
var Port = "22"

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
	}

	App.Commands = []*cli.Command{
		{
			Name:  "list-users",
			Usage: "list users",
			Action: func(c *cli.Context) error {
				client := sshkeymanager.NewClient(Host, Port, sshkeymanager.DefaultConfig)
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
			},
		},
		{
			Name:  "add-user",
			Usage: "add user",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "name",
					Usage: "user name",
				},
				cli.IntFlag{
					Name:  "uid",
					Usage: "user uid",
				},
				cli.IntFlag{
					Name:  "gid",
					Usage: "user gid",
				},
			},
			Action: func(c *cli.Context) error {
				projectId := c.Int64("project")
				reindexer.Reindex(projectId)
				return nil
			},
			Action: func(c *cli.Context) error {
				//client := sshkeymanager.NewClient(Host, Port, sshkeymanager.DefaultConfig)
			},
		},
		{
			Name:  "list-keys",
			Usage: "list keys",
			Action: func(c *cli.Context) error {
				//client := sshkeymanager.NewClient(Host, Port, sshkeymanager.DefaultConfig)
			},
		},
		{
			Name:  "add-key",
			Usage: "add key",
			Action: func(c *cli.Context) error {
				//client := sshkeymanager.NewClient(Host, Port, sshkeymanager.DefaultConfig)
			},
		},
	}

	err := App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
