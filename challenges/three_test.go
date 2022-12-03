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
