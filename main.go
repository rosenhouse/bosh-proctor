package main

import (
	"github.com/onsi/say"
	"github.com/rosenhouse/proctor/commands"
)

func main() {
	say.Invoke(say.Executable{
		Name: "proctor",
		CommandGroups: []say.CommandGroup{
			{
				Name:        "Actions",
				Description: "Classroom management",
				Commands: []say.Command{
					commands.NewCreateCommand(),
					commands.NewDescribeCommand(),
					commands.NewDestroyCommand(),
					commands.NewListCommand(),
					commands.NewRunCommand(),
				},
			},
		},
	})
}
