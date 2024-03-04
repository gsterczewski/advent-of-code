package solutions

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type scratchCard struct {
	winningNumbers []int
	playerNumbers  []int
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
func createScratchCard(str string) scratchCard {
	arr := strings.Split(str, ":")
	values := strings.Split(arr[1], "|")

	winningNumbers := mapStringsToNumbers(strings.Split(strings.Trim(values[0], " "), " "))
	playerNumbers := mapStringsToNumbers(strings.Split(strings.Trim(values[1], " "), " "))
	return scratchCard{winningNumbers, playerNumbers}

}
func (sc scratchCard) calculateScore() int {
	score := 0
	winningMap := map[int]bool{}
	for _, v := range sc.winningNumbers {
		winningMap[v] = false
	}
	for _, n := range sc.playerNumbers {
		used, exist := winningMap[n]
		if exist && !used {
			if score == 0 {
				score += 1
			} else {
				score *= 2
			}
		}
		winningMap[n] = true
	}

	return score
}
func SolutionDay4A(inputScanner *bufio.Scanner) int {
	result := 0
	for inputScanner.Scan() {
		sc := createScratchCard(inputScanner.Text())
		fmt.Printf("winning numbers: %v ", sc.winningNumbers)
		fmt.Printf("player numbers: %v ", sc.playerNumbers)
		fmt.Printf("score: %v\n", sc.calculateScore())

		result += sc.calculateScore()
	}
	return result
}
