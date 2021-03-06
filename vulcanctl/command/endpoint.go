package command

import (
	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func NewEndpointCommand(cmd *Command) cli.Command {
	return cli.Command{
		Name:  "endpoint",
		Usage: "Operations with vulcan endpoint",
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Usage:  "Add a new endpoint to location",
				Action: cmd.addEndpointAction,
				Flags: []cli.Flag{
					cli.StringFlag{"id", "", "endpoint id, autogenerated if empty"},
					cli.StringFlag{"upstream, up", "", "upstream id"},
					cli.StringFlag{"url", "", "url in form <scheme>://<host>:<port>"},
				},
			},
			{
				Name:  "rm",
				Usage: "Remove endpoint from location",
				Flags: []cli.Flag{
					cli.StringFlag{"id", "", "endpoint id"},
					cli.StringFlag{"upstream, up", "", "upstream id"},
				},
				Action: cmd.deleteEndpointAction,
			},
		},
	}
}

func (cmd *Command) addEndpointAction(c *cli.Context) {
	cmd.printStatus(cmd.client.AddEndpoint(c.String("up"), c.String("id"), c.String("url")))
}

func (cmd *Command) deleteEndpointAction(c *cli.Context) {
	cmd.printStatus(cmd.client.DeleteEndpoint(c.String("up"), c.String("id")))
}
