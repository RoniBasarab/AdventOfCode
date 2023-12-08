package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func getDataFile() *os.File {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return nil
	}

	var userDirectoryInput string
	fmt.Print("Please enter the year of the input file and the name of the input file. Example:\n2015\\daythree\\daythreeinput.txt")
	_, err := fmt.Scanln(&userDirectoryInput)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil
	}

	workingDirectory := viper.GetString("BASE_WORKING_DIRECTORY") + "\\"

	var sb strings.Builder
	sb.WriteString(workingDirectory)
	sb.WriteString(userDirectoryInput)

	input, errInput := os.Open(sb.String())
	if errInput != nil {
		fmt.Println("We encountered an error opening the data file", errInput)
		return nil
	}

	return input
}

func getActualDataFromFile() string {
	inputFile := getDataFile()
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

func getActualDataFromFileByLine(fileName string, dir string) string {
	inputFile := getDataFile()
	if inputFile == nil {
		return ""
	}

	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var data strings.Builder

	for scanner.Scan() {
		data.WriteString(scanner.Text())
		data.WriteString("\n")
	}

	if scannerError(&scanner) {
		return ""
	}

	return data.String()
}

func scannerError(scanner **bufio.Scanner) bool {
	if errScanner := (*scanner).Err(); errScanner != nil {
		fmt.Println("Error reading file:", errScanner)
		return true
	}
	return false
}
