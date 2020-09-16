package group

import (
	"log"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/davecgh/go-spew/spew"
)

type Group struct {
	Name     string
	Password string
	GID      string
	Members  string
}

func Parse(raw string) ([]Group, error) {
	strs := strings.Split(raw, "\n")

	groups := make([]Group, 0)
	for _, s := range strs {
		// skip empty lines
		if s == "" {
			continue
		}
		g := strings.Split(s, ":")
		if len(g) < 4 {
			log.Println("bad /etc/group entry:")
			spew.Dump(s)
			continue
		}
		var group Group
		group.Name = g[0]
		group.Password = g[1]
		group.GID = g[2]
		group.Members = g[3]
		groups = append(groups, group)
	}

	return groups, nil
}

func (g *Group) Serialize() string {
	return strings.Join([]string{
		g.Name,
		g.Password,
		g.GID,
		g.Members,
	}, ":")
}

func (g *Group) GroupAdd() string {
	command := []string{
		"groupadd",
	}
	if g.GID != "" {
		command = append(command, "-g "+shellescape.Quote(g.GID))
	}
	command = append(command, g.Name)
	return strings.Join(command, " ")
}

func (g *Group) GroupDelete() string {
	command := []string{
		"groupdel",
		g.Name,
	}
	return strings.Join(command, " ")
}
