package challenges

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
		if isFullyContained(firstSchedule, secondSchedule) {
			fullyContained += 1
		}
	}
	return fullyContained
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
