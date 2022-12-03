package challenges

import (
	"fmt"
	"testing"
)

func TestChallengeThree(t *testing.T) {
	input := readInputFromFile("input-three")

	priorities := ChallengeThree(input)

	fmt.Println(priorities)
}
