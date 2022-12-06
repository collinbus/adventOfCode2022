package challengeSix

import "github.com/golang-collections/collections/set"

func challengeSix(input string) int {
	return findFirstDistinctCharacterIndex(input, 4)
}

func challengeSixPartTwo(input string) int {
	return findFirstDistinctCharacterIndex(input, 14)
}

func findFirstDistinctCharacterIndex(input string, amount int) int {
	offset := amount - 1
	for i := offset; i < len(input); i++ {
		charactersSet := set.New()
		for j := i - offset; j <= i; j++ {
			charactersSet.Insert(input[j])
		}
		if charactersSet.Len() == amount {
			return i + 1
		}
	}
	return 0
}
