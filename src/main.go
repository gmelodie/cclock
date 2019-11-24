package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

type normalTime struct {
	seconds int
	minutes int
	hours   int
}

type centhTime struct {
	centhconds int
	centhutes  int
	centhours  int
}

func toCenth(normal normalTime) centhTime {
	return centhTime{}
}

func main() {
	var app = cli.NewApp()

	app.Name = "cclock"
	app.Usage = "cclock's command-line interface"
	app.Version = "0.1"

	app.Action = func(c *cli.Context) error {
		n := normalTime{}

		// Print usage on wrong arguments
		if c.NArg() == 0 || c.NArg() > 3 {
			cli.ShowAppHelp(c)
			return nil
		}

		if c.NArg() >= 1 {
			n.seconds, _ = strconv.Atoi(c.Args().Get(0))
		}
		if c.NArg() >= 2 {
			n.minutes, _ = strconv.Atoi(c.Args().Get(1))
		}
		if c.NArg() == 3 {
			n.hours, _ = strconv.Atoi(c.Args().Get(2))
		}

		fmt.Println(n)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
