package challengeTen

import (
	"adventOfCode2022/challenges"
	"fmt"
	"testing"
)

func TestChallengeTen(t *testing.T) {
	expectedSignalSum := 13180
	input := challenges.ReadInputFromFile("input-ten")

	sumSignalStrengths := challengeTen(input)
	fmt.Println(sumSignalStrengths)

	if sumSignalStrengths != expectedSignalSum {
		t.Fatalf("%d is not the correct signal sum\n", sumSignalStrengths)
	}
}
