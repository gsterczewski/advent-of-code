package solutions

import (
	"bufio"
	"regexp"
	"strconv"
)

func SolutionDay1A(is *bufio.Scanner) int {
	sum := 0
	regex := regexp.MustCompile(`\d`)
	for is.Scan() {
		matches := regex.FindAllString(is.Text(), -1)
		if len(matches) > 0 {
			firstMatch := matches[0]
			lastMatch := matches[len(matches)-1]
			num, _ := strconv.Atoi(firstMatch + lastMatch)
			sum += num
		}

	}
	return sum
}

func SolutionDay1B(is *bufio.Scanner) int {
	return 0

}
