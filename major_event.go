package suriya

import (
	"github.com/soh335/ical"
	"time"
)

/*
The major calendar events for convenient access.
- Magha Puja
- Vesakha Puja
- Asalha Puja
- Vassa begins
- Pavarana Day
- Vassa ends
*/

type MajorEvent Event

func (e MajorEvent) AddToDay(day_p *CalDay) {
	day_p.Date = e.Date
	day_p.MajorEvents = append(day_p.MajorEvents, e)
	return
}

func (e MajorEvent) GetDate() time.Time {
	return e.Date
}

func (e MajorEvent) String() string {
	return e.Summary
}

func (d MajorEvent) IcalEvent() ical.VEvent {
	return icalEvent(d)
}
