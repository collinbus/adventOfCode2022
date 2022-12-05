package challengeThree

import (
	"log"
)

func ChallengeThree(input []string) int {
	sumPriorities := 0

	for _, ruckSack := range input {
		duplicate := findDuplicateInRucksack(ruckSack)
		sumPriorities += calculateScore(duplicate)
	}
	return sumPriorities
}

func ChallengeThreePartTwo(input []string) int {
	sumPriorities := 0
	for i := 0; i < len(input); i += 3 {
		duplicate := findDuplicateInSeries(input[i : i+3])
		sumPriorities += calculateScore(duplicate)
	}
	return sumPriorities
}

func calculateScore(value uint8) int {
	lowerCaseScoreOffset := uint8(96)
	upperCaseScoreOffset := uint8(38)
	if value > lowerCaseScoreOffset {
		return int(value - lowerCaseScoreOffset)
	} else {
		return int(value - upperCaseScoreOffset)
	}
}

func findDuplicateInSeries(ruckSack []string) uint8 {
	items := make(map[uint8]int)
	for i := range ruckSack[0] {
		item := ruckSack[0][i]
		items[item] = 1
	}
	for i := range ruckSack[1] {
		item := ruckSack[1][i]
		if occurrences, exists := items[item]; exists {
			if occurrences == 1 {
				items[item] = 2
			}
		}
	}
	for i := range ruckSack[2] {
		item := ruckSack[2][i]
		if occurrences, exists := items[item]; exists {
			if occurrences == 2 {
				return item
			}
		}
	}
	log.Fatal("No doubles in series!")
	return 0
}

func findDuplicateInRucksack(ruckSack string) uint8 {
	length := len(ruckSack)
	middle := length / 2
	items := make(map[uint8]bool)
	for i := 0; i < middle; i++ {
		item := ruckSack[i]
		items[item] = true
	}
	for i := middle; i < length; i++ {
		item := ruckSack[i]
		if _, exists := items[item]; exists {
			return item
		}
	}
	log.Fatal("No doubles in rucksack!")
	return 0
}
