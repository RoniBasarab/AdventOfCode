package year2015

import (
	"fmt"
	"year2015/dayfive"
	"year2015/dayfour"
	"year2015/daysix"
	"year2015/daythree"
)

const (
	one    = "one"
	two    = "two"
	three  = "three"
	four   = "four"
	five   = "five"
	six    = "six"
	seven  = "seven"
	eight  = "eight"
	nine   = "nine"
	ten    = "ten"
	eleven = "eleven"
)

func LaunchByDay(day string) {
	switch day {
	case three:
		{
			daythree.Solution()
		}
	case four:
		{
			dayfour.Solution()
		}
	case five:
		{
			dayfive.Solution()
		}
	case six:
		{
			daysix.Solution()
		}
	case seven:
	case eight:
	case nine:
	case ten:
	case eleven:
	default:
		{
			fmt.Println("There is no solution for this day! Please relaunch the application")
		}
	}
}
