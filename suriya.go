package suriya

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Eade, p.10. South Asian traditional number of days in 800 years

const (
	EraDays          = 292207
	EraYears         = 800
	EraHorakhun      = 373 // The Horakhun at the beginning of the CS Era, Ephemeris p.15, H2 element
	EraUccabala      = 2611
	EraAvoman        = 650
	EraMasaken       = 0
	MonthLength      = 30
	CycleTithi       = 703 // For every 692 solar days that elapse there are also 703 tithi = 692 + 11 / 692
	CycleSolar       = 692
	CycleDaily       = 11
	KammacubalaDaily = 800 // Daily increase
)

// Whether to apply the (adhikavāra) exceptions where the official calendar
// differed from the formulas. Default is false, to generate calendar data that
// is "pure" in its consistency. Set to true if you want to match official past
// calendars which differed from the regular pattern.
var UseExceptions bool = false

var AdhikavaraExceptions = map[int]bool{
	1994: false,
	1997: true,
}

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
}

type SuriyaDay struct {
	Year        int // Common Era
	BE_Year     int // Buddhist Era, CE + 543
	CS_Year     int // Chulasakkarat Era, CE - 638
	Day         int // nth day in the Lunar Year
	Horakhun    int
	Kammacubala int
	Uccabala    int
	Avoman      int
	Masaken     int
	Tithi       int
	MeanSun     float64 // position in degrees
	TrueSun     float64
	MeanMoon    float64
	TrueMoon    float64
	Raek        float64
}

type UposathaMoon struct {
	Date          time.Time
	Calendar      int    // mahanikaya, dhammayut, srilanka, myanmar
	Status        int    // draft, predicted, confirmed
	Phase         string // only new or full. waxing and waning will be derived.
	Event         string // magha, vesakha, asalha, pavarana
	S_Number      int    // 1 of 8 in Hemanta
	S_Total       int    // total number of uposathas in the season, 8 in Hemanta
	U_Days        int    // uposatha days, 14 or 15
	M_Days        int    // month days, 29 or 30
	LunarMonth    int    // 1-12, 13 is 2nd Asalha (adhikamasa). Odd numbers are 30 day months.
	LunarSeason   int    // 1-3, an int code to an []string array of names
	LunarYear     int
	HasAdhikavara bool
}

type HalfMoon struct {
	Date  time.Time
	Phase string
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
	su.BE_Year = su.Year + 543
	su.CS_Year = su.Year - 638
	su.calculateSuriyaValues()
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

func (su *SuriyaYear) calculateSuriyaValues() {
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

// (x; y : z) in Eade's notation means 30*60*x + 60*y + z in arcmins, so x and y are deg originally
func DegreeToRal(degree float64) (x, y, z int) {
	// how many times 30 degrees
	x = int(math.Floor(degree / 30))

	// the remainder degrees
	y = int(math.Floor(degree)) % 30

	// plus the arcmins
	z = int(math.Floor((degree - math.Floor(degree)) * 60))

	// TODO This is rasi, angsa, lipda
	return x, y, z
}

func DegreeToRalString(degree float64) string {
	u, v, t := DegreeToRal(degree)
	return fmt.Sprintf("%d:%d°%d'", u, v, t)
}

func RalToDegree(x, y, z int) float64 {
	// Multiply up and divide down by 10000 for better arcmin (z) accuracy
	// Floor to keep only 4 decimal places
	return math.Floor(float64(30*x+y)*10000+float64(z*10000)/60) / 10000
}

// Keep it within 360 deg
func NormalizeDegree(deg float64) float64 {
	if deg <= 360 {
		return deg
	}
	return deg - math.Floor(deg/360)*360
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

// Steps resolved with the answers at:
// http://astronomy.stackexchange.com/questions/12052/from-mean-moon-to-true-moon-in-an-old-procedural-calendar
// http://astronomy.stackexchange.com/questions/11753/how-to-interpret-this-old-degree-notation

func (suDay *SuriyaDay) Init(ce_year int, lunar_year_day int) {
	suYear := SuriyaYear{}
	suYear.Init(ce_year)

	suDay.Year = ce_year
	suDay.BE_Year = ce_year + 543
	suDay.CS_Year = ce_year - 638
	suDay.Day = lunar_year_day

	// This is elapsedDays = suDay.Horakhun - suYear.Horakhun, but the meaning is
	// perhaps clearer as below.
	elapsedDays := suDay.Day - suYear.Tithi

	// Horakhun of the day
	suDay.Horakhun = suYear.Horakhun + elapsedDays

	// Kammacubala of the day
	suDay.Kammacubala = KammacubalaDaily - (suDay.CS_Year*EraDays+EraHorakhun)%EraYears + elapsedDays*KammacubalaDaily

	// Uccabala of the day
	suDay.Uccabala = (suDay.Horakhun + EraUccabala) % 3232

	var ai, bi int // int helpers

	// Avoman of the day
	ai = (suDay.Horakhun * CycleDaily) + EraAvoman
	suDay.Avoman = ai % CycleSolar

	// Masaken of the day
	bi = int(math.Floor(float64(ai)/CycleSolar)) + EraMasaken + suDay.Horakhun
	suDay.Masaken = int(math.Floor(float64(bi / MonthLength)))

	// Tithi of the day
	suDay.Tithi = bi % MonthLength

	var a, b float64

	// === B. Find the position of the Mean and true Sun on Asalha 15 ===

	// Sample values in the comments are for lunar_year_day = 103, Asalha 15

	// Length of the months, Thai months ending on New Moon:
	// Citta   Full + New = 15+14
	// Vesakha Full + New = 15+15
	// Jettha  Full + New = 15+14
	// Asalha  Full       = 15
	// ---------------------------
	//                    = 103

	// interval from 1 Caitra (aka Citta) to Asalha Full Moon, minus New Year day

	a = float64((elapsedDays * EraYears) + suYear.Kammacubala)
	// a = 64552

	b = (a / EraDays) * 360
	// b = 79.5282796100025

	// The -3 arcmin is a geographical correction. Mentioned in "Interpolation..." and "Calendrical".

	// (x; y : z) in Eade's notation means 30*60*x + 60*y + z in arcmins, so x and y are deg originally
	x, y, z := DegreeToRal(b)
	z -= 3

	// Do convert the degree to Ral and back. If we only do b -= 3/60, we get
	// slightly different results than in Eade's papers.

	suDay.MeanSun = RalToDegree(x, y, z)
	// MeanSun = 2; 19 : 28
	// MeanSun = 79.4666

	// The -80 degree is mentioned in Calendrical, sth to do with the Sun's Apogee?

	a = math.Abs(suDay.MeanSun - 80)

	// math.Sin takes radians
	radconv := math.Pi / 180
	b = math.Floor(134 * math.Sin(a*radconv))
	// b = math.Floor(1.2473)
	// b = 1

	// Floor it to get degree only to 4th decimal place, to avoid results such as TrueSun: 79.48326666666667
	suDay.TrueSun = math.Floor(suDay.MeanSun*10000+(b*10000)/60) / 10000
	// TrueSun = 2; 19 : 29

	// === C. Find the Mean and True Moon on Asalha 15 ===

	// step 12.

	// divide with 60 to covert value in degrees from minutes
	a = (float64(suDay.Avoman) + math.Floor(float64(suDay.Avoman)/25)) / 60
	// 0; 4 : 17
	// b = 4.3

	// step 13.

	/* The -40 arcmin is a geographical correction. In "Interpolation...": The
	routine subtraction of 3 arcmins is a geographical longitude correction for
	the sun, as is the subtraction of 40 arcmins for the moon (sec. C13). */

	// Use RalToDegree() instead of 40/60. RalToDegree() gives only a four decimal
	// place value, which produces results closer to Eade's papers.

	suDay.MeanMoon = NormalizeDegree(suDay.TrueSun + a + (float64(suDay.Tithi) * 12) - RalToDegree(0, 0, 40))
	// Mean Moon: 8; 11 : 7
	// Mean Moon: 251.116666

	// step 14.

	var meanUccabala float64

	// all in one, see below for step-by-step
	meanUccabala = ((((float64(suYear.Uccabala+elapsedDays) * 3 * 30) / 808) * 60) + 2) / 60
	// Mean Uccabala = 6; 27 : 12
	// Mean Uccabala = 207.2115

	/*
		Multiply with 30 to conform with (x; y : z) = 30*60*x + 60*y + z

		808 / 30 is 26.9333, perhaps reproducing the length of the lunar month.

		meanUccabala *= 30

		Which gives Mean Uccabala = 6; 27 : 10

		Convert to arcmin:

		meanUccabala *= 60

		Add 2, possibly correction for geographical position

		meanUccabala += 2

		Convert back to degree:

		meanUccabala = meanUccabala / 60

		Mean Uccabala = 6; 27 : 12
	*/

	// step 15.

	a = suDay.MeanMoon - meanUccabala
	// b = 1; 13 : 54
	// b = 43.9051

	// NOTE: Eade has 1; 3 : 55, but that doesn't work. This is a typo in the paper.

	// step 16.

	b = (296 * math.Sin(a*radconv)) / 60
	// d = 0; 3 : 24
	// d = 3.4

	// step 17.

	suDay.TrueMoon = math.Floor((suDay.MeanMoon-b)*10000) / 10000
	// True Moon = 8; 7 : 43
	// True Moon = 247.716666

	// (0; 13:20) = 13.33 degree is one raek, i.e. 360 deg / 27 mansions
	a = RalToDegree(0, 13, 20)
	// Raek aka Mula
	suDay.Raek = suDay.TrueMoon/a + 1
	// Raek = 0; 19 : 34
	// Raek = 19.5771

}

var monthToInt = map[string]int{
	"null":       0,
	"magasira":   1,
	"phussa":     2,
	"magha":      3,
	"phagguna":   4,
	"citta":      5,
	"vesakha":    6,
	"jettha":     7,
	"asalha":     8,
	"savana":     9,
	"bhaddapada": 10,
	"assayuja":   11,
	"kattika":    12,
	"2nd asalha": 13,
}

func MonthToInt(month string) int {
	return monthToInt[month]
}

var seasonToInt = map[string]int{
	"null":    0,
	"hemanta": 1,
	"gimhana": 2,
	"vassana": 3,
}

func SeasonToInt(season string) int {
	return seasonToInt[season]
}

var seasonName = map[int]string{
	0: "",
	1: "Hemanta",
	2: "Gimha",
	3: "Vassāna",
}

func SeasonName(number int) string {
	return seasonName[number]
}

var calendarToInt = map[string]int{
	"mahanikaya": 0,
	"dhammayut":  1,
	"srilanka":   2,
	"myanmar":    3,
}

func CalendarToInt(calendar string) int {
	return calendarToInt[calendar]
}

var statusToInt = map[string]int{
	"draft":     0,
	"predicted": 1,
	"confirmed": 2,
}

func StatusToInt(status string) int {
	return statusToInt[status]
}

func NextUposatha(last_uposatha UposathaMoon) UposathaMoon {

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

		// Event: magha, vesakha, asalha, pavarana

		// In Adhikamāsa Years the major moons shift with one month
		if is_adhikamasa_year {
			switch nu.LunarMonth {
			case 4:
				nu.Event = "magha"
			case 7:
				nu.Event = "vesakha"
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
				nu.Event = "vesakha"
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

// Calculate the kattika full moon before this year
func CalculatePreviousKattika(solar_year int) time.Time {

	dFmt := "2006-01-02"

	// Step from a known Kattika date as epoch date
	kattika_date, _ := time.Parse(dFmt, "2015-11-25")

	// Determine the direction of stepping
	var direction int
	if kattika_date.Year() < solar_year-1 {
		direction = 1
	} else if kattika_date.Year() > solar_year-1 {
		direction = -1
	}

	// Step in direction until the Kattika in the prev. solar year
	for y := kattika_date.Year(); y != solar_year-1; y += direction {
		var su_year SuriyaYear
		var n int
		if direction == 1 {
			su_year.Init(y + 1)
		} else {
			su_year.Init(y)
		}
		n = 6*29 + 6*30
		if su_year.Is_Adhikamasa() {
			n += 30
		} else if su_year.Is_Adhikavara() {
			n += 1
		}
		kattika_date = kattika_date.Add(time.Duration(n*direction) * time.Hour * 24)
	}

	return kattika_date
}

type SimpleCalDay struct {
	Date  time.Time
	Phase string // full, waning, new, waxing
	Event string // magha, vesakha, asalha, pavarana
}

func GenerateSolarYear(solar_year int) []SimpleCalDay {

	var days []SimpleCalDay

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
		LunarYear:   date.Year() + 543,
	}

	for last_uposatha.Date.Year() <= solar_year {
		var uposatha UposathaMoon
		var day SimpleCalDay
		uposatha = NextUposatha(last_uposatha)
		last_uposatha = uposatha

		// Uposatha

		day.Date = uposatha.Date
		day.Phase = uposatha.Phase
		day.Event = uposatha.Event

		if uposatha.Date.Year() == solar_year {
			days = append(days, day)
		}

		// Half Moon

		var phase string
		switch uposatha.Phase {
		case "new":
			phase = "waxing"
		case "full":
			phase = "waning"
		}

		day = SimpleCalDay{
			Date:  uposatha.Date.AddDate(0, 0, 8),
			Phase: phase,
			Event: "",
		}

		if day.Date.Year() == solar_year {
			days = append(days, day)
		}

	}

	return days
}
