package challenges

import (
	"github.com/gammazero/deque"
	"strconv"
)

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

func challengeOnePartTwo(calories []string) *deque.Deque[int] {
	mostCalories := deque.New[int](3)
	mostCalories.PushBack(0)
	mostCalories.PushBack(0)
	mostCalories.PushBack(0)

	carryingCalories := 0

	for _, caloryString := range calories {
		if caloryString == "" {
			mostCalories = updateMostCarryingCalories(mostCalories, carryingCalories)
			carryingCalories = 0
		} else {
			calory, _ := strconv.Atoi(caloryString)
			carryingCalories += calory
		}
	}

	return mostCalories
}

func updateMostCarryingCalories(calories *deque.Deque[int], currentCalory int) *deque.Deque[int] {
	for i := 0; i < calories.Len(); i++ {
		if calories.At(i) < currentCalory {
			calories.Insert(i, currentCalory)
			calories.PopBack()
			return calories
		}
	}
	return calories
}
