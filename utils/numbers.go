package main

import (
	"fmt"

	"github.com/splendidmoons/suriya-go"
)

func main() {
	//suYear := suriya.SuriyaYear{}
	//suYear.Init(1963) // CS 1325

	suDay := suriya.SuriyaDay{}

	//suDay.Init(1958, 103+30) // CS 1320, Raek 22 : 28 (+4), adhikamāsa
	//suDay.Init(1959, 103) // CS 1321, Raek
	//suDay.Init(1960, 103+10) // CS 1322, Raek 20 : 41 (+1)
	//suDay.Init(1961, 103+10) // CS 1323, Raek 19 : 13 (+1), adhikamāsa
	//suDay.Init(1962, 103) // CS 1324, Raek 20 : 39 (+1)

	suDay.Init(1963, 103) // CS 1325, Raek 19 : 34

	//fmt.Printf("Mean Sun: %s\n", suriya.DegreeToEadeString(suDay.MeanSun))
	//fmt.Printf("Mean Moon: %s\n", suriya.DegreeToEadeString(suDay.MeanMoon))
	//fmt.Printf("True Moon: %s\n", suriya.DegreeToEadeString(suDay.TrueMoon))
	fmt.Printf("Raek: %s\n", suriya.DegreeToEadeString(suDay.Raek))
	fmt.Printf("Raek: %v\n", suDay.Raek)

	//suYear.Init(1496) // CS 858

	//suYear.Init(2016)

	//fmt.Print(suYear.SuriyaValuesString())
	//fmt.Printf("%v\n", suYear.AsalhaPuja())

	/*
		fmt.Println("---")

		suNext := suriya.SuriyaYear{}
		suNext.Init(2016)

		fmt.Print(suNext.SuriyaValuesString())

		fmt.Printf("%d\n", suNext.Horakhun-suYear.Horakhun)
	*/
}
