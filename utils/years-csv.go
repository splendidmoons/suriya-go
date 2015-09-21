package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/splendidmoons/suriya-go"
)

type testYear struct {
	Asalha string
	Source string
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
		1992: {"1992-07-15", "bot.or.th"},         // FAIL GOT: 1992-07-14
		2001: {"2001-07-05", "fs-cal"},            // FAIL GOT: 2001-08-04
		2004: {"2004-07-31", "fs-cal"},            // PASS
		2005: {"2005-07-22", "bot.or.th"},         // PASS
		2006: {"2006-07-10", "fs-cal"},            // PASS
		2007: {"2007-07-29", "fs-cal"},            // PASS
		2008: {"2008-07-17", "bot.or.th, fs-cal"}, // PASS
		2009: {"2009-07-07", "bot.or.th, fs-cal"}, // FAIL GOT: 2009-08-05
		2010: {"2010-07-26", "bot.or.th, fs-cal"}, // FAIL GOT: 2010-07-27
		2011: {"2011-07-15", "bot.or.th, fs-cal"}, // PASS
		2012: {"2012-08-02", "bot.or.th, fs-cal"}, // PASS
		2013: {"2013-07-22", "bot.or.th, fs-cal"}, // PASS
		2014: {"2014-07-11", "bot.or.th, fs-cal"}, // PASS
		2015: {"2015-07-30", "bot.or.th, fs-cal"}, // PASS
		2016: {"2016-07-19", "bot.or.th, fs-cal"}, // PASS
	}

	csvString = "CE year;BE year;K;A;T;nM;dM;nV;dV;Asalha by Calc;Asalha in Calendar;source;check;comments\n"

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
			"%v;",  // source
			"%v;",  // check
			"%v\n", // comments
		)

		asalhaStr := su.AsalhaPuja().Format("2006-01-02")
		pass := ""

		var tYear testYear
		tYear, ok := testYears[year]
		if ok {
			if asalhaStr == tYear.Asalha {
				pass = "OK"
			} else {
				pass = "FAIL"
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
			tYear.Source,
			pass,
			"",
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

	n, err := io.WriteString(file, yearsCsv(1957, 2040))
	if err != nil {
		fmt.Println(n, err)
	}

}
