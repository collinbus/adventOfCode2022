package challengeSix

import "github.com/golang-collections/collections/set"

func challengeSix(input string) int {

	for i := 3; i < len(input); i++ {
		charactersSet := set.New()
		for j := i - 3; j <= i; j++ {
			charactersSet.Insert(input[j])
		}
		if charactersSet.Len() == 4 {
			return i + 1
		}
	}
	return 0
}
