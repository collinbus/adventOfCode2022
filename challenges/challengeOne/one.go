package challengeOne

import (
	"fmt"
	"github.com/gammazero/deque"
	"strconv"
)

func challengeOne(calories []string) int {
	top3mostCalories := getTop3MostCalories(calories)
	return top3mostCalories.At(0)
}

func challengeOnePartTwo(calories []string) int {
	sumOfCalories := 0
	top3mostCalories := getTop3MostCalories(calories)

	for i := 0; i < top3mostCalories.Len(); i++ {
		calory := top3mostCalories.At(i)
		fmt.Printf("Elf %d is carrying %d\n", 1+i, calory)
		sumOfCalories += calory
	}

	fmt.Printf("The sum of the top three calories is %d\n", sumOfCalories)

	return sumOfCalories
}

func getTop3MostCalories(calories []string) *deque.Deque[int] {
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
