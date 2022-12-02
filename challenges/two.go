package challenges

import "strings"

type Game struct {
	firstSore   int
	secondScore int
}

type Score struct {
	winningHand int
	losingHand  int
	drawingHand int
}

type RockPaperScissorGame interface {
	result() int
}

const PAPER = 1
const SCISSORS = 2
const ROCK = 3

func (game Game) result() int {
	scoreTable := make(map[int]Score)
	scoreTable[PAPER] = Score{
		winningHand: SCISSORS,
		losingHand:  ROCK,
		drawingHand: PAPER,
	}
	scoreTable[SCISSORS] = Score{
		winningHand: ROCK,
		losingHand:  PAPER,
		drawingHand: SCISSORS,
	}
	scoreTable[ROCK] = Score{
		winningHand: PAPER,
		losingHand:  SCISSORS,
		drawingHand: ROCK,
	}

	currentScore := scoreTable[game.firstSore]
	ownHandScore := game.getOwnHandScore(game.secondScore)
	resultScore := game.getResultScore(game.secondScore, currentScore)
	return ownHandScore + resultScore
}

func (game Game) getResultScore(yourScore int, targetScore Score) int {
	if yourScore == targetScore.winningHand {
		return 6
	} else if yourScore == targetScore.drawingHand {
		return 3
	}
	return 0
}

func (game Game) getOwnHandScore(score int) int {
	switch score {
	case ROCK:
		return 1
	case PAPER:
		return 2
	default:
		return 3
	}
}

func getGameValue(input string) int {
	if input == "B" || input == "Y" {
		return PAPER
	} else if input == "C" || input == "Z" {
		return SCISSORS
	}
	return ROCK
}

func challengeTwo(scores []string) int {
	ownScore := 0
	for _, score := range scores {
		splittedScore := strings.Split(score, " ")
		game := &Game{
			firstSore:   getGameValue(splittedScore[0]),
			secondScore: getGameValue(splittedScore[1]),
		}
		ownScore += game.result()
	}
	return ownScore
}
