package main

import (
	"2015/datastruct"
	"fmt"
)

func main() {

	//data := getActualDataFromFile("daytwoinput.txt")
	// data := getActualDataFromFileByLine("daytwoinput.txt")
	// data := getActualDataFromFile("daythreeinput.txt")
	data := "^>v<"
	totalUniqueHousesVisitedBySanta := 1
	santasPath := &datastruct.FourWayNode{Up: nil, Down: nil, Left: nil, Right: nil}

	for _, char := range data {
		santasPath = santasPath.AppendAndReturnCurrentPosition(char, &totalUniqueHousesVisitedBySanta)
	}

	fmt.Println(data)
	fmt.Println(totalUniqueHousesVisitedBySanta)
}
