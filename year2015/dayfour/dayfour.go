package dayfour

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"utils"
)

func Solution() {

	/* PART ONE */
	// data := utils.GetActualDataFromFile("four", "2015")
	// lowestNumberToHavePrefix := 1
	// for true {
	// 	dataToCheck := data + strconv.Itoa(lowestNumberToHavePrefix)
	// 	hashedString := MD5(dataToCheck)
	// 	if hasPrefix(hashedString) {
	// 		break
	// 	}
	// 	lowestNumberToHavePrefix++
	// }
	// fmt.Println(lowestNumberToHavePrefix)

	/* PART TWO */

	data := utils.GetActualDataFromFile("four", "2015")
	lowestNumberToHavePrefix := 1
	for true {
		dataToCheck := data + strconv.Itoa(lowestNumberToHavePrefix)
		hashedString := MD5(dataToCheck)
		if hasPrefix(hashedString) {
			break
		}
		lowestNumberToHavePrefix++
	}
	fmt.Println(lowestNumberToHavePrefix)

}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func hasPrefix(hashedString string) bool {

	if hashedString[:6] == "000000" {
		return true
	}
	return false
}
