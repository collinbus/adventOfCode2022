package challengeFour

import (
	"strconv"
	"strings"
)

type Schedule struct {
	start int
	end   int
}

func challengeFour(input []string) int {
	fullyContained := 0
	for _, schedule := range input {
		firstSchedule, secondSchedule := setupSchedules(schedule)
		if isFullyContained(firstSchedule, secondSchedule) {
			fullyContained += 1
		}
	}
	return fullyContained
}

func challengeFourPartTwo(input []string) int {
	partiallyContained := 0
	for _, schedule := range input {
		firstSchedule, secondSchedule := setupSchedules(schedule)
		if isPartiallyContained(firstSchedule, secondSchedule) {
			partiallyContained += 1
		}
	}
	return partiallyContained
}

func setupSchedules(schedule string) (Schedule, Schedule) {
	cleanAreas := strings.Split(schedule, ",")
	firstStringSchedule := strings.Split(cleanAreas[0], "-")
	start, _ := strconv.Atoi(firstStringSchedule[0])
	end, _ := strconv.Atoi(firstStringSchedule[1])
	firstSchedule := Schedule{
		start: start,
		end:   end,
	}
	secondStringSchedule := strings.Split(cleanAreas[1], "-")
	start, _ = strconv.Atoi(secondStringSchedule[0])
	end, _ = strconv.Atoi(secondStringSchedule[1])
	secondSchedule := Schedule{
		start: start,
		end:   end,
	}
	return firstSchedule, secondSchedule
}

func isPartiallyContained(firstSchedule Schedule, secondSchedule Schedule) bool {
	lowerBounds := min(firstSchedule.start, secondSchedule.start)
	if firstSchedule.start == lowerBounds {
		if firstSchedule.end >= secondSchedule.start || firstSchedule.end >= secondSchedule.end {
			return true
		}
	}
	if secondSchedule.start == lowerBounds {
		if secondSchedule.start <= firstSchedule.start && secondSchedule.end >= firstSchedule.start {
			return true
		}
	}
	return false
}

func isFullyContained(firstSchedule Schedule, secondSchedule Schedule) bool {
	lowerBounds := min(firstSchedule.start, secondSchedule.start)
	if firstSchedule.start == lowerBounds {
		if secondSchedule.end <= firstSchedule.end && secondSchedule.start <= firstSchedule.end {
			return true
		}
	}
	if secondSchedule.start == lowerBounds {
		if firstSchedule.start >= secondSchedule.start && firstSchedule.end <= secondSchedule.end {
			return true
		}
	}
	return false
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
