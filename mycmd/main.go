package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "Make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Boom ! I say!")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
