package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"github.com/splendidmoons/suriya-go"
	"log"
	"os"
	"time"
)

const (
	isoDateFmt = "2006-01-02"
)

func cliInit(c *cli.Context) (dates map[string]time.Time) {
	var err error

	dates = make(map[string]time.Time)

	if len(c.String("from")) > 0 {
		dates["fromDate"], err = time.Parse("2006-01-02", c.String("from"))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	} else {
		dates["fromDate"] = time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	}

	if len(c.String("to")) > 0 {
		dates["toDate"], err = time.Parse(isoDateFmt, c.String("to"))
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	} else {
		dates["toDate"] = time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	}

	return dates
}

func actionCalDays(c *cli.Context) {
	dates := cliInit(c)

	// group the days by year
	var days_by_year = make(map[string][]suriya.CalDay)

	for year := dates["fromDate"].Year(); year <= dates["toDate"].Year(); year++ {
		cal_days := suriya.GetCalDays(dates["fromDate"], dates["toDate"])
		days_by_year[fmt.Sprintf("%d", year)] = cal_days
	}

	a, err := json.Marshal(days_by_year)
	str := string(a)
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}

	if len(c.String("output")) > 0 {
		f, err := os.OpenFile(c.String("output"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("%v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		_, err = f.WriteString(str)
		if err != nil {
			log.Printf("%v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s\n", str)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "suriya"
	app.Usage = "uposatha moondays on the command line"

	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Usage: "from date as YYYY-MM-DD, defaults to Jan 1st of this year",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "to date as YYYY-MM-DD, defaults to Dec 31 of this year",
		},
		cli.StringFlag{
			Name:  "output",
			Usage: "output file name",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "caldays",
			Usage:  "CalDays JSON output for splendidmoons",
			Action: actionCalDays,
			Flags:  commonFlags,
		},
	}

	app.Action = func(c *cli.Context) {
		dates := cliInit(c)
		spew.Dump(dates)
	}

	app.Run(os.Args)
}
