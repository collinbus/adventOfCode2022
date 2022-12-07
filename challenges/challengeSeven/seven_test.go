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

func TestChallengeSevenPartTwo(t *testing.T) {
	expectedSmallestDirectorySize := 8278005
	input := challenges.ReadInputFromFile("input-seven")

	smallestDirectorySize := challengeSevenPartTwo(input)
	fmt.Println(smallestDirectorySize)

	if smallestDirectorySize != expectedSmallestDirectorySize {
		t.Fatalf("%d is not the correct smallest directorySize", smallestDirectorySize)
	}
}
