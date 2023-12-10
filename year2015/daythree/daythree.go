package daythree

import (
	"fmt"
	"utils"
	"year2015/datastruct"
)

func Solution() {

	/* PART 1 */
	/*
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
	*/

	/* PART 2 */
	data := utils.GetActualDataFromFile("three", "2015")
	totalUniqueHousesVisitedBySanta := 1
	totalUniqueHousesVisitedByRobot := 0
	lastStepSanta := datastruct.Coordinates2D{X: 0, Y: 0}
	lastStepRobot := datastruct.Coordinates2D{X: 0, Y: 0}
	santasPath := datastruct.Set{}
	robotPath := datastruct.Set{}
	santasPath[lastStepSanta] = lastStepSanta
	robotPath[lastStepRobot] = lastStepRobot
	north := datastruct.Coordinates2D{X: 0, Y: 1}
	south := datastruct.Coordinates2D{X: 0, Y: -1}
	west := datastruct.Coordinates2D{X: -1, Y: 0}
	east := datastruct.Coordinates2D{X: 1, Y: 0}

	for i, direction := range data {
		if i%2 == 0 {
			calculateStep(direction, &lastStepRobot, &north, &south, &east, &west)
			lastStepRobot = santasPath.AddAndReturnCurrentSantaPosition(lastStepRobot, &totalUniqueHousesVisitedByRobot)
		} else {
			calculateStep(direction, &lastStepSanta, &north, &south, &east, &west)
			lastStepSanta = robotPath.AddAndReturnCurrentSantaPosition(lastStepSanta, &totalUniqueHousesVisitedBySanta)
		}
	}

	uniquePath := santasPath.SymmetricDifference(&robotPath)

	fmt.Println(data)
	fmt.Println(uniquePath)
	fmt.Println(uniquePath.Size())

}

func calculateStep(direction rune, step, north, south, east, west *datastruct.Coordinates2D) {
	switch direction {
	case '^':
		step.Plus(north)
	case 'v':
		step.Plus(south)
	case '<':
		step.Plus(west)
	case '>':
		step.Plus(east)
	}
}
