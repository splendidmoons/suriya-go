package suriya

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SuriyaYear struct {
	Year        int // Common Era
	BE_year     int // Buddhist Era, CE - 543
	CS_year     int // Chulasakkarat Era, CE - 638
	Horakhun    int // Elapsed days of the era
	Kammacubala int // Remaining 800ths of a day
	Uccabala    int // Age of the moon's Apogee
	Avoman      int // For the Moon's mean motion
	Masaken     int // Elapsed months of the era
	Tithi       int // Age of the moon at the start of the year, also called Thaloengsok or New Year's Day
	//MeanSun     float64
	//TrueSun     float64
	//MeanMoon    float64
	//TrueMoon    float64
}

func (su SuriyaYear) Is_Adhikamasa() bool {
	t := su.Tithi
	// TODO: check this against the definition again in the papers.
	return (t >= 21 && t <= 29) || (t >= 0 && t <= 1)
}

func (su SuriyaYear) Is_Adhikavara() bool {
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
	su.BE_year = su.Year + 543
	su.CS_year = su.Year - 638
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

	return fmt.Sprintf(fmtStr, su.Year, su.BE_year, su.CS_year, su.Horakhun, su.Kammacubala, su.Uccabala, su.Avoman, su.Masaken, su.Tithi)
}

func (su *SuriyaYear) calculateSuriyaValues() {
	// Eade, p.10. South Asian traditional number of days in 800 years
	const eraDays = 292207
	const eraYears = 800
	var a int // just a helper variable

	// Take CE 1963, CS 1325 (as in the paper: "Rules for Interpolation...")

	// === A. Find the relevant values for the astronomical New Year ===

	a = (su.CS_year * eraDays) + 373
	su.Horakhun = int(math.Floor(float64(a/800 + 1)))
	// Horakhun = 483969

	su.Kammacubala = 800 - a%800
	// Kammacubala = 552

	su.Uccabala = (su.Horakhun + 2611) % 3232
	// Uccabala = 1780

	su.Avoman = ((su.Horakhun * 11) + 650) % 692
	// Avoman = 61

	a = int(math.Floor(float64(((su.Horakhun * 11) + 650) / 692)))
	su.Masaken = (a + su.Horakhun) / 30
	// Masaken = 16388

	su.Tithi = (a + su.Horakhun) % 30
	// Tithi = 23

	// === B. Find the position of the Mean and true Sun on Asalha 15 ===

	/*
		// TODO
		// interval from 1 Caitra (Visakha New Moon - 1) to Ashalha Full Moon, minus NY day
		deltaVA := 103 - su.Tithi
		// deltaVA = 80

		a = (deltaVA * eraYears) + su.Kammacubala
		su.MeanSun = float64(((a / eraDays) * 360)) - math.Pow(3, 13)
		fmt.Printf("MeanSun: %v\n", su.MeanSun)
	*/

	// === C. Find the Mean and True Moon on Asalha 15 ===

}

func (su SuriyaYear) KatString() string {
	kat := strings.Join(
		[]string{
			strconv.Itoa(su.Kammacubala),
			strconv.Itoa(su.Avoman),
			strconv.Itoa(su.Tithi),
		}, ",",
	)
	return kat
}

func (su SuriyaYear) Is_Suriya_Leap() bool {
	return su.Kammacubala <= 207
}

func (su SuriyaYear) Would_Be_Adhikavara() bool {
	if su.Is_Suriya_Leap() {
		// TODO: is it <= or < ?
		return su.Avoman <= 126
	} else {
		// TODO: <= 137 doesn't work. check this in the papers.
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
	return int(math.Abs(float64(1984-su.Year)))%57 + 1
}

// Determine the position in the 19 year cycle.
func (su SuriyaYear) AdhikamasaCyclePos() int {
	return int(math.Abs(float64(1984-su.Year)))%19 + 1
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
		if su.Year-check.Year > 3 {
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
		if su.Year-check.Year > 6 {
			break
		}
	}
	return -1
}
