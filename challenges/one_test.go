package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeOne(t *testing.T) {
	calories := readInputFromFile("input-one")
	expectedMostCalories := 70374

	mostCalories := challengeOne(calories)
	fmt.Println(mostCalories)

	if mostCalories != expectedMostCalories {
		t.Fatalf("%d is not the amount of calories the elf with the most calories is carrying", mostCalories)
	}
}

func TestChallengeOnePartTwo(t *testing.T) {
	calories := readInputFromFile("input-one")
	expectedSumOfCalories := 204610

	sumOfCalories := challengeOnePartTwo(calories)

	if sumOfCalories != expectedSumOfCalories {
		t.Fatalf("%d is not the amount of the top 3 elves carrying calories", sumOfCalories)
	}
}
