package solutions

import (
	"advent-of-code/utils"
	"testing"
)

func TestSolutionDay2A(t *testing.T) {
	testInput := utils.GetInput("day2-test-input.txt")
	result := SolutionDay2A(testInput)
	expected := 8
	if result != expected {
		t.Errorf("Day 1 Solution B should return %d, got %d", expected, result)
	}
}
func TestSolutionDay2B(t *testing.T) {
	testInput := utils.GetInput("day2-test-input.txt")
	result := SolutionDay2B(testInput)
	expected := 2286
	if result != expected {
		t.Errorf("Day 1 Solution B should return %d, got %d", expected, result)
	}
}
