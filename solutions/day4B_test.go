package solutions

import (
	"advent-of-code/utils"
	"bufio"
	"testing"
)

func TestSolutionDay4B(t *testing.T) {
	testCases := []*bufio.Scanner{utils.GetInput("day4-test-input.txt")}
	expectedResults := []int{30}

	for i, tc := range testCases {
		result := SolutionDay4B(tc)
		expected := expectedResults[i]

		if result != expected {
			t.Errorf("SolutionDay4B: expected result %d, got %d", expected, result)
		}
	}
}
