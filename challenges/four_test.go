package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeFour(t *testing.T) {
	input := readInputFromFile("input-four")

	solelyContainedPairs := challengeFour(input)
	fmt.Println(solelyContainedPairs)

}

func TestChallengeFourExtraInput(t *testing.T) {
	expectedPairs := 3
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
		"28-81,28-82",
	}

	solelyContainedPairs := challengeFour(input)
	fmt.Println(solelyContainedPairs)

	if solelyContainedPairs != expectedPairs {
		t.Fatalf("%d is not the right amount of pairs", solelyContainedPairs)
	}
}
