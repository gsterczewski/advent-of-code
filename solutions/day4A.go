package solutions

import (
	"bufio"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	winningNumbers []int
	playerNumbers  []int
	cardNumber     int
}

func mapStringsToNumbers(strs []string) []int {
	result := []int{}
	for _, v := range strs {
		num, _ := strconv.Atoi(string(v))
		// HACK: get rid of empty space that gets converted to 0
		if num != 0 {
			result = append(result, num)
		}
	}
	return result
}
func CreateScratchCard(str string) ScratchCard {
	arr := strings.Split(str, ":")
	regex := regexp.MustCompile(`\d`)
	matched := regex.FindAllString(arr[0], -1)
	cardNumber, _ := strconv.Atoi(strings.Join(matched, ""))
	values := strings.Split(arr[1], "|")
	winningNumbers := mapStringsToNumbers(strings.Split(strings.Trim(values[0], " "), " "))
	playerNumbers := mapStringsToNumbers(strings.Split(strings.Trim(values[1], " "), " "))
	return ScratchCard{winningNumbers, playerNumbers, cardNumber}

}
func (sc ScratchCard) CalculateHits() int {
	hits := 0
	winningMap := map[int]bool{}
	for _, v := range sc.winningNumbers {
		winningMap[v] = false
	}
	for _, n := range sc.playerNumbers {
		used, exist := winningMap[n]
		if exist && !used {
			hits++
		}
		winningMap[n] = true
	}
	return hits
}
func (sc ScratchCard) calculateScore() int {

	hits := float64(sc.CalculateHits())
	if hits == 0 {
		return 0
	}
	return int(math.Pow(2, hits-1))
}

func SolutionDay4A(inputScanner *bufio.Scanner) int {
	result := 0
	for inputScanner.Scan() {
		sc := CreateScratchCard(inputScanner.Text())
		result += sc.calculateScore()
	}
	return result
}
