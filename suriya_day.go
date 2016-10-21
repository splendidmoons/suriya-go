package suriya

import (
	"math"
	"time"
)

type SuriyaDay struct {
	Year        int // Common Era
	BE_Year     int // Buddhist Era, CE + 543
	CS_Year     int // Chulasakkarat Era, CE - 638
	Day         int // nth day in the Lunar Year
	Date        time.Time
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

// Steps resolved with the answers at:
// http://astronomy.stackexchange.com/questions/12052/from-mean-moon-to-true-moon-in-an-old-procedural-calendar
// http://astronomy.stackexchange.com/questions/11753/how-to-interpret-this-old-degree-notation

func (suDay *SuriyaDay) Init(ce_year int, lunar_year_day int) {
	suYear := SuriyaYear{}
	suYear.Init(ce_year)

	suDay.Year = ce_year
	suDay.BE_Year = ce_year + BEdiff
	suDay.CS_Year = ce_year - CSdiff
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
	// Visakha Full + New = 15+15
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
