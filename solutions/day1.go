package solutions

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

func extractNumFromString(str string) int {
	numStr := ""
	for _, char := range str {
		s := string(char)
		_, err := strconv.Atoi(s)
		if err == nil {
			numStr += string(char)
		}
	}
	result, err := strconv.Atoi(numStr[0:1] + numStr[len(numStr)-1:])
	if err != nil {
		fmt.Println(err)
	}
	return result
}
func SolutionDay1() int {
	inputScanner := utils.GetInput("inputs/day1-input.txt")
	sum := 0
	for inputScanner.Scan() {
		sum += extractNumFromString(inputScanner.Text())
	}
	return sum
}
