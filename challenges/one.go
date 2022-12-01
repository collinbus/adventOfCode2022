package challenges

import "strconv"

func challengeOne(calories []string) int {
	mostCalories := 0
	carryingCalories := 0

	for _, caloryString := range calories {
		if caloryString == "" {
			if carryingCalories > mostCalories {
				mostCalories = carryingCalories
			}
			carryingCalories = 0
		} else {
			calory, _ := strconv.Atoi(caloryString)
			carryingCalories += calory
		}
	}

	return mostCalories
}
