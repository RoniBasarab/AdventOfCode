package daythree

import (
	"fmt"
	"utils"
	"year2015/datastruct"
)

func Solution() {

	data := utils.GetActualDataFromFile("three", "2015")
	totalUniqueHousesVisitedBySanta := 1
	santasPath := &datastruct.FourWayNode{Up: nil, Down: nil, Left: nil, Right: nil}

	for _, char := range data {
		santasPath = santasPath.AppendAndReturnCurrentPosition(char, &totalUniqueHousesVisitedBySanta)
	}

	fmt.Println(data)
	fmt.Println(totalUniqueHousesVisitedBySanta)
}
