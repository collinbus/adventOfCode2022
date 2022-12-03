package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeThree(t *testing.T) {
	expectedPriorities := 7727
	input := readInputFromFile("input-three")

	priorities := ChallengeThree(input)
	fmt.Println(priorities)

	if priorities != expectedPriorities {
		t.Fatalf("%d is not the right sum of priorities", priorities)
	}
}

func TestChallengeThreePartTwo(t *testing.T) {
	expectedPriorities := 2609
	input := readInputFromFile("input-three")

	priorities := ChallengeThreePartTwo(input)
	fmt.Println(priorities)

	if priorities != expectedPriorities {
		t.Fatalf("%d is not the right sum of priorities", priorities)
	}
}
