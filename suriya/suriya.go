package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"github.com/soh335/ical"
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

func actionCalDays(c *cli.Context) error {
	dates := cliInit(c)

	// group the days by year
	var days_by_year = make(map[string][]suriya.CalDay)

	// GetCalDays returns sorted days
	cal_days := suriya.GetCalDays(dates["fromDate"], dates["toDate"])

	for _, day := range cal_days {
		y := fmt.Sprintf("%d", day.Date.Year())
		days_by_year[y] = append(days_by_year[y], day)
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

	return nil
}

func actionIcal(c *cli.Context) error {
	dates := cliInit(c)

	// GetCalDays returns sorted days
	cal_days := suriya.GetCalDays(dates["fromDate"], dates["toDate"])

	// https://tools.ietf.org/html/draft-ietf-calext-extensions-01

	/*
		  http://stackoverflow.com/a/17187346/195141

			BEGIN:VCALENDAR
			VERSION:2.0
			PRODID:-//My Company//NONSGML Event Calendar//EN
			URL:http://my.calendar/url
			NAME:My Calendar Name
			X-WR-CALNAME:My Calendar Name
			DESCRIPTION:A description of my calendar
			X-WR-CALDESC:A description of my calendar
			TIMEZONE-ID:Europe/London
			X-WR-TIMEZONE:Europe/London
			REFRESH-INTERVAL;VALUE=DURATION:PT12H
			X-PUBLISHED-TTL:PT12H
			COLOR:34:50:105
			CALSCALE:GREGORIAN
			METHOD:PUBLISH
	*/

	calendar := "mahanikaya"
	calendarTxt := "Mahānikāya"
	calendarName := "Uposatha Moondays (" + calendarTxt + ")"

	icalendar := ical.VCalendar{
		VERSION:      "2.0",
		PRODID:       "Uposatha Moondays " + calendarTxt + " EN",
		URL:          "http://splendidmoons.github.io/ical/" + calendar + ".ical",
		NAME:         calendarName,
		X_WR_CALNAME: calendarName,
		DESCRIPTION:  calendarName,
		X_WR_CALDESC: calendarName,
		//TIMEZONE_ID:      "Europe/London",
		//X_WR_TIMEZONE:    "Europe/London",
		REFRESH_INTERVAL: "PT12H",
		X_PUBLISHED_TTL:  "PT12H",
		COLOR:            "244:196:48", // Saffron
		CALSCALE:         "GREGORIAN",
		METHOD:           "PUBLISH",
	}

	for _, day := range cal_days {

		// UposathaMoon[0]
		// HalfMoon[0]
		// NOT AstroMoon[0]
		// MajorEvents[]
		// Events[]

		if len(day.UposathaMoon) != 0 {
			e := day.UposathaMoon[0].IcalEvent()
			icalendar.VComponent = append(icalendar.VComponent, &e)
		}

		if len(day.HalfMoon) != 0 {
			e := day.HalfMoon[0].IcalEvent()
			icalendar.VComponent = append(icalendar.VComponent, &e)
		}

		for _, d := range day.MajorEvents {
			e := d.IcalEvent()
			icalendar.VComponent = append(icalendar.VComponent, &e)
		}

		for _, d := range day.Events {
			e := d.IcalEvent()
			icalendar.VComponent = append(icalendar.VComponent, &e)
		}

	}

	buf := bytes.NewBufferString("")
	icalendar.Encode(buf)
	str := buf.String()

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
		fmt.Printf("%s", str)
	}

	return nil
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
		{
			Name:   "ical",
			Usage:  "Icalendar output for splendidmoons",
			Action: actionIcal,
			Flags:  commonFlags,
		},
	}

	app.Action = func(c *cli.Context) {
		dates := cliInit(c)
		spew.Dump(dates)
	}

	app.Run(os.Args)
}
