package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeFive(t *testing.T) {
	expectedTopCrates := "VGBBJCRMN"
	input := readInputFromFile("input-five")

	topCrates := challengeFive(input)
	fmt.Println(topCrates)

	if topCrates != expectedTopCrates {
		t.Fatalf("%s are not the correct top crates", expectedTopCrates)
	}
}
