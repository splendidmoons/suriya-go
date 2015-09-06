package suriya

import (
	"fmt"
	"testing"
)

func TestAdhikamasa(t *testing.T) {
	adhikamasaYears := map[int]bool{
		1998: false, //
		1999: true,  // 3
		2000: false, //
		2001: true,  // 2
		2002: false, //
		2003: false, //
		2004: true,  // 3
		2005: false, //
		2006: false, //
		2007: true,  // 3
		2008: false, //
		2009: true,  // 2
		2010: false, //
		2011: false, //
		2012: true,  // 3
		2013: false, //
		2014: false, //
		2015: true,  // 3
		2016: false, //
	}

	for year, expect := range adhikamasaYears {
		su := SuriyaYear{}
		su.Init(year)
		res := su.Is_Adhikamasa()
		if res != expect {
			t.Errorf("%d.Is_Adhikamasa() should be %v, but got %v", su.Year, expect, res)
		}
	}
}

func TestAdhikavara(t *testing.T) {
	// Adhikavāra in FS-Cal: 2005, 2010, 2016.
	adhikavaraYears := map[int]bool{
		1998: false, //
		1999: false, //
		2000: true,  // 6
		2001: false, //
		2002: false, //
		2003: false, //
		2004: false, //
		2005: true,  // 5
		2006: false, //
		2007: false, //
		2008: false, //
		2009: false, //
		2010: true,  // 5
		2011: false, //
		2012: false, //
		2013: false, //
		2014: false, //
		2015: false, //
		2016: true,  // 6
	}

	for year, expect := range adhikavaraYears {
		su := SuriyaYear{}
		su.Init(year)
		res := su.Is_Adhikavara()
		if res != expect {
			t.Errorf("%d.Is_Adhikavara() should be %v, but got %v", su.Year, expect, res)
		}
	}
}

func TestCalculateSuriyaValues(t *testing.T) {

	var expectSuYears []SuriyaYear
	var su SuriyaYear

	// Take CE 1963, CS 1325 (as in the paper: "Rules for Interpolation...", JC Eade)
	su = SuriyaYear{
		Year:        1963,
		BE_year:     2506,
		CS_year:     1325,
		Horakhun:    483969,
		Kammacubala: 552,
		Uccabala:    1780,
		Avoman:      61,
		Masaken:     16388,
		Tithi:       23,
	}
	expectSuYears = append(expectSuYears, su)

	// Take CE 1496, CS 858 (as in "South Asian Ephemeris", JC Eade)
	// https://books.google.com/books?id=g_JEgc5C-OYC
	su = SuriyaYear{
		Year:        1496,
		BE_year:     2039,
		CS_year:     858,
		Horakhun:    313393,
		Kammacubala: 421,
		Uccabala:    2500,
		Avoman:      429,
		Masaken:     10612,
		Tithi:       15,
	}
	expectSuYears = append(expectSuYears, su)

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

	for _, expectSu := range expectSuYears {
		var su SuriyaYear
		su.Init(expectSu.Year)

		suStr := fmt.Sprintf(fmtStr, su.Year, su.BE_year, su.CS_year, su.Horakhun, su.Kammacubala, su.Uccabala, su.Avoman, su.Masaken, su.Tithi)
		expectSuStr := fmt.Sprintf(fmtStr, expectSu.Year, expectSu.BE_year, expectSu.CS_year, expectSu.Horakhun, expectSu.Kammacubala, expectSu.Uccabala, expectSu.Avoman, expectSu.Masaken, expectSu.Tithi)

		if suStr != expectSuStr {
			t.Errorf("expected: %s\n but got: %s\n", expectSuStr, suStr)
		}
	}
}
