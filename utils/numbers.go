package main

import (
	"fmt"

	"github.com/splendidmoons/suriya-go"
)

func main() {
	suYear := suriya.SuriyaYear{}
	//suYear.Init(1963) // CS 1325
	//suYear.Init(1496) // CS 858

	suYear.Init(2015)

	fmt.Print(suYear.SuriyaValuesString())

	fmt.Println("---")

	suNext := suriya.SuriyaYear{}
	suNext.Init(2016)

	fmt.Print(suNext.SuriyaValuesString())

	fmt.Printf("%d\n", suNext.Horakhun-suYear.Horakhun)
}
