package challenges

import (
	"log"
)

func ChallengeThree(input []string) int {
	sumPriorities := 0
	lowerCaseScoreOffset := uint8(96)
	upperCaseScoreOffset := uint8(38)

	for _, ruckSack := range input {
		duplicate := findDuplicateInRucksack(ruckSack)
		if duplicate > lowerCaseScoreOffset {
			sumPriorities += int(duplicate - lowerCaseScoreOffset)
		} else {
			sumPriorities += int(duplicate - upperCaseScoreOffset)
		}
	}
	return sumPriorities
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
