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

// Default value of cs (in seconds)
var csVal float = 0.35999997

// Default value of ct (in seconds)
var ctVal float = 35.999997

// Default value of ch (in seconds)
var chVal float = 3599.9997

func toCenth(n normalTime) centhTime {
	return centhTime{
		n.seconds * csVal,
		n.minutes * ctVal,
		n.hours * chVal,
	}
}

func toNormal(c centhTime) normalTime {
	return normalTime{
		c.centhconds / csVal,
		c.centhutes / ctVal,
		c.centhours / chVal,
	}
}

func main() {
	var app = cli.NewApp()

	app.Name = "cclock"
	app.Usage = "cclock's command-line interface"
	app.Version = "0.2"

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
