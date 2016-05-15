package suriya

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type SuriyaYear struct {
	Year        int // Common Era
	BE_Year     int // Buddhist Era, CE + 543
	CS_Year     int // Chulasakkarat Era, CE - 638
	Horakhun    int // Elapsed days of the era, aka Ahargana or Sawana
	Kammacubala int // Remaining 800ths of a day
	Uccabala    int // Age of the moon's Apogee
	Avoman      int // For the Moon's mean motion
	Masaken     int // Elapsed months of the era
	Tithi       int // Age of the moon at the start of the year, aka Thaloengsok or New Year's Day
	FirstDay    time.Time
}

func (su SuriyaYear) Is_Adhikamasa() bool {
	// If next year also qualifies for adhikamāsa, then this year isn't
	var su_next SuriyaYear
	su_next.Init(su.Year + 1)
	return !su_next.Would_Be_Adhikamasa() && su.Would_Be_Adhikamasa()
}

func (su SuriyaYear) Would_Be_Adhikamasa() bool {
	t := su.Tithi
	// Eade says t >= 25, but then 2012 (t=24) would not be adhikamāsa.
	return (t >= 24 && t <= 29) || (t >= 0 && t <= 5)
}

func (su SuriyaYear) Is_Adhikavara() bool {
	if UseExceptions {
		if _, ok := AdhikavaraExceptions[su.Year]; ok {
			return AdhikavaraExceptions[su.Year]
		}
	}

	if su.Is_Adhikamasa() {
		return false
	}
	if su.Has_Carried_Adhikavara() {
		return true
	} else {
		return su.Would_Be_Adhikavara()
	}
}

func (su SuriyaYear) String() string {
	n := strconv.Itoa(su.Year)
	return n
}

func (su *SuriyaYear) Init(ce_year int) {
	su.Year = ce_year
	su.BE_Year = su.Year + BEdiff
	su.CS_Year = su.Year - CSdiff

	var a, b int // just helper variables

	// Take CE 1963, CS 1325 (as in the paper: "Rules for Interpolation...")

	// === A. Find the relevant values for the astronomical New Year ===

	// +1 is another constant correction, H3
	a = (su.CS_Year * EraDays) + EraHorakhun
	su.Horakhun = int(math.Floor(float64(a/KammacubalaDaily + 1)))
	// Horakhun = 483969

	su.Kammacubala = KammacubalaDaily - a%KammacubalaDaily
	// Kammacubala = 552

	su.Uccabala = (su.Horakhun + EraUccabala) % 3232
	// Uccabala = 1780

	a = (su.Horakhun * CycleDaily) + EraAvoman
	su.Avoman = a % CycleSolar
	// Avoman = 61

	b = int(math.Floor(float64(a) / CycleSolar))
	su.Masaken = int(math.Floor(float64((b + EraMasaken + su.Horakhun) / MonthLength)))
	// Masaken = 16388

	su.Tithi = (b + su.Horakhun) % MonthLength
	// Tithi = 23
}

func (su *SuriyaYear) SuriyaValuesString() string {
	fmtStr := `CE: %d
BE: %d
CS: %d
Horakhun: %d
Kammacubala: %d
Uccabala: %d
Avoman: %d
Masaken: %d
Tithi: %d
`

	return fmt.Sprintf(fmtStr, su.Year, su.BE_Year, su.CS_Year, su.Horakhun, su.Kammacubala, su.Uccabala, su.Avoman, su.Masaken, su.Tithi)
}

func (su SuriyaYear) Is_Suriya_Leap() bool {
	return su.Kammacubala <= 207
}

/*
Eade, in Rules for Interpolation...:

> if the kammacubala value is 207 or less, then the year is a leap year.
> in a leap year, if the avoman is 126 or less, the year will have an extra day
> in a normal year, if the avoman is 137 or less the year will have an extra day.
*/

func (su SuriyaYear) Would_Be_Adhikavara() bool {
	if su.Is_Suriya_Leap() {
		// Both <= and < seems to work. Eade phrases it as <=.
		return su.Avoman <= 126
	} else {
		// Eade says Avoman <= 137, but that doesn't work.
		return su.Avoman < 137
	}
}

func (su SuriyaYear) Has_Carried_Adhikavara() bool {
	last_year := SuriyaYear{}
	last_year.Init(su.Year - 1)
	return last_year.Is_Adhikamasa() && last_year.Would_Be_Adhikavara()
}

// Determine the position in the 57 year cycle. Assume 1984 = 1, 2040 = 57, 2041 = 1.
func (su SuriyaYear) AdhikavaraCyclePos() int {
	return int(math.Abs(float64(1984-57*10-su.Year)))%57 + 1
}

// Determine the position in the 19 year cycle.
func (su SuriyaYear) AdhikamasaCyclePos() int {
	return int(math.Abs(float64(1984-19*10-su.Year)))%19 + 1
}

// Years since last adhikamāsa.
func (su SuriyaYear) DeltaAdhikamasa() int {
	for year := su.Year - 1; true; year-- {
		check := SuriyaYear{}
		check.Init(year)
		if check.Is_Adhikamasa() {
			return su.Year - check.Year
		}
		// Avoid looking forever.
		if su.Year-check.Year > 6 {
			break
		}
	}
	return -1
}

// Years since last adhikavāra.
func (su SuriyaYear) DeltaAdhikavara() int {
	for year := su.Year - 1; true; year-- {
		check := SuriyaYear{}
		check.Init(year)
		if check.Is_Adhikavara() {
			return su.Year - check.Year
		}
		// Avoid looking forever.
		if su.Year-check.Year > 12 {
			break
		}
	}
	return -1
}

// Length of the lunar year in days
func (su SuriyaYear) YearLength() int {
	// In a common year, there are six alternating 29 and 30 day lunar months.
	days := 6 * (30 + 29)
	if su.Is_Adhikamasa() {
		// In an adhikamāsa year, there is an extra 30 day month.
		days = days + 30
	} else if su.Is_Adhikavara() {
		// In an adhikavāra year, there is an extra day.
		days = days + 1
	}
	return days
}

// Date of Asalha Puja
func (su SuriyaYear) AsalhaPuja() time.Time {
	// In a common year, Asalha Puja is the last day of the 8th month.
	days := 4 * (29 + 30)
	if su.Is_Adhikamasa() {
		// In an adhikamāsa year, the extra month (2nd Asalha) is a 30 day month.
		days = days + 30
	} else if su.Is_Adhikavara() {
		// In an adhikavāra year, the 8th month (Asalha) is 30 days instead of 29 days.
		days = days + 1
	}

	prev_kattika := CalculatePreviousKattika(su.Year)
	date := prev_kattika.Add(time.Duration(days) * time.Hour * 24)
	return date
}
