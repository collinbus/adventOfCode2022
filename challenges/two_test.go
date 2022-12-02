package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeTwo(t *testing.T) {
	expectedEndScore := 12794
	scores := readInputFromFile("input-two")

	endScore := challengeTwo(scores)
	fmt.Println(endScore)

	if endScore != expectedEndScore {
		t.Fatalf("%d is not the correct endScore\n", endScore)
	}
}
