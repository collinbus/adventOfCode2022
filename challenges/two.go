package challenges

import "strings"

type Game struct {
	firstSore   string
	secondScore string
}

type Score struct {
	winningHand string
	losingHand  string
	drawingHand string
}

type RockPaperScissorGame interface {
	result() int
}

func (game Game) result() int {
	scoreTable := make(map[string]Score)
	scoreTable["A"] = Score{
		winningHand: "Y",
		losingHand:  "Z",
		drawingHand: "X",
	}
	scoreTable["B"] = Score{
		winningHand: "Z",
		losingHand:  "X",
		drawingHand: "Y",
	}
	scoreTable["C"] = Score{
		winningHand: "X",
		losingHand:  "Y",
		drawingHand: "Z",
	}

	currentScore := scoreTable[game.firstSore]
	ownHandScore := game.getOwnHandScore()
	resultScore := game.getResultScore(game.secondScore, currentScore)
	return ownHandScore + resultScore
}

func (game Game) getResultScore(yourScore string, targetScore Score) int {
	if yourScore == targetScore.winningHand {
		return 6
	} else if yourScore == targetScore.drawingHand {
		return 3
	}
	return 0
}

func (game Game) getOwnHandScore() int {
	switch game.secondScore {
	case "X":
		return 1
	case "Y":
		return 2
	default:
		return 3
	}
}

func challengeTwo(scores []string) int {
	ownScore := 0
	for _, score := range scores {
		splittedScore := strings.Split(score, " ")
		game := &Game{
			firstSore:   splittedScore[0],
			secondScore: splittedScore[1],
		}
		ownScore += game.result()
	}
	return ownScore
}
