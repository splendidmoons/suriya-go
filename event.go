package suriya

import (
	"github.com/soh335/ical"
	"time"
)

/*
Events are notes for a date in a calendar. They may be anniversaries, Kathinas or
occasional notes.
*/

type Event struct {
	Date        time.Time
	Calendar    int // mahanikaya, dhammayut, srilanka, myanmar
	Summary     string
	Description string
}

func (e Event) AddToDay(day_p *CalDay) {
	day_p.Date = e.Date
	day_p.Events = append(day_p.Events, e)
	return
}

func (e Event) GetDate() time.Time {
	return e.Date
}

func (e Event) String() string {
	return e.Summary
}

func (d Event) IcalEvent() ical.VEvent {
	return icalEvent(d)
}
