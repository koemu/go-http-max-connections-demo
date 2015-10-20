package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/koemu/go-http-max-connections-demo/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "default",
		Usage:  "Default commection mode.",
		Action: command.CmdDefault,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port, p",
				Value: 8080,
				Usage: "Set httpd listen port",
			},
		},
	},
	{
		Name:   "limited",
		Usage:  "Limited connection mode.",
		Action: command.CmdLimited,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port, p",
				Value: 8080,
				Usage: "Set httpd listen port",
			},
			cli.IntFlag{
				Name:  "max-connections, m",
				Value: 5,
				Usage: "Set httpd max connections",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
