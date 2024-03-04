package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

type gameRound struct {
	red   int
	green int
	blue  int
}
type Game struct {
	id     int
	rounds []gameRound
}

func createGameRound(str string) gameRound {
	arr := strings.Split(str, ",")
	redCount := 0
	blueCount := 0
	greenCount := 0
	for _, sentence := range arr {
		sentence = strings.Trim(sentence, " ")
		words := strings.Split(sentence, " ")
		color := words[1]
		count, _ := strconv.Atoi(words[0])
		switch color {
		case "red":
			redCount += count

		case "blue":
			blueCount += count

		case "green":
			greenCount += count

		}

	}
	gr := gameRound{
		red:   redCount,
		blue:  blueCount,
		green: greenCount,
	}

	return gr
}
func createGame(str string) Game {
	strArr := strings.Split(str, ":")
	id, _ := strconv.Atoi(strings.Split(strArr[0], " ")[1])
	rounds := strings.Split(strings.Join(strArr[1:], ""), ";")
	game := Game{id: id, rounds: make([]gameRound, 0)}
	for _, r := range rounds {
		round := createGameRound(r)
		game.rounds = append(game.rounds, round)

	}

	return game
}

func SolutionDay2A(inputScanner *bufio.Scanner) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	result := 0
	for inputScanner.Scan() {
		game := createGame(inputScanner.Text())
		possible := true
		for _, round := range game.rounds {
			if (round.red > maxRed) || (round.blue > maxBlue) || (round.green > maxGreen) {
				possible = false
			}
		}
		if possible {
			result += game.id
		}
	}
	return result
}
func SolutionDay2B(inputScanner *bufio.Scanner) int {
	result := 0
	for inputScanner.Scan() {
		game := createGame(inputScanner.Text())
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, round := range game.rounds {
			if round.red > maxRed {
				maxRed = round.red
			}
			if round.blue > maxBlue {
				maxBlue = round.blue
			}
			if round.green > maxGreen {
				maxGreen = round.green
			}
		}
		result += maxRed * maxBlue * maxGreen
	}
	return result
}
