package challenges

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestChallengeOne(t *testing.T) {
	calories := readInputFromFile()
	expectedMostCalories := 70374

	mostCalories := challengeOne(calories)
	fmt.Println(mostCalories)

	if mostCalories != expectedMostCalories {
		t.Fatalf("%d is not the amount of calories the elf with the most calories is carrying", mostCalories)
	}
}

func TestChallengeOnePartTwo(t *testing.T) {
	calories := readInputFromFile()
	expectedSumOfCalories := 204610

	sumOfCalories := challengeOnePartTwo(calories)

	if sumOfCalories != expectedSumOfCalories {
		t.Fatalf("%d is not the amount of the top 3 elves carrying calories", sumOfCalories)
	}
}

func readInputFromFile() []string {
	var content []string
	file, err := os.Open("input-one")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}
