package suriya

import (
	"math"
	"strconv"
	"strings"
)

type SuriyaYear struct {
	Year        int // Common Era
	BE_year     int // Buddhist Era
	CS_year     int // Thai Era
	Horakhun    int
	Kammacubala int
	Uccabala    int
	Avoman      int
	Thaloengsok int
}

func (su SuriyaYear) Is_Adhikamasa() bool {
	t := su.Thaloengsok
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
	su.kat()
}

func (su *SuriyaYear) kat() {
	su.CS_year = su.Year - 638
	a := (su.CS_year * 292207) + 373
	su.Horakhun = int(math.Floor(float64(a)/800 + 1))
	su.Kammacubala = 800 - a%800
	//su.Uccabala = (su.Horakhun + 2611) % 3232 // This doesn't seem to be used.
	su.Avoman = ((su.Horakhun * 11) + 650) % 692
	su.Thaloengsok = int((math.Floor(((float64(su.Horakhun)*11)+650)/692) + float64(su.Horakhun))) % 30
}

func (su SuriyaYear) KatString() string {
	kat := strings.Join(
		[]string{
			strconv.Itoa(su.Kammacubala),
			strconv.Itoa(su.Avoman),
			strconv.Itoa(su.Thaloengsok),
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
