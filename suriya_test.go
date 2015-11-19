package suriya

import (
	"fmt"
	"testing"
)

func TestAdhikamasa(t *testing.T) {
	adhikamasaYears := map[int]bool{
		// --- T = thaiorc.com, M = myhora.com, F = fs-cal, K = Khemanando
		//              T M F K
		1950: true,  // 0 0
		1951: false, //
		1952: false, //
		1953: true,  // 3 3
		1954: false, //
		1955: false, //
		1956: true,  // 3 3
		1957: false, //
		1958: true,  // 2 2
		1959: false, //
		1960: false, //
		1961: true,  // 3 3
		1962: false, //
		1963: false, //
		1964: true,  // 3 3
		1965: false, //
		1966: true,  // 2 2
		1967: false, //
		1968: false, //
		1969: true,  // 3 3
		1970: false, //
		1971: false, //
		1972: true,  // 3 3
		1973: false, //
		1974: false, //
		1975: true,  // 3 3
		1976: false, //
		1977: true,  // 2 2
		1978: false, //
		1979: false, //
		1980: true,  // 3 3
		1981: false, //
		1982: false, //
		1983: true,  // 3 3
		1984: false, //
		1985: true,  // 2 2  	K
		1986: false, //
		1987: false, //
		1988: true,  // 3 3  	K
		1989: false, //
		1990: false, //       K
		1991: true,  // 3 3
		1992: false, //
		1993: true,  // 2 2  	K
		1994: false, //
		1995: false, //
		1996: true,  // 3 3  	K
		1997: false, //
		1998: false, //
		1999: true,  // 3 3  	K
		2000: false, //
		2001: false, //       K
		2002: true,  // 3 3 ?
		2003: false, //
		2004: true,  // 2 2 2	K
		2005: false, //
		2006: false, //
		2007: true,  // 3 3 3
		2008: false, //
		2009: false, //
		2010: true,  // 3 3 3
		2011: false, //
		2012: true,  // 2 2 2
		2013: false, //
		2014: false, //
		2015: true,  // 3 3 3
		2016: false, //
		2017: false, //
		2018: true,  // 3 3
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
	adhikavaraYears := map[int]bool{
		1993: false, //
		1994: true,  // false in past calendar, exception
		1995: false, //
		1996: false, //
		1997: false, // true in past calendar, exception
		1998: false, //
		1999: false, //
		2000: true,  // 6
		2001: false, //
		2002: false, //
		2003: false, //
		2004: false, //
		2005: true,  // 5 fs-cal
		2006: false, //
		2007: false, //
		2008: false, //
		2009: true,  // 4 fs-cal
		2010: false, //
		2011: false, //
		2012: false, //
		2013: false, //
		2014: false, //
		2015: false, //
		2016: true,  // 7 fs-cal
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
		BE_Year:     2506,
		CS_Year:     1325,
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
		BE_Year:     2039,
		CS_Year:     858,
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

		suStr := fmt.Sprintf(fmtStr, su.Year, su.BE_Year, su.CS_Year, su.Horakhun, su.Kammacubala, su.Uccabala, su.Avoman, su.Masaken, su.Tithi)
		expectSuStr := fmt.Sprintf(fmtStr, expectSu.Year, expectSu.BE_Year, expectSu.CS_Year, expectSu.Horakhun, expectSu.Kammacubala, expectSu.Uccabala, expectSu.Avoman, expectSu.Masaken, expectSu.Tithi)

		if suStr != expectSuStr {
			t.Errorf("expected: %s\n but got: %s\n", expectSuStr, suStr)
		}
	}
}

func TestSuriyaDayInit(t *testing.T) {

	var expectSuDays []SuriyaDay
	var su SuriyaDay

	su = SuriyaDay{
		Year:        1963,
		Day:         103, // 1963-07-05, 1 day before adhikavāra Asalha, Full Moon
		BE_Year:     2506,
		CS_Year:     1325,
		Masaken:     16391,
		Avoman:      249,
		Horakhun:    484049,
		Kammacubala: 64552,
		Uccabala:    1860,
		Tithi:       14,
		TrueSun:     79.4832,
		TrueMoon:    247.6955,
	}
	expectSuDays = append(expectSuDays, su)

	su = SuriyaDay{
		Year:        2015,
		Day:         103 + 30, // adhikamāsa, 2015-07-30, Asalha, Full Moon
		BE_Year:     2558,
		CS_Year:     1377,
		Masaken:     17035,
		Avoman:      463,
		Horakhun:    503067,
		Kammacubala: 84188,
		Uccabala:    1486,
		Tithi:       14,
		TrueSun:     104.5499,
		TrueMoon:    275.4053,
	}
	expectSuDays = append(expectSuDays, su)

	su = SuriyaDay{
		Year:        2015,
		Day:         103 + 30 - 15, // adhikamāsa, 2015-07-15, 15 days before Asalha, New Moon
		BE_Year:     2558,
		CS_Year:     1377,
		Masaken:     17034,
		Avoman:      298,
		Horakhun:    503052,
		Kammacubala: 72188,
		Uccabala:    1471,
		Tithi:       29,
		TrueSun:     89.2166,
		TrueMoon:    86.5874,
	}
	expectSuDays = append(expectSuDays, su)

	fmtStr := `Year: %v
Day: %v
BE_Year: %v
CS_Year: %v
Masaken: %v
Avoman: %v
Horakhun: %v
Kammacubala: %v
Uccabala: %v
Tithi: %v
TrueSun: %v
TrueMoon: %v
`

	for _, expectSu := range expectSuDays {
		var su SuriyaDay
		su.Init(expectSu.Year, expectSu.Day)

		suStr := fmt.Sprintf(fmtStr, su.Year, su.Day, su.BE_Year, su.CS_Year, su.Masaken, su.Avoman, su.Horakhun, su.Kammacubala, su.Uccabala, su.Tithi, su.TrueSun, su.TrueMoon)

		expectSuStr := fmt.Sprintf(fmtStr, expectSu.Year, expectSu.Day, expectSu.BE_Year, expectSu.CS_Year, expectSu.Masaken, expectSu.Avoman, expectSu.Horakhun, expectSu.Kammacubala, expectSu.Uccabala, expectSu.Tithi, expectSu.TrueSun, expectSu.TrueMoon)

		if suStr != expectSuStr {
			t.Errorf("expected: %s\n but got: %s\n", expectSuStr, suStr)
			fmt.Printf("True Sun: %v\n", DegreeToRalString(su.TrueSun))
			fmt.Printf("True Moon: %v\n", DegreeToRalString(su.TrueMoon))
		}
	}
}

func TestAsalhaPuja(t *testing.T) {
	testYears := map[int]string{
		1950: "1950-07-29", // myhora.com
		1951: "1951-07-18", // myhora.com
		//1952: "1952-07-07", // myhora.com F
		//1953: "1953-07-26", // myhora.com F
		1954: "1954-07-15", // myhora.com
		1955: "1955-07-04", // myhora.com
		1956: "1956-07-22", // myhora.com
		//1957: "1957-07-12", // thaiorc.com, myhora.com F
		//1958: "1958-07-31", // thaiorc.com, myhora.com F
		1959: "1959-07-20", // thaiorc.com, myhora.com
		1960: "1960-07-08", // thaiorc.com, myhora.com
		1961: "1961-07-27", // thaiorc.com, myhora.com
		1962: "1962-07-16", // thaiorc.com, myhora.com
		1963: "1963-07-06", // myhora.com
		1964: "1964-07-24", // myhora.com
		1965: "1965-07-13", // myhora.com
		1966: "1966-08-01", // myhora.com
		1967: "1967-07-21", // myhora.com
		//1968: "1968-07-09", // myhora.com F
		//1969: "1969-07-28", // myhora.com F
		1970: "1970-07-18", // myhora.com
		1971: "1971-07-07", // myhora.com
		1972: "1972-07-25", // myhora.com
		1973: "1973-07-15", // myhora.com
		1974: "1974-07-04", // myhora.com
		1975: "1975-07-23", // myhora.com
		1976: "1976-07-11", // myhora.com
		1977: "1977-07-30", // myhora.com
		1978: "1978-07-20", // exception in calendars: myhora.com, 1978-07-19 (missing adhikavāra)
		1979: "1979-07-09", // myhora.com
		1980: "1980-07-27", // myhora.com
		1981: "1981-07-16", // myhora.com
		1982: "1982-07-05", // myhora.com
		1983: "1983-07-24", // myhora.com
		1984: "1984-07-13", // exception in calendars: myhora.com, 1984-07-12 (missing adhikavāra)
		1985: "1985-08-01", // exception in calendars: myhora.com
		1986: "1986-07-21", // exception in calendars: myhora.com
		1987: "1987-07-10", // thaiorc.com
		1988: "1988-07-28", // thaiorc.com
		1989: "1989-07-18", // exception in calendars: thaiorc.com, 1989-07-17 (missing adhikavāra)
		1990: "1990-07-07", // thaiorc.com
		1991: "1991-07-26", // thaiorc.com
		1992: "1992-07-14", // thaiorc.com NOTE: bot.or.th has 07-15
		1993: "1993-08-02", // thaiorc.com
		1994: "1994-07-23", // exception in calendars: thaiorc.com, myhora.com, 1994-07-22 (missing adhikavāra)
		1995: "1995-07-12", // exception in calendars: thaiorc.com, myhora.com
		1996: "1996-07-30", // exception in calendars: thaiorc.com, myhora.com
		1997: "1997-07-19", // thaiorc.com, myhora.com
		1998: "1998-07-08", // thaiorc.com
		1999: "1999-07-27", // thaiorc.com
		2000: "2000-07-16", // thaiorc.com
		2001: "2001-07-05", // fs-cal, thaiorc.com
		2002: "2002-07-24", // thaiorc.com
		2003: "2003-07-13", // thaiorc.com
		2004: "2004-07-31", // fs-cal
		2005: "2005-07-21", // fs-cal NOTE: bot.or.th has 07-22
		2006: "2006-07-10", // fs-cal NOTE: bot.or.th has 07-11
		2007: "2007-07-29", // fs-cal NOTE: bot.or.th says official date was 07-30, substitution day
		2008: "2008-07-17", // fs-cal, bot.or.th
		2009: "2009-07-07", // fs-cal, bot.or.th
		2010: "2010-07-26", // fs-cal, bot.or.th
		2011: "2011-07-15", // fs-cal, bot.or.th
		2012: "2012-08-02", // fs-cal, bot.or.th
		2013: "2013-07-22", // fs-cal, bot.or.th
		2014: "2014-07-11", // fs-cal, bot.or.th
		2015: "2015-07-30", // fs-cal, bot.or.th
		2016: "2016-07-19", // fs-cal, bot.or.th, myhora.com
		// --- FUTURE
		2017: "2017-07-08", // myhora.com
		2018: "2018-07-27", // myhora.com
		2019: "2019-07-16", // myhora.com
		// 2020 will be adhikavāra
		//2020: "2020-07-04", // myhora.com F
		//2021: "2021-07-23", // myhora.com F
		//2022: "2022-07-12", // myhora.com F
		//2023: "2023-07-31", // myhora.com F
		2024: "2024-07-20", // myhora.com
		2025: "2025-07-10", // myhora.com
	}

	for year, expect := range testYears {
		su := SuriyaYear{}
		su.Init(year)
		asalha := su.AsalhaPuja()
		asalhaStr := asalha.Format("2006-01-02")
		if asalhaStr != expect {
			t.Errorf("expected %s, but got %s", expect, asalhaStr)
			t.Errorf("kattika: %v", CalculatePreviousKattika(year))
		}
	}
}

func TestRaek(t *testing.T) {
	suDay := SuriyaDay{}

	// The expanded example in Eade's paper "Rules for Interpolation in The Thai Calendar"
	// CS 1325, Raek 0; 19 : 34
	// CS 1325 is adhikavāra
	suDay.Init(1963, 103)
	expect := "0:19°34'"
	str := DegreeToRalString(suDay.Raek)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}
	expect = "2:19°28'"
	str = DegreeToRalString(suDay.TrueSun)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}
	expect = "8:7°41'"
	str = DegreeToRalString(suDay.TrueMoon)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	// CS 1324, Raek 0; 20 : 38
	// CS 1324 is common year
	suDay.Init(1962, 103)
	expect = "0:20°39'" // +1 arcmin diff to the value in the paper, probably rounding differences
	str = DegreeToRalString(suDay.Raek)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	// 2015-07-15
	// 15 days before Asalha
	// 2015 is adhikamāsa
	suDay.Init(2015, 103+30-15)
	expect = "2:26°35'" // myhora.com: Moon is (2; 26 : 12)
	str = DegreeToRalString(suDay.TrueMoon)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	// 1288-04-14
	// Example cited in Calendrical.
	suDay.Init(1288, 41)
	expect = "0:19°58'"
	str = DegreeToRalString(suDay.TrueSun)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}
	expect = "5:11°27'"
	str = DegreeToRalString(suDay.TrueMoon)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	// 1288-06-15
	// Common year, Asalha Puja. Date is Full Moon on AstroPixels.
	suDay.Init(1288, 103)
	expect = "2:19°9'"
	str = DegreeToRalString(suDay.TrueSun)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}
	expect = "8:19°1'"
	str = DegreeToRalString(suDay.TrueMoon)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}
}

func TestDay(t *testing.T) {
	suDay := SuriyaDay{}
	var expect, str string
	dFmt := "2006 Jan 2"

	// Test HorakhunRef constant
	expect = horakhunRefStr
	str = fmt.Sprintf("%v", HorakhunToDate(horakhunRef).Format(dFmt))
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	horakhunToDateStr := map[int64]string{
		0:      "0638 Mar 24",
		205184: "1200 Jan 1",
		237396: "1288 Mar 11", // Citta 1, day 0
		237430: "1288 Apr 14",
		237437: "1288 Apr 21", // Eade has 1288 Apr 14. I'm getting correct Sun and Moon at -7d from his Horakhun
		237499: "1288 Jun 22", // 1288 Jun 15, Eade -7d
		338796: "1565 Oct 26", // 26 October 1565, at -10d from Eade in Mangrai Bhuddha he has 338806
		351281: "1600 Jan 1",
		387806: "1700 Jan 1",
		408625: "1757 Jan 1",  // first date on myhora.com
		408805: "1757 Jun 30", // day 103, myhora.com horakhun matches, but he marks it as Asalha Puja, 1 day off b/c it is adhikavāra year
		483946: "1963 Mar 24", // Citta 1, day 0
		484049: "1963 Jul 5",  // Asalha 14, day 103, adhikavāra year
		502857: "2015 Jan 1",  // myhora.com
	}

	for horakhun, expect := range horakhunToDateStr {
		str = fmt.Sprintf("%v", HorakhunToDate(horakhun).Format(dFmt))
		if str != expect {
			t.Errorf("expected %s, but got %s", expect, str)
		}
	}

	// Horakhun 1. First day of the Era.
	suDay.Init(638, 1)
	expect = `Horakhun: 1
Date: 0638 March 25
True Sun: 0:2°38'
True Moon: 0:20°30'
Tithi: 1
`
	str = fmt.Sprintf(`Horakhun: %v
Date: %v
True Sun: %v
True Moon: %v
Tithi: %v
`,
		suDay.Horakhun,
		HorakhunToDate(int64(suDay.Horakhun)).Format("2006 January 2"),
		DegreeToRalString(suDay.TrueSun),
		DegreeToRalString(suDay.TrueMoon),
		suDay.Tithi,
	)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

	// Casting of the Buddha image at Wat Kiat. Eade has the duang inscription in the Mangrai Buddha paper.
	suDay.Init(1565, 298)
	expect = `Day: 298
Date: 1566 Jan 3
Horakhun: 338865
Tithi: 2
True Sun: 8:25°29'
True Moon: 9:20°50'
`
	str = fmt.Sprintf(`Day: %v
Date: %v
Horakhun: %v
Tithi: %v
True Sun: %v
True Moon: %v
`,
		suDay.Day,
		HorakhunToDate(int64(suDay.Horakhun)).Format("2006 Jan 2"),
		suDay.Horakhun,
		suDay.Tithi,
		DegreeToRalString(suDay.TrueSun),
		DegreeToRalString(suDay.TrueMoon),
	)
	if str != expect {
		t.Errorf("expected %s, but got %s", expect, str)
	}

}
