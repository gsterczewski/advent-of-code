package solutions

import (
	"advent-of-code/utils"
	"bufio"
)

func getNextNumbers(current, n int) []int {
	results := []int{}
	for i := 1; i <= n; i++ {
		results = append(results, current+i)
	}
	return results
}

// TODO: OPTIMISE IT TO RUN FASTER !!
func SolutionDay4B(inputScanner *bufio.Scanner) int {
	totalCards := 0
	cardsStack := utils.CreateStack[ScratchCard]()
	cardsArray := []ScratchCard{}
	for inputScanner.Scan() {
		card := CreateScratchCard(inputScanner.Text())
		cardsArray = append(cardsArray, card)
		cardsStack.Push(card)
	}
	for !cardsStack.IsEmpty() {
		card := cardsStack.Pop()
		totalCards++

		hits := card.CalculateHits()
		for _, bonusCardNumber := range getNextNumbers(card.cardNumber, hits) {
			cardsStack.Push(cardsArray[bonusCardNumber-1])
		}
	}
	return totalCards
}
