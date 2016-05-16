package suriya

import (
	"fmt"
	"github.com/soh335/ical"
	s "strings"
	"time"
)

type HalfMoon struct {
	Date  time.Time
	Phase string
}

func (m HalfMoon) String() string {
	if len(m.Phase) != 0 {
		return fmt.Sprintf("%s Moon", s.Title(m.Phase))
	}
	return ""
}

func (m HalfMoon) AddToDay(day_p *CalDay) {
	day_p.Date = m.Date
	day_p.SetHalfMoon(m)
	return
}

func (m HalfMoon) GetDate() time.Time {
	return m.Date
}

func (m HalfMoon) IcalEvent() ical.VEvent {
	return icalEvent(m)
}
