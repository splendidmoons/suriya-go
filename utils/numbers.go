package main

import (
	"fmt"

	"github.com/splendidmoons/suriya-go"
)

func main() {
	//suYear := suriya.SuriyaYear{}

	//suYear.Init(1962) // CS 1324
	//suYear.Init(1963) // CS 1325
	//suYear.Init(1964) // CS 1326

	suDay := suriya.SuriyaDay{}
	suDay.Init(1963, 103)

	//suDay.Init(1963, 1)

	fmt.Printf("%s\n", suriya.DegreeToEadeString(suDay.MeanMoon))
	fmt.Printf("%s\n", suriya.DegreeToEadeString(suDay.TrueMoon))

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
