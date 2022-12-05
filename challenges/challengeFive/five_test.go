package challengeFive

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeFive(t *testing.T) {
	expectedTopCrates := "VGBBJCRMN"
	input := challenges.ReadInputFromFile("input-five")

	topCrates := challengeFive(input)
	fmt.Println(topCrates)

	if topCrates != expectedTopCrates {
		t.Fatalf("%s are not the correct top crates", expectedTopCrates)
	}
}

func TestChallengeFivePartTwo(t *testing.T) {
	expectedTopCrates := "LBBVJBRMH"
	input := challenges.ReadInputFromFile("input-five")

	topCrates := challengeFivePartTwo(input)
	fmt.Println(topCrates)

	if topCrates != expectedTopCrates {
		t.Fatalf("%s are not the correct top crates", expectedTopCrates)
	}
}
