package main

import (
	"fmt"
	"log"
	"math"
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
var secToCSVal = 2.777778
var csToSecVal = 0.35999997

func toCenth(n normalTime) centhTime {
	totalsecs := n.hours*3600 + n.minutes*60 + n.seconds
	totalcs := int(math.Round(float64(totalsecs) * secToCSVal))

	res := centhTime{}
	res.centhours = totalcs / 10000
	totalcs = totalcs - 10000*res.centhours
	res.centhutes = totalcs / 100
	totalcs = totalcs - 100*res.centhutes
	res.centhconds = totalcs

	return res
}

func toNormal(c centhTime) normalTime {
	totalcs := c.centhours*10000 + c.centhutes*100 + c.centhconds
	totalsecs := int(math.Round(float64(totalcs) * csToSecVal))

	res := normalTime{}
	res.hours = totalsecs / 3600
	totalsecs = totalsecs - 3600*res.hours
	res.minutes = totalsecs / 60
	totalsecs = totalsecs - 60*res.minutes
	res.seconds = totalsecs

	return res
}

func main() {
	var app = cli.NewApp()

	app.Name = "cclock"
	app.Usage = "cclock's command-line interface"
	app.Version = "0.4"

	app.Action = func(c *cli.Context) error {
		n := normalTime{}

		// Print usage on wrong arguments
		if c.NArg() == 0 || c.NArg() > 3 {
			cli.ShowAppHelp(c)
			return nil
		}

		if c.NArg() >= 1 {
			n.hours, _ = strconv.Atoi(c.Args().Get(0))
		}
		if c.NArg() >= 2 {
			n.minutes, _ = strconv.Atoi(c.Args().Get(1))
		}
		if c.NArg() == 3 {
			n.seconds, _ = strconv.Atoi(c.Args().Get(2))
		}

		centh := toCenth(n)
		fmt.Printf("Normal Time:\t%dh\t%dmin\t%ds\n", n.hours, n.minutes, n.seconds)
		fmt.Printf("CTime:\t%dch\t%dct\t%dcs\n", centh.centhours, centh.centhutes, centh.centhconds)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
