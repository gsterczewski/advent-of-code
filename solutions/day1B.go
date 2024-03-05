package solutions

import (
	"bufio"
	"math"
	"regexp"
	"strconv"
)

func getMatchPosition(str string) ([2]int, bool) {

	regex := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	indexes := regex.FindIndex([]byte(str))
	if indexes != nil {
		return [2]int(indexes), true
	}
	return [2]int{-1, -1}, false
}
func findOverlappingSubstrings(str string) []string {

	result := []string{}
	index := 0
	for index < len(str) {
		currentStr := str[index:]

		positions, found := getMatchPosition(currentStr)
		if found {
			foundStr := currentStr[positions[0]:positions[1]]

			result = append(result, foundStr)
			if positions[1]-positions[0] > 1 {
				index += int(math.Max(float64(positions[1])-1, 1))
			} else {
				index += positions[1]
			}
		} else {
			index++
		}
	}
	return result
}
func SolutionDay1B(is *bufio.Scanner) int {
	result := 0
	digitsMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	for is.Scan() {
		matches := findOverlappingSubstrings(is.Text())
		first := digitsMap[matches[0]]
		last := digitsMap[matches[len(matches)-1]]

		sum, err := strconv.Atoi(first + last)

		if err == nil {
			result += sum
		}
	}
	return result

}
