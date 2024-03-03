package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

func GetInput(inputPath string) *bufio.Scanner {
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
