package main

import (
	"advent-of-code/solutions"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	fmt.Printf("day 1 solution A : %d\n", solutions.SolutionDay1A(utils.GetInput("day1-input.txt")))
	// fmt.Printf("day 1 solution B : %d\n", solutions.SolutionDay1B(utils.GetInput("day1-input.txt")))
	fmt.Printf("day 2 solution A : %d\n", solutions.SolutionDay2A(utils.GetInput("day2-input.txt")))
	fmt.Printf("day 2 solution B : %d\n", solutions.SolutionDay2B(utils.GetInput("day2-input.txt")))
}
