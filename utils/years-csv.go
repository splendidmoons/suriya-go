package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/splendidmoons/suriya-go"
)

type testYear struct {
	Asalha   string
	Source   string
	Comments string
}

func yearsCsvSimple(firstYear int, lastYear int) (csvString string) {
	// Assume the 57 year cycle as 1984-2040.
	//
	// 1984 is adhikavāra (diff 6), 1985 is adhikamāsa (diff 3), and this will not
	// offset the familiar 332-3332 pattern that Aj Khemanando used in the old
	// circular diagram.

	csvString = "CE year,BE year,nM,dM,nV,dV,K,A,T\n"

	for year := firstYear; year <= lastYear; year++ {
		su := suriya.SuriyaYear{}
		su.Init(year)

		dM := ""
		if su.Is_Adhikamasa() {
			dM = strconv.Itoa(su.DeltaAdhikamasa())
		}
		dV := ""
		if su.Is_Adhikavara() {
			dV = strconv.Itoa(su.DeltaAdhikavara())
		}

		csvString = csvString + fmt.Sprintf("%d,%d,%d,%s,%d,%s,%s\n",
			su.Year,
			su.BE_year,
			su.AdhikamasaCyclePos(),
			dM,
			su.AdhikavaraCyclePos(),
			dV,
			su.Kammacubala,
			su.Avoman,
			su.Tithi,
		)
	}

	return csvString
}

func yearsCsv(firstYear int, lastYear int) (csvString string) {
	// Assume the 57 year cycle as 1984-2040.
	//
	// 1984 is adhikavāra (diff 6), 1985 is adhikamāsa (diff 3), and this will not
	// offset the familiar 332-3332 pattern that Aj Khemanando used in the old
	// circular diagram.

	testYears := map[int]testYear{
		1970: {"1970-07-18", "myhora.com", ""},
		1971: {"1971-07-07", "myhora.com", ""},
		1972: {"1972-07-25", "myhora.com", ""},
		1973: {"1973-07-15", "myhora.com", ""},
		1974: {"1974-07-04", "myhora.com", ""},
		1975: {"1975-07-23", "myhora.com", ""},
		1976: {"1976-07-11", "myhora.com", ""},
		1977: {"1977-07-30", "myhora.com", ""},
		1978: {"1978-07-19", "myhora.com", "adhikavāra is missing from the calendar"}, // FAIL
		1979: {"1979-07-09", "myhora.com", ""},
		1980: {"1980-07-27", "myhora.com", ""},
		1981: {"1981-07-16", "myhora.com", ""},
		1982: {"1982-07-05", "myhora.com", ""},
		1983: {"1983-07-24", "myhora.com", ""},
		1984: {"1984-07-12", "myhora.com", "adhikavāra is missing from the calendar"}, // FAIL
		1985: {"1985-07-31", "myhora.com", ""},                                        // FAIL
		1986: {"1986-07-20", "myhora.com", ""},                                        // FAIL
		1987: {"1987-07-10", "thaiorc.com", ""},
		1988: {"1988-07-28", "thaiorc.com", ""},
		1989: {"1989-07-17", "thaiorc.com", ""}, // FAIL
		1990: {"1990-07-07", "thaiorc.com", ""},
		1991: {"1991-07-26", "thaiorc.com", ""},
		1992: {"1992-07-14", "thaiorc.com", ""},
		1993: {"1993-08-02", "thaiorc.com", ""},
		1994: {"1994-07-22", "thaiorc.com, myhora.com", "calendar is missing adhikavāra, passing with exception"},
		1995: {"1995-07-11", "thaiorc.com, myhora.com", ""},
		1996: {"1996-07-29", "thaiorc.com, myhora.com", ""},
		1997: {"1997-07-19", "thaiorc.com, myhora.com", "missing adhikavāra was added back here, passing with exception"},
		1998: {"1998-07-08", "thaiorc.com", ""},
		1999: {"1999-07-27", "thaiorc.com", ""},
		2000: {"2000-07-16", "thaiorc.com", ""},
		2001: {"2001-07-05", "fs-cal, thaiorc.com", ""},
		2002: {"2002-07-24", "thaiorc.com", ""},
		2003: {"2003-07-13", "thaiorc.com", ""},
		2004: {"2004-07-31", "fs-cal", ""},
		2005: {"2005-07-21", "fs-cal", ""},
		2006: {"2006-07-10", "fs-cal", ""},
		2007: {"2007-07-29", "fs-cal", ""},
		2008: {"2008-07-17", "fs-cal, bot.or.th", ""},
		2009: {"2009-07-07", "fs-cal, bot.or.th", ""},
		2010: {"2010-07-26", "fs-cal, bot.or.th", ""},
		2011: {"2011-07-15", "fs-cal, bot.or.th", ""},
		2012: {"2012-08-02", "fs-cal, bot.or.th", ""},
		2013: {"2013-07-22", "fs-cal, bot.or.th", ""},
		2014: {"2014-07-11", "fs-cal, bot.or.th", ""},
		2015: {"2015-07-30", "fs-cal, bot.or.th", ""},
		2016: {"2016-07-19", "fs-cal, bot.or.th, myhora.com", ""},
	}

	csvString = "CE year;BE year;K;A;T;nM;dM;nV;dV;Asalha by Calc.;Asalha in Calendar;test;source;comments\n"

	for year := firstYear; year <= lastYear; year++ {
		su := suriya.SuriyaYear{}
		su.Init(year)

		dM := ""
		if su.Is_Adhikamasa() {
			dM = strconv.Itoa(su.DeltaAdhikamasa())
		}
		dV := ""
		if su.Is_Adhikavara() {
			dV = strconv.Itoa(su.DeltaAdhikavara())
		} else if su.Would_Be_Adhikavara() {
			dV = "x"
		}

		fmtStr := fmt.Sprint(
			"%v;",  // CE year
			"%v;",  // BE year
			"%v;",  // K
			"%v;",  // A
			"%v;",  // T
			"%v;",  // nM
			"%v;",  // dM
			"%v;",  // nV
			"%v;",  // dV
			"%v;",  // Asalha by Calc
			"%v;",  // Asalha in Calendar
			"%v;",  // test
			"%v;",  // source
			"%v\n", // comments
		)

		asalhaStr := su.AsalhaPuja().Format("2006-01-02")
		testRes := ""

		var tYear testYear
		tYear, ok := testYears[year]
		if ok {
			if asalhaStr == tYear.Asalha {
				testRes = "OK"
			} else {
				testRes = "FAIL"
			}
		}

		csvString = csvString + fmt.Sprintf(fmtStr,
			su.Year,
			su.BE_year,
			su.Kammacubala,
			su.Avoman,
			su.Tithi,
			su.AdhikamasaCyclePos(),
			dM,
			su.AdhikavaraCyclePos(),
			dV,
			asalhaStr,
			tYear.Asalha,
			testRes,
			tYear.Source,
			tYear.Comments,
		)
	}

	return csvString
}

func main() {
	file, err := os.Create("years.csv")
	defer file.Close()

	if err != nil {
		panic(err)
		return
	}

	n, err := io.WriteString(file, yearsCsv(1970, 2041))
	if err != nil {
		fmt.Println(n, err)
	}

}
