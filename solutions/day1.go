package solutions

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func getInput(inputPath string) *bufio.Scanner {
	wd, _ := os.Getwd()
	fullPath := path.Join(wd, inputPath)
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("invalid input file path")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

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
	inputScanner := getInput("inputs/day1-input.txt")
	sum := 0
	for inputScanner.Scan() {
		sum += extractNumFromString(inputScanner.Text())
	}
	return sum
}
