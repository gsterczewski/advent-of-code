package solutions

import (
	"advent-of-code/utils"
	"testing"
)

func TestSolutionDay1A(t *testing.T) {
	testInput := utils.GetInput("day1-test-inputA.txt")
	result := SolutionDay1A(testInput)
	expected := 142
	if result != expected {
		t.Errorf("Day 1 Solution B should return %d, got %d", expected, result)
	}
}
func TestSolutionDay1B(t *testing.T) {
	testInput := utils.GetInput("day1-test-inputB.txt")
	result := SolutionDay1B(testInput)
	expected := 281
	if result != expected {
		t.Errorf("Day 1 Solution B should return %d, got %d", expected, result)
	}
}
