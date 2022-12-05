package challengeTwo

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeTwo(t *testing.T) {
	expectedEndScore := 12794
	scores := challenges.ReadInputFromFile("input-two")

	endScore := challengeTwo(scores)
	fmt.Println(endScore)

	if endScore != expectedEndScore {
		t.Fatalf("%d is not the correct endScore\n", endScore)
	}
}

func TestChallengeTwoPartTwo(t *testing.T) {
	expectedEndScore := 14979
	scores := challenges.ReadInputFromFile("input-two")

	endScore := challengeTwoPartTwo(scores)
	fmt.Println(endScore)

	if endScore != expectedEndScore {
		t.Fatalf("%d is not the correct endScore\n", endScore)
	}
}
