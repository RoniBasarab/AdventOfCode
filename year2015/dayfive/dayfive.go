package dayfive

import (
	"fmt"
	"strings"
	"utils"
)

func Solution() {
	data := utils.GetActualDataFromFileByLine("five", "2015")
	dataArray := strings.Split(data, "\n")
	niceStrings := 0

	for _, word := range dataArray {
		if betterIsNiceString(word) {
			niceStrings++
		}
	}

	fmt.Println("Number of total words: ", len(dataArray))
	fmt.Println("Nice strings:", niceStrings)

}

func isNiceString(word string) bool {

	atleastThreeVowels := false
	atleastOneLetterAppearsTwice := false

	vowels := "aeiou"
	forbiddenStrings := []string{"ab", "cd", "pq", "xy"}

	vowelCounter := 0

	for _, badWord := range forbiddenStrings {
		if strings.Contains(word, badWord) {
			fmt.Println("This word contained a forbbiden string: ", word, "-", badWord)
			return false
		}
	}

	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			vowelCounter++
			if vowelCounter >= 3 {
				atleastThreeVowels = true
				break
			}
		}
	}

	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			atleastOneLetterAppearsTwice = true
			break
		}
	}
	return atleastThreeVowels && atleastOneLetterAppearsTwice
}

func betterIsNiceString(word string) bool {
	hasPairOfTwoLetters := false
	hasLetterRepeating := false

	for i := 1; i < len(word)-2; i++ {
		pattern := word[i-1 : i+1]
		if strings.Contains(word[i+1:], pattern) {
			hasPairOfTwoLetters = true
			// fmt.Println("Found pattern that appears twice: (", pattern, ") in word: ", word)
			break
		}
	}

	if hasPairOfTwoLetters == false {
		fmt.Println("This word doesnt have the pattern! - ", word)
		return false
	}

	//      qjhvhtzxzqqjkmpb
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			// fmt.Printf("Found repeating letter: (%c%c%c) in word: %s\n", rune(word[i]), rune(word[i+1]), rune(word[i+2]), word)
			hasLetterRepeating = true
			break
		}
	}

	if hasLetterRepeating == false {
		fmt.Println("This word doesnt have the repeating letter! - ", word)
		return false
	}

	return hasPairOfTwoLetters && hasLetterRepeating
}
