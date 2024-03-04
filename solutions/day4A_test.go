package solutions

import (
	"advent-of-code/utils"
	"bufio"
	"testing"
)

func TestSolutionDay4A(t *testing.T) {
	testCases := []*bufio.Scanner{utils.GetInput("day4-test-input.txt")}
	expectedResults := []int{13}

	for i, tc := range testCases {
		result := SolutionDay4A(tc)
		expected := expectedResults[i]

		if result != expected {
			t.Errorf("SolutionDay4A: expected result %d, got %d", expected, result)
		}
	}
}
