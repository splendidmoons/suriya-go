package suriya

import (
	"math"
	"strconv"
	"strings"
)

type SuriyaYear struct {
	Year        int // Common Era
	Cs_year     int // Thai Era
	Horakhun    int
	Kammacubala int
	Uccabala    int
	Avoman      int
	Thaloengsok int
}

func (su SuriyaYear) Is_Adhikamasa() bool {
	t := su.Thaloengsok

	// >= 25 doesn't work
	//return (t >= 25 && t <= 29) || (t >= 0 && t <= 6)
	return (t > 25 && t <= 29) || (t >= 0 && t <= 6)

	//return (t >= 25 && t <= 29) || (t >= 0 && t < 6)
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
	su.kat()
}

func (su *SuriyaYear) kat() {
	su.Cs_year = su.Year - 638
	a := (su.Cs_year * 292207) + 373
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
		return su.Avoman <= 126
	} else {
		return su.Avoman <= 137
	}
}

func (su SuriyaYear) Has_Carried_Adhikavara() bool {
	last_year := SuriyaYear{}
	last_year.Init(su.Year - 1)
	return last_year.Is_Adhikamasa() && last_year.Would_Be_Adhikavara()
}
