package challengeSix

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeSix(t *testing.T) {
	expectedFirstCharacterAvailable := 1598
	input := challenges.ReadInputFromFile("input-six")

	firstCharacterAvailable := challengeSix(input[0])
	fmt.Println(firstCharacterAvailable)

	if firstCharacterAvailable != expectedFirstCharacterAvailable {
		t.Fatalf("%d is not the correct index of the first character available\n", firstCharacterAvailable)
	}
}

func TestChallengeSixPartTwo(t *testing.T) {
	expectedFirstCharacterAvailable := 2414
	input := challenges.ReadInputFromFile("input-six")

	firstCharacterAvailable := challengeSixPartTwo(input[0])
	fmt.Println(firstCharacterAvailable)

	if firstCharacterAvailable != expectedFirstCharacterAvailable {
		t.Fatalf("%d is not the correct index of the first character available\n", firstCharacterAvailable)
	}
}
