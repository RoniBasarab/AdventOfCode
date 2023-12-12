package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func GetDataFile(specificDay, yearNumber string) *os.File {

	workingDir, errWd := os.Getwd()

	if errWd != nil {
		fmt.Println("Error getting working directory", errWd)
		return nil
	}

	viper.SetConfigFile(workingDir + "\\utils\\config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return nil
	}

	workingDirectory := viper.GetString("BASE_WORKING_DIRECTORY") + "\\"

	var sb strings.Builder
	sb.WriteString(workingDirectory)
	sb.WriteString("year" + yearNumber + "\\" + "day" + specificDay + "\\" + "day" + specificDay + "input.txt")

	input, errInput := os.Open(sb.String())
	if errInput != nil {
		fmt.Println("We encountered an error opening the data file", errInput)
		return nil
	}

	return input
}

func GetActualDataFromFile(specificDay, year string) string {
	inputFile := GetDataFile(specificDay, year)
	if inputFile == nil {
		return ""
	}

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var data strings.Builder

	for scanner.Scan() {
		data.WriteString(scanner.Text())
	}

	if scannerError(&scanner) {
		return ""
	}

	return data.String()
}

func GetActualDataFromFileByLine(specificDay, year string) []string {
	inputFile := GetDataFile(specificDay, year)
	if inputFile == nil {
		return nil
	}

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var data strings.Builder

	for scanner.Scan() {
		data.WriteString(scanner.Text())
		data.WriteString("\n")
	}

	if scannerError(&scanner) {
		return nil
	}
	dataArray := strings.Split(data.String(), "\n")
	return dataArray[:len(dataArray)-1]
}

func scannerError(scanner **bufio.Scanner) bool {
	if errScanner := (*scanner).Err(); errScanner != nil {
		fmt.Println("Error reading file:", errScanner)
		return true
	}
	return false
}
