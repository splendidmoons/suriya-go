package suriya

import (
	"github.com/satori/go.uuid"
	"github.com/soh335/ical"
	"time"
)

type CalendarEvent interface {
	String() string
	AddToDay(day_p *CalDay)
	GetDate() time.Time
}

// A function that we can call from all the UposathaMoon, etc. types in a
// m.IcalEvent() function of each
func icalEvent(d CalendarEvent) ical.VEvent {
	e := ical.VEvent{
		UID:     uuid.NewV4().String(),
		DTSTAMP: time.Now(),
		DTSTART: d.GetDate(),
		DTEND:   d.GetDate().Add(24 * time.Hour),
		SUMMARY: d.String(),
		AllDay:  true,
		TZID:    d.GetDate().Location().String(), // TODO: is the timezone going to mess up the dates?
	}
	return e
}
