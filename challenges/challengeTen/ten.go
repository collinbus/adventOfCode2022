package challengeTen

import (
	"fmt"
	"strconv"
	"strings"
)

const NO_OPERATION = 0
const ADD = 1

type Instruction struct {
	operation int
	amount    int
}

type Cycle struct {
	cycleNumber int
	instruction *Instruction
}

func NewCycle(cycleNumber int, instruction *Instruction) *Cycle {
	return &Cycle{cycleNumber: cycleNumber, instruction: instruction}
}

func NewInstruction(operation int, amount int) *Instruction {
	return &Instruction{operation: operation, amount: amount}
}

func challengeTen(input []string) int {
	cycleNumber := 1
	var cycles []*Cycle
	cycleToCheckIndex := 0
	cyclesToCheck := []int{20, 60, 100, 140, 180, 220}
	for _, instructionString := range input {
		splitInstruction := strings.Split(instructionString, " ")
		var instruction *Instruction
		if splitInstruction[0] == "addx" {
			amount, _ := strconv.Atoi(splitInstruction[1])
			instruction = NewInstruction(ADD, amount)
			cycleNumber++
		} else {
			instruction = NewInstruction(NO_OPERATION, 0)
		}
		cycles = append(cycles, NewCycle(cycleNumber, instruction))
		cycleNumber++
	}

	cycleIndex := 0
	currentSignalStrength := 1
	signalSum := 0
	for currentCycleNumber := 1; currentCycleNumber <= 220; currentCycleNumber++ {
		if cycleToCheckIndex >= len(cyclesToCheck) {
			continue
		}
		if currentCycleNumber == cyclesToCheck[cycleToCheckIndex] {
			fmt.Printf("Cycle %d: current sum: %d\n", currentCycleNumber, currentSignalStrength*currentCycleNumber)
			signalSum += currentSignalStrength * currentCycleNumber
			cycleToCheckIndex++
		}

		if cycleIndex >= len(cycles) {
			continue
		}
		currentCycle := cycles[cycleIndex]
		if cycles[cycleIndex].cycleNumber == currentCycleNumber {
			//fmt.Printf("Cycle %d: current amount: %d\n", currentCycleNumber, currentCycle.instruction.amount)
			currentSignalStrength += currentCycle.instruction.amount
			//fmt.Printf("Cycle %d current strength %d\n", currentCycleNumber, currentSignalStrength)
			cycleIndex++
		}
	}

	return signalSum
}
