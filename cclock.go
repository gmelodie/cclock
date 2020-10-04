package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli"
	"github.com/k0kubun/go-ansi"
)

type normalTime struct {
	nanoseconds int
	seconds     int
	minutes     int
	hours       int
}

type centhTime struct {
	centhconds int
	centhutes  int
	centhours  int
}

// Default value of cs (in seconds)
const secToCSVal = 2.777778
const csToNanoSec = 359999970
const csToSecVal = csToNanoSec / 1e9

func toCenth(n normalTime) centhTime {
	totalsecs := float64(n.hours*3600+n.minutes*60+n.seconds) + float64(n.nanoseconds)/1e9
	totalcs := int(math.Round(totalsecs * secToCSVal))

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

func gotimeToNormalTime(t time.Time) normalTime {
	return normalTime{
		nanoseconds: t.Nanosecond(),
		seconds:     t.Second(),
		minutes:     t.Minute(),
		hours:       t.Hour(),
	}
}

func convertAndPrintSummary(normal normalTime) {
	centh := toCenth(normal)

	fmt.Printf("Normal Time:\t%2dh\t%02dmin\t%02ds\n", normal.hours, normal.minutes, normal.seconds)
	fmt.Printf("CTime:\t\t%2dch\t%02dct\t%02dcs\n", centh.centhours, centh.centhutes, centh.centhconds)
}

func moveCursorUp(lines uint) {
	ansi.Printf("\u001b[%dA", lines)
}

func main() {
	var app = cli.NewApp()

	app.Name = "cclock"
	app.Usage = "cclock's command-line interface"
	app.Version = "0.4"

	app.Commands = []cli.Command{
		{
			Name: "clock",
			Action: func(c *cli.Context) error {
				for {
					n := gotimeToNormalTime(time.Now())
					convertAndPrintSummary(n)
					time.Sleep(35999997)
					moveCursorUp(2)
				}
				return nil
			},
		},
		{
			Name:  "now",
			Usage: "Show current time (decimal and normal)",
			Action: func(c *cli.Context) error {
				n := gotimeToNormalTime(time.Now())
				convertAndPrintSummary(n)
				return nil
			},
		},
	}

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

		convertAndPrintSummary(n)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
