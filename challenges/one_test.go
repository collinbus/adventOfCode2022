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

	mostCalories := challengeOne(calories)

	fmt.Println(mostCalories)
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
