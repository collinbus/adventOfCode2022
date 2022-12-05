package challengeThree

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeThree(t *testing.T) {
	expectedPriorities := 7727
	input := challenges.ReadInputFromFile("input-three")

	priorities := ChallengeThree(input)
	fmt.Println(priorities)

	if priorities != expectedPriorities {
		t.Fatalf("%d is not the right sum of priorities", priorities)
	}
}

func TestChallengeThreePartTwo(t *testing.T) {
	expectedPriorities := 2609
	input := challenges.ReadInputFromFile("input-three")

	priorities := ChallengeThreePartTwo(input)
	fmt.Println(priorities)

	if priorities != expectedPriorities {
		t.Fatalf("%d is not the right sum of priorities", priorities)
	}
}
