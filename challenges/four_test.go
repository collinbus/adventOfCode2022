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

func TestChallengeFourPartTwo(t *testing.T) {
	input := readInputFromFile("input-four")

	partiallyContainedPairs := challengeFourPartTwo(input)
	fmt.Println(partiallyContainedPairs)
}

func TestChallengeFourPartTwoExtraInput(t *testing.T) {
	expectedPairs := 4
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
		"75-75,72-74",
	}

	partiallyContainedPairs := challengeFourPartTwo(input)
	fmt.Println(partiallyContainedPairs)

	if partiallyContainedPairs != expectedPairs {
		t.Fatalf("%d is not the right amount of pairs", partiallyContainedPairs)
	}
}
