package main

import (
	"fmt"

	"github.com/splendidmoons/suriya-go"
)

func main() {
	suYear := suriya.SuriyaYear{}
	//suYear.Init(1963) // CS 1325
	suYear.Init(1496) // CS 858

	fmt.Print(suYear.SuriyaValuesString())
}
