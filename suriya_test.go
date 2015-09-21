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

func TestAsalhaPuja(t *testing.T) {
	/*
		testYears := map[int]string{
			1957: "1957-07-12", // PASS thaiorc.com
			1958: "1958-07-31", // FAIL thaiorc.com GOT: 1958-07-30
			1959: "1959-07-20", // FAIL thaiorc.com GOT: 1959-07-21
			1960: "1960-07-08", // PASS thaiorc.com
			1961: "1961-07-27", // FAIL thaiorc.com GOT: 1961-08-27
			1962: "1962-07-16", // PASS thaiorc.com

			1983: "1983-07-24", // thaiorc.com
			1984: "1984-07-12", // thaiorc.com
			1985: "1985-07-31", // thaiorc.com
			1986: "1986-07-20", // thaiorc.com
			1987: "1987-07-10", // thaiorc.com
			1988: "1988-07-28", // thaiorc.com
			1989: "1989-07-17", // thaiorc.com
			1990: "1990-07-07", // thaiorc.com
			1991: "1991-07-26", // thaiorc.com
			1992: "1992-07-14", // thaiorc.com
			1993: "1993-08-02", // thaiorc.com
			1994: "1994-07-22", // thaiorc.com
			1995: "1995-07-11", // thaiorc.com
			1996: "1996-07-29", // thaiorc.com
			1997: "1997-07-19", // thaiorc.com
			1998: "1998-07-08", // thaiorc.com
			1999: "1999-07-27", // thaiorc.com
			2000: "2000-07-16", // thaiorc.com
			2001: "2001-07-05", // thaiorc.com
			2002: "2002-07-24", // thaiorc.com
			2003: "2003-07-13", // thaiorc.com
	*/

	/*
		//1992: "1992-07-15", // FAIL bot.or.th GOT: 1992-07-14. 1992 could have been a "before" adhikavāra.
		//2001: "2001-07-05", // FAIL fs-cal GOT: 2001-08-04. 2001 is adhikamāsa.
		2004: "2004-07-31", // PASS fs-cal
		2005: "2005-07-22", // PASS bot.or.th NOTE: fs-cal has 07-21
		2006: "2006-07-10", // PASS fs-cal NOTE: bot.or.th has 07-11
		2007: "2007-07-29", // PASS fs-cal has 07-29 NOTE: bot.or.th says official date was 07-30, substitution day
		2008: "2008-07-17", // PASS bot.or.th, fs-cal
		//2009: "2009-07-07", // FAIL bot.or.th, fs-cal GOT: 2009-08-05. 2009 is adhikamāsa. 2009 is adhikavāra in fs-cal.
		//2010: "2010-07-26", // FAIL bot.or.th, fs-cal GOT: 2010-07-27. 2010 is adhikavāra. 2010 too is adhikavāra in fs-cal.
		2011: "2011-07-15", // PASS bot.or.th, fs-cal
		2012: "2012-08-02", // PASS bot.or.th, fs-cal
		2013: "2013-07-22", // PASS bot.or.th, fs-cal
		2014: "2014-07-11", // PASS bot.or.th, fs-cal
		2015: "2015-07-30", // PASS bot.or.th, fs-cal
		2016: "2016-07-19", // PASS bot.or.th, fs-cal
	*/

	/*
		}
	*/

	/*
		for year, expect := range testYears {
			su := SuriyaYear{}
			su.Init(year)
			asalha := su.AsalhaPuja()
			asalhaStr := asalha.Format("2006-01-02")
			if asalhaStr != expect {
				t.Errorf("expected %s, but got %s", expect, asalhaStr)
			}
		}
	*/
}

/*
func TestAsalhaPujaStepping(t *testing.T) {
	testYears := map[int]string{
		//1992: "1992-07-15", // FAIL bot.or.th GOT: 1992-07-14. 1992 could have been a "before" adhikavāra.
		//2001: "2001-07-05", // FAIL fs-cal GOT: 2001-08-04. 2001 is adhikamāsa.
		2004: "2004-07-31", // PASS fs-cal
		2005: "2005-07-22", // PASS bot.or.th NOTE: fs-cal has 07-21
		2006: "2006-07-10", // PASS fs-cal NOTE: bot.or.th has 07-11
		2007: "2007-07-29", // PASS fs-cal has 07-29 NOTE: bot.or.th says official date was 07-30, substitution day
		2008: "2008-07-17", // PASS bot.or.th, fs-cal
		//2009: "2009-07-07", // FAIL bot.or.th, fs-cal GOT: 2009-08-05. 2009 is adhikamāsa. 2009 is adhikavāra in fs-cal.
		//2010: "2010-07-26", // FAIL bot.or.th, fs-cal GOT: 2010-07-27. 2010 is adhikavāra. 2010 too is adhikavāra in fs-cal.
		2011: "2011-07-15", // PASS bot.or.th, fs-cal
		2012: "2012-08-02", // PASS bot.or.th, fs-cal
		2013: "2013-07-22", // PASS bot.or.th, fs-cal
		2014: "2014-07-11", // PASS bot.or.th, fs-cal
		2015: "2015-07-30", // PASS bot.or.th, fs-cal
		2016: "2016-07-19", // PASS bot.or.th, fs-cal
	}

	for year, expect := range testYears {
		su := SuriyaYear{}
		su.Init(year)
		asalha := su.AsalhaPujaStepping()
		asalhaStr := asalha.Format("2006-01-02")
		if asalhaStr != expect {
			t.Errorf("expected %s, but got %s", expect, asalhaStr)
		}
	}
}
*/
