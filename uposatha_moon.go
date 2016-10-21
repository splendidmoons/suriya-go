package suriya

import (
	"fmt"
	"github.com/soh335/ical"
	s "strings"
	"time"
)

type UposathaMoon struct {
	Date          time.Time
	Calendar      int    // 0 mahanikaya, 1 dhammayut, 2 srilanka, 3 myanmar
	Status        int    // 0 draft, 1 predicted, 2 confirmed
	Phase         string // only new or full. waxing and waning will be derived.
	Event         string // magha, visakha, asalha, pavarana
	S_Number      int    // 1 of 8 in Hemanta
	S_Total       int    // total number of uposathas in the season, 8 in Hemanta
	U_Days        int    // uposatha days, 14 or 15
	M_Days        int    // month days, 29 or 30
	LunarMonth    int    // 1-12, 13 is 2nd Asalha (adhikamasa). Odd numbers are 30 day months.
	LunarSeason   int    // 1-3, an int code to an []string array of names
	LunarYear     int
	HasAdhikavara bool
	Source        string
	Comments      string
}

func (last_uposatha UposathaMoon) NextUposatha() UposathaMoon {

	lu := last_uposatha
	var nu UposathaMoon // next uposatha

	var su_year SuriyaYear
	su_year.Init(lu.Date.Year())

	is_adhikamasa_year := su_year.Is_Adhikamasa()
	is_adhikavara_year := su_year.Is_Adhikavara()

	nu.Status = 0   // predicted
	nu.Calendar = 0 // mahanikaya

	// Alternating New Moon and Full Moon uposathas.

	if lu.Phase == "new" {
		nu.Phase = "full"
	} else {
		nu.Phase = "new"
	}

	if nu.Phase == "full" {

		// A Full Moon uposatha is always 15 days in the same month, season and year as the last uposatha.

		nu.S_Number = lu.S_Number + 1
		nu.S_Total = lu.S_Total
		nu.U_Days = 15
		nu.M_Days = lu.M_Days
		nu.LunarMonth = lu.LunarMonth
		nu.LunarSeason = lu.LunarSeason
		nu.LunarYear = lu.LunarYear
		nu.HasAdhikavara = false // Adhikavara is only added to New Moons

		// Event: magha, visakha, asalha, pavarana

		// In Adhikamāsa Years the major moons shift with one month
		if is_adhikamasa_year {
			switch nu.LunarMonth {
			case 4:
				nu.Event = "magha"
			case 7:
				nu.Event = "visakha"
			case 13:
				nu.Event = "asalha"
			case 11:
				nu.Event = "pavarana"
			default:
				nu.Event = ""
			}
		} else {
			// Common Year and Adhikavara Year
			switch nu.LunarMonth {
			case 3:
				nu.Event = "magha"
			case 6:
				nu.Event = "visakha"
			case 8:
				nu.Event = "asalha"
			case 11:
				nu.Event = "pavarana"
			default:
				nu.Event = ""
			}
		}

	} else {

		// The New Moon uposatha begins a new month.

		if lu.LunarMonth == 13 {
			nu.LunarMonth = 9 // Savana after 2nd Asalha
		} else if lu.LunarMonth == 8 && is_adhikamasa_year {
			nu.LunarMonth = 13 // 2nd Asalha
		} else if lu.LunarMonth == 12 {
			nu.LunarMonth = 1
		} else {
			nu.LunarMonth = lu.LunarMonth + 1
		}

		// Odd numbered months are 30 days, except in adhikavāra years when the 8th month is 30 days.

		if is_adhikavara_year && nu.LunarMonth == 8 {
			nu.HasAdhikavara = true
			nu.M_Days = 30
		} else {
			if nu.LunarMonth%2 == 1 {
				nu.M_Days = 30
			} else {
				nu.M_Days = 29
			}
		}

		if nu.M_Days == 29 {
			nu.U_Days = 14
		} else {
			nu.U_Days = 15
		}

		// Season

		// In an adhikamāsa year the Hot Season is 10 uposatha long

		if is_adhikamasa_year && ((nu.LunarMonth >= 5 && nu.LunarMonth <= 8) || nu.LunarMonth == 13) {
			nu.S_Total = 10
		} else {
			nu.S_Total = 8
		}

		// If the last uposatha was not the last of the season, increment

		if lu.S_Number < lu.S_Total {
			nu.S_Number = lu.S_Number + 1
			nu.LunarSeason = lu.LunarSeason
			nu.LunarYear = lu.LunarYear
		} else {

			// Else, it is the first uposatha of the season

			nu.S_Number = 1
			// is it a new lunar year?
			if lu.LunarMonth == 12 {
				nu.LunarSeason = 1
				nu.LunarYear = lu.LunarYear + 1
			} else {
				nu.LunarSeason = lu.LunarSeason + 1
				nu.LunarYear = lu.LunarYear
			}
		}
	}

	nu.Date = lu.Date.Add(time.Duration(nu.U_Days) * time.Hour * 24)

	return nu
}

func (m UposathaMoon) AddToDay(day_p *CalDay) {
	day_p.Date = m.Date
	day_p.SetUposathaMoon(m)
	return
}

func (m UposathaMoon) GetDate() time.Time {
	return m.Date
}

func (m UposathaMoon) String() string {
	if len(m.Phase) != 0 {
		return fmt.Sprintf("%s Moon - %d day %s %d/%d", s.Title(m.Phase), m.U_Days, SeasonName(m.LunarSeason), m.S_Number, m.S_Total)
	}
	return ""
}

func (m UposathaMoon) IcalEvent() ical.VEvent {
	return icalEvent(m)
}
