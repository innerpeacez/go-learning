package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Apache SkyWalking CLI"
	app.Description = "SkyWalking CLI is a command interaction tool for the SkyWalking user or OPS team, as an alternative besides using browser GUI.\n It is based on SkyWalking GraphQL query protocol, same as GUI."
	app.Usage = "https://github.com/apache/skywalking-cli"
	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "lang", Value: "english", Usage: "Language for the greeting"},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
