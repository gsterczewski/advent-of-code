package solutions

import (
	"advent-of-code/utils"
	"bufio"
	"testing"
)

func TestSolutionDay1B(t *testing.T) {
	testCases := []*bufio.Scanner{utils.GetInput("day1-test-inputB.txt"), utils.GetInput("day1B-test1.txt")}
	expectedResults := []int{281, 396}
	for i, tc := range testCases {
		result := SolutionDay1B(tc)
		expected := expectedResults[i]
		if result != expected {
			t.Errorf("Day 1 Solution B should return %d, got %d", expected, result)
		}
	}
}
