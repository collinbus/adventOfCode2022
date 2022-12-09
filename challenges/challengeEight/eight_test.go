package challengeEight

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeEight(t *testing.T) {
	expectedVisibleTrees := 1870
	input := challenges.ReadInputFromFile("input-eight")

	amountOfTreesVisible := challengeEight(input)
	fmt.Println(amountOfTreesVisible)

	if amountOfTreesVisible != expectedVisibleTrees {
		t.Fatalf("%d is not the expected amount of visible trees\n", amountOfTreesVisible)
	}
}

func TestChallengeEightPartTwo(t *testing.T) {
	expectedHighestSceneryScore := 517440
	input := challenges.ReadInputFromFile("input-eight")

	highestSceneryScore := challengeEightPartTwo(input)
	fmt.Println(highestSceneryScore)

	if highestSceneryScore != expectedHighestSceneryScore {
		t.Fatalf("%d is not the expected amount of visible trees\n", highestSceneryScore)
	}
}
