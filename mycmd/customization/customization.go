package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	cli.BashCompletionFlag = cli.BoolFlag{
		Name:   "compgen",
		Hidden: true,
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "wat",
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
