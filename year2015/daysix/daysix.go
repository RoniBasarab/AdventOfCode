package daysix

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
	"year2015/datastruct"
)

const (
	toggle  = "toggle"
	turnOn  = "turn on"
	turnOff = "turn off"
)

func Solution() {
	data := utils.GetActualDataFromFileByLine("six", "2015")
	grid := initGrid(1000, 1000)

	for _, input := range data {

		var command string
		switch {
		case strings.Contains(input, toggle):
			command = toggle
		case strings.Contains(input, turnOn):
			command = turnOn
		case strings.Contains(input, turnOff):
			command = turnOff
		}

		point1, point2 := extractCoordinatesFromLine(input)
		adjustLights(point1.X, point2.X, point1.Y, point2.Y, &grid, command)
	}
	fmt.Println(grid)
	fmt.Println(countLights(&grid))

}

func initGrid(rows, cols int) [][]int {

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func extractCoordinatesFromLine(input string) (datastruct.Coordinates2D, datastruct.Coordinates2D) {

	re := regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)`)

	matches := re.FindStringSubmatch(input)

	x1, _ := strconv.Atoi(matches[1])
	y1, _ := strconv.Atoi(matches[2])
	x2, _ := strconv.Atoi(matches[3])
	y2, _ := strconv.Atoi(matches[4])

	return datastruct.Coordinates2D{X: x1, Y: y1}, datastruct.Coordinates2D{X: x2, Y: y2}
}

func adjustLights(fromRow, toRow, fromColumn, toColumn int, grid *[][]int, command string) {

	for i := fromRow; i <= toRow; i++ {
		for j := fromColumn; j <= toColumn; j++ {
			switch command {
			case toggle:
				(*grid)[i][j] += 2
			case turnOn:
				(*grid)[i][j] += 1
			case turnOff:
				if (*grid)[i][j] > 0 {
					(*grid)[i][j] -= 1
				}
			}
		}
	}
}

func countLights(grid *[][]int) int {
	totalBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += (*grid)[i][j]
		}
	}
	return totalBrightness
}
