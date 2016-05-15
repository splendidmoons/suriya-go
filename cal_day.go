package suriya

import (
	"errors"
	"sort"
	s "strings"
	"time"
)

// Use a slice for the moons, with only one item, so that empty values are not
// filled with initialized defaults

type CalDay struct {
	Date         time.Time
	UposathaMoon UposathaMoonSliceSingle `json:",omitempty"`
	HalfMoon     HalfMoonSliceSingle     `json:",omitempty"`
	AstroMoon    AstroMoonSliceSingle    `json:",omitempty"`
	MajorEvents  []MajorEvent            `json:",omitempty"`
	Events       []Event                 `json:",omitempty"`
}

type UposathaMoonSliceSingle []UposathaMoon
type HalfMoonSliceSingle []HalfMoon
type AstroMoonSliceSingle []AstroMoon
type CalDaySlice []CalDay

func (c CalDay) GetUposathaMoon() UposathaMoon {
	var m UposathaMoon
	if len(c.UposathaMoon) != 0 {
		m = c.UposathaMoon[0]
	} else {
		m = UposathaMoon{}
	}
	return m
}

func (c *CalDay) SetUposathaMoon(m UposathaMoon) {
	if len(c.UposathaMoon) != 0 {
		c.UposathaMoon[0] = m
	} else {
		c.UposathaMoon = append(c.UposathaMoon, m)
	}
}

func (c CalDay) GetAstroMoon() AstroMoon {
	var m AstroMoon
	if len(c.AstroMoon) != 0 {
		m = c.AstroMoon[0]
	} else {
		m = AstroMoon{}
	}
	return m
}

func (c *CalDay) SetAstroMoon(m AstroMoon) {
	if len(c.AstroMoon) != 0 {
		c.AstroMoon[0] = m
	} else {
		c.AstroMoon = append(c.AstroMoon, m)
	}
}

func (c CalDay) GetHalfMoon() HalfMoon {
	var m HalfMoon
	if len(c.HalfMoon) != 0 {
		m = c.HalfMoon[0]
	} else {
		m = HalfMoon{}
	}
	return m
}

func (c *CalDay) SetHalfMoon(m HalfMoon) {
	if len(c.HalfMoon) != 0 {
		c.HalfMoon[0] = m
	} else {
		c.HalfMoon = append(c.HalfMoon, m)
	}
}

func (c CalDay) String() string {
	// TODO Do better. also MajorEvents and Events.
	a := []string{c.GetUposathaMoon().String(), c.GetAstroMoon().String(), c.GetHalfMoon().String()}
	return s.Join(a, ", ")
}

func (s CalDaySlice) Len() int {
	return len(s)
}

func (s CalDaySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s CalDaySlice) Less(i, j int) bool {
	return s[j].Date.After(s[i].Date)
}

func findCalDay(cal_days []CalDay, date time.Time) (*CalDay, error) {
	date_day := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	for k, day := range cal_days {
		day_day := time.Date(day.Date.Year(), day.Date.Month(), day.Date.Day(), 0, 0, 0, 0, day.Date.Location())
		if day_day.Equal(date_day) {
			return &cal_days[k], nil
		}
	}
	return &CalDay{}, errors.New("CalDay not found")
}

func AddHalfMoonDays(cal_days []CalDay) []CalDay {
	var half_moon_days []CalDay

	for _, day := range cal_days {
		// If there is an uposatha on the day
		if len(day.UposathaMoon) != 0 {
			// Eight days from its date
			eighth_day := CalDay{Date: day.Date.AddDate(0, 0, 8)}

			// determine the phase of the half moon
			var phase string
			switch day.GetUposathaMoon().Phase {
			case "new":
				phase = "waxing"
			case "full":
				phase = "waning"
			default:
				// just skip in case of an invalid phase
				continue
			}

			// add a new half-moon
			moon := HalfMoon{
				Date:  eighth_day.Date,
				Phase: phase,
			}
			moon.AddToDay(&eighth_day)

			half_moon_days = append(half_moon_days, eighth_day)
		}
	}

	// add half-moons to the calendar days
	cal_days = append(cal_days, half_moon_days...)

	return cal_days
}

func mergeIntoCalDays(cal_days []CalDay, event CalendarEvent) []CalDay {
	day_p, err := findCalDay(cal_days, event.GetDate())
	if err != nil {
		// no such day in the slice, so add a new one
		event.AddToDay(day_p)
		cal_days = append(cal_days, *day_p)
	} else {
		event.AddToDay(day_p)
	}
	return cal_days
}

func GetCalDays(fromDate time.Time, toDate time.Time) []CalDay {
	var cal_days []CalDay

	for _, d := range GetAstroMoons(fromDate, toDate) {
		cal_days = mergeIntoCalDays(cal_days, d)
	}

	for year := fromDate.Year(); year <= toDate.Year(); year++ {
		for _, d := range GenerateSolarYear(year) {
			if d.GetDate().Before(fromDate) || d.GetDate().After(toDate) {
				continue
			} else {
				cal_days = mergeIntoCalDays(cal_days, d)
			}
		}
	}

	sort.Sort(CalDaySlice(cal_days))
	return cal_days
}

func GenerateSolarYear(solar_year int) []CalendarEvent {
	var events []CalendarEvent

	var su_year SuriyaYear
	su_year.Init(solar_year)

	date := CalculatePreviousKattika(solar_year)

	last_uposatha := UposathaMoon{
		Date:        date,
		Calendar:    0, // mahanikaya
		Phase:       "full",
		S_Number:    8,
		S_Total:     8,
		U_Days:      15,
		M_Days:      29,
		LunarMonth:  12,
		LunarSeason: 3,
		LunarYear:   date.Year() + BEdiff,
	}

	for last_uposatha.Date.Year() <= solar_year {
		var uposatha UposathaMoon
		uposatha = last_uposatha.NextUposatha()
		last_uposatha = uposatha

		// Uposatha

		// assume confirmed
		uposatha.Status = 2

		if uposatha.Date.Year() == solar_year {
			events = append(events, uposatha)
		}

		// Half Moon

		var phase string
		switch uposatha.Phase {
		case "new":
			phase = "waxing"
		case "full":
			phase = "waning"
		}

		halfmoon := HalfMoon{
			Date:  uposatha.Date.AddDate(0, 0, 8),
			Phase: phase,
		}

		if halfmoon.GetDate().Year() == solar_year {
			events = append(events, halfmoon)
		}

		// Major Events

		// If it is this year, Full Moon, and Mahanikaya
		if uposatha.Date.Year() == solar_year &&
			uposatha.Phase == "full" &&
			uposatha.Calendar == CalendarToInt("mahanikaya") {

			var e MajorEvent

			// Magha Puja
			if uposatha.LunarMonth == MonthToInt("magha") {
				e = MajorEvent{
					Date:        uposatha.Date,
					Calendar:    uposatha.Calendar,
					Summary:     "Māgha Pūjā",
					Description: "Māgha Pūjā",
				}
				events = append(events, e)
			}

			// Visakha Puja
			if uposatha.LunarMonth == MonthToInt("visakha") {
				e = MajorEvent{
					Date:        uposatha.Date,
					Calendar:    uposatha.Calendar,
					Summary:     "Visākha Pūjā",
					Description: "Visākha Pūjā",
				}
				events = append(events, e)
			}

			// Asalha Puja
			// check for 2nd Asalha too
			if uposatha.LunarMonth == MonthToInt("2nd asalha") ||
				uposatha.LunarMonth == MonthToInt("asalha") {
				e = MajorEvent{
					Date:        uposatha.Date,
					Calendar:    uposatha.Calendar,
					Summary:     "Āsāḷha Pūjā",
					Description: "Āsāḷha Pūjā",
				}
				events = append(events, e)

				// First day of Vassa
				e = MajorEvent{
					Date:        uposatha.Date.AddDate(0, 0, 1),
					Calendar:    uposatha.Calendar,
					Summary:     "First day of Vassa",
					Description: "First day of Vassa",
				}
				events = append(events, e)
			}

			// Pavarana Day
			if uposatha.LunarMonth == MonthToInt("assayuja") {
				e = MajorEvent{
					Date:        uposatha.Date,
					Calendar:    uposatha.Calendar,
					Summary:     "Pavāraṇā Day",
					Description: "Pavāraṇā Day",
				}
				events = append(events, e)

				// Last day of Vassa
				e = MajorEvent{
					Date:        uposatha.Date,
					Calendar:    uposatha.Calendar,
					Summary:     "Last day of Vassa",
					Description: "Last day of Vassa",
				}
				events = append(events, e)
			}
		}
	}

	return events
}
