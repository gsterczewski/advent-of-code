package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func GetInput(inputPath string) *bufio.Scanner {

	wd := "C:/Users/GS/Programowanie/Advent-Of-Code/inputs"
	fullPath := path.Join(wd, inputPath)
	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Println("invalid input file path")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
