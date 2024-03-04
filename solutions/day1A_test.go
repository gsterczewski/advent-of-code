package solutions

import (
	"advent-of-code/utils"
	"bufio"
	"testing"
)

func TestSolutionDay1A(t *testing.T) {
	testCases := []*bufio.Scanner{utils.GetInput("day1-test-inputA.txt")}
	expectedResults := []int{142}
	for i, tc := range testCases {
		result := SolutionDay1A(tc)
		expected := expectedResults[i]
		if result != expected {
			t.Errorf("Day 1 Solution A should return %d, got %d", expected, result)
		}
	}
}
