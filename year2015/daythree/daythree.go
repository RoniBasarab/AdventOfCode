package daythree

import (
	"fmt"
	"utils"
	"year2015/datastruct"
)

func Solution() {

	/* PART 1 */
	data := utils.GetActualDataFromFile("three", "2015")
	totalUniqueHousesVisitedBySanta := 1
	lastStep := datastruct.Coordinates2D{X: 0, Y: 0}
	santasPath := datastruct.Set{Items: []interface{}{lastStep}}
	north := datastruct.Coordinates2D{X: 0, Y: 1}
	south := datastruct.Coordinates2D{X: 0, Y: -1}
	west := datastruct.Coordinates2D{X: -1, Y: 0}
	east := datastruct.Coordinates2D{X: 1, Y: 0}

	for _, direction := range data {
		switch direction {
		case '^':
			lastStep.Plus(&north)
		case 'v':
			lastStep.Plus(&south)
		case '<':
			lastStep.Plus(&west)
		case '>':
			lastStep.Plus(&east)
		}
		lastStep = santasPath.AddAndReturnCurrentSantaPosition(lastStep, &totalUniqueHousesVisitedBySanta)
	}
	fmt.Println(data)
	fmt.Println(santasPath)
	fmt.Println(totalUniqueHousesVisitedBySanta)

	/* PART 2 */

}
