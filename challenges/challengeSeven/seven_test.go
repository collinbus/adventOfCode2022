package challengeSeven

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeSeven(t *testing.T) {
	expectedDirectorySizes := 1141028
	input := challenges.ReadInputFromFile("input-seven")

	largeDirectorySizes := challengeSeven(input)
	fmt.Println(largeDirectorySizes)

	if largeDirectorySizes != expectedDirectorySizes {
		t.Fatalf("%d is not the correct size of all directories that can be removed", largeDirectorySizes)
	}
}
