package main

import (
	"fmt"
	"year2015"
)

func main() {
	fmt.Println("Please specify which day you wish to see the solution of (Example: one,two,three... (maximum day: eleven)):")
	var day string
	_, err := fmt.Scan(&day)
	if err != nil {
		fmt.Println("Error in input ", err)
	}
	year2015.LaunchByDay(day)
}
