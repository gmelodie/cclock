package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

type normalTime struct {
	seconds float64
	minutes float64
	hours   float64
}

type centhTime struct {
	centhconds float64
	centhutes  float64
	centhours  float64
}

// Default value of cs (in seconds)
var csVal = 0.35999997

// Default value of ct (in seconds)
var ctVal = 35.999997

// Default value of ch (in seconds)
var chVal = 3599.9997

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
	app.Version = "0.3"

	app.Action = func(c *cli.Context) error {
		n := normalTime{}

		// Print usage on wrong arguments
		if c.NArg() == 0 || c.NArg() > 3 {
			cli.ShowAppHelp(c)
			return nil
		}

		if c.NArg() >= 1 {
			n.seconds, _ = strconv.ParseFloat(c.Args().Get(0), 32)
		}
		if c.NArg() >= 2 {
			n.minutes, _ = strconv.ParseFloat(c.Args().Get(1), 32)
		}
		if c.NArg() == 3 {
			n.hours, _ = strconv.ParseFloat(c.Args().Get(2), 32)
		}

		centh := toCenth(n)
		fmt.Printf("Sleeping for\t%fch\t%fct\t%fcs\n", centh.centhours, centh.centhutes, centh.centhconds)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
