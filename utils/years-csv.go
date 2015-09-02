package main

import (
	"fmt"
	"strconv"

	"github.com/splendidmoons/suriya-go"
)

func yearsCsv(firstYear int, lastYear int) (csvString string) {
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
			su.KatString(),
		)
	}

	return csvString
}

func main() {
	fmt.Print(yearsCsv(1984, 2040))
}
