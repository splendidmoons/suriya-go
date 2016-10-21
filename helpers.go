package suriya

import (
	"fmt"
	"math"
	"time"
)

// (x; y : z) in Rasi, Angsa (degree), Lipda (minute) means 30*60*x + 60*y + z
// in arcmins, so x and y are deg originally
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

func HorakhunToDate(horakhun int64) time.Time {
	// Make sure it is not a pointer to horakhunRefDate, but is the same time.
	date, _ := time.Parse("2006 Jan 2", horakhunRefStr)
	// At midnight
	date = date.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	var delta, stepDays, direction int64

	// Duration is max 290 solar years. Increment the date in 290 year steps.

	delta = horakhun - horakhunRef
	if delta < 0 {
		direction = -1
	} else {
		direction = 1
	}
	stepDays = 290 * 356
	for math.Abs(float64(delta)) > float64(stepDays) {
		date = date.Add(time.Duration(direction*stepDays) * time.Hour * 24)
		delta += direction * -1 * stepDays
	}
	// Add any remaining delta.
	date = date.Add(time.Duration(delta) * time.Hour * 24)

	return date
}

/* TODO: Rewrite this to calculate based on year values. Stepping like this only
works for a few years backward and forward. CE 1288 is off for example. */

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

var monthToInt = map[string]int{
	"null":       0,
	"magasira":   1,
	"phussa":     2,
	"magha":      3,
	"phagguna":   4,
	"citta":      5,
	"visakha":    6,
	"jettha":     7,
	"asalha":     8,
	"2nd asalha": 13,
	"savana":     9,
	"bhaddapada": 10,
	"assayuja":   11,
	"kattika":    12,
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
