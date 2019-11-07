package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "Apache Skywalking CLI"
	app.Usage = "Pleace refer to https://github.com/apache/skywalking-cli"
	app.Description = "SkyWalking CLI is a command interaction tool for the SkyWalking user or OPS team, as an alternative besides using browser GUI.\n" +
		"It is based on SkyWalking GraphQL query protocol, same as GUI."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Value: "path",
			Usage: "swctl --config `path`",
		},
		cli.StringFlag{
			Name:  "debug",
			Value: "path",
			Usage: "swctl --config `path`",
		},
		cli.StringFlag{
			Name:  "start-time,st",
			Value: "time",
		},
		cli.StringFlag{
			Name:  "end-time,et",
			Value: "time",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "service",
			Aliases: []string{"svc"},
			Usage:   "service",
			Action: func(c *cli.Context) error {
				// TODD
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "list,ls",
					Value: "Name",
					Usage: "List all available services.",
				},
				cli.StringFlag{
					Name:  "metrics-value,mv",
					Value: "path",
					Usage: "Metrics value in the given duration and metrics name.",
				},
				cli.StringFlag{
					Name:  "metrics-linear,ml",
					Value: "time",
					Usage: "Show the metrics linear graph based on response values.",
				},
			},
		},
		{
			Name:    "instance",
			Aliases: []string{"inst"},
			Usage:   "instance of service",
			Action: func(c *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "list,ls",
					Value: "Name",
					Usage: "List all available instances.",
				},
				cli.StringFlag{
					Name:  "service,svc",
					Value: "Name",
					Usage: "Set the service code in current command context.",
				},
				cli.StringFlag{
					Name:  "search",
					Value: "time",
					Usage: "Filter the instance from the existing service instance list by given --service parameter.",
				},
				cli.StringFlag{
					Name:  "metrics-value,mv",
					Value: "time",
					Usage: "Metrics value in the given duration and metrics name.",
				},
				cli.StringFlag{
					Name:  "metrics-linear,ml",
					Value: "time",
					Usage: "Show the metrics linear graph based on response values.",
				},
			},
		},
		{
			Name:    "endpoint",
			Aliases: []string{"ep"},
			Usage:   "endpoint",
			Action: func(c *cli.Context) error {
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "service,svc",
					Value: "Name",
					Usage: "(Required)Set the service code in current command context.",
				},
				cli.StringFlag{
					Name:  "search",
					Value: "Name",
					Usage: "Use the backend endpoint search capability to get the endpoint list by given --service parameter.",
				},
				cli.StringFlag{
					Name:  "metrics-value,mv",
					Value: "time",
					Usage: "Metrics value in the given duration and metrics name.",
				},
				cli.StringFlag{
					Name:  "metrics-linear,ml",
					Value: "time",
					Usage: "Show the metrics linear graph based on response values.",
				},
			},
		},
		{
			Name:    "service",
			Aliases: []string{"svc"},
			Usage:   "service",
			Action: func(c *cli.Context) error {
				// TODD
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
