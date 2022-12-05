package challenges

import (
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
)

type Move struct {
	amount int
	from   int
	to     int
}

func challengeFive(input []string) string {
	stacks, startMoveIndex := parseCrateInput(input)
	stacks = performMovesCrateMover9000(stacks, input[startMoveIndex:])
	topStackString := ""
	for _, currentStack := range stacks {
		topStackString += currentStack.Peek().(string)
	}
	return topStackString
}

func challengeFivePartTwo(input []string) string {
	stacks, startMoveIndex := parseCrateInput(input)
	stacks = performMovesCrateMover9001(stacks, input[startMoveIndex:])
	topStackString := ""
	for _, currentStack := range stacks {
		topStackString += currentStack.Peek().(string)
	}
	return topStackString
}

func performMovesCrateMover9001(stacks []stack.Stack, moves []string) []stack.Stack {
	for _, moveString := range moves {
		move := parseMove(moveString)
		elements := make([]string, 0)
		for i := 0; i < move.amount; i++ {
			element := stacks[move.from-1].Pop().(string)
			elements = append(elements, element)
		}
		for i := len(elements) - 1; i >= 0; i-- {
			stacks[move.to-1].Push(elements[i])
		}
	}
	return stacks
}

func performMovesCrateMover9000(stacks []stack.Stack, moves []string) []stack.Stack {
	for _, moveString := range moves {
		move := parseMove(moveString)
		for i := 0; i < move.amount; i++ {
			element := stacks[move.from-1].Pop()
			stacks[move.to-1].Push(element)
		}
	}
	return stacks
}

func parseMove(moveString string) Move {
	splittedMove := strings.Split(moveString, "from")
	splittedMoveWithoutSpaces := strings.ReplaceAll(splittedMove[0], " ", "")
	splittedOperationWithoutSpaces := strings.ReplaceAll(splittedMove[1], " ", "")
	operationArray := strings.Split(splittedOperationWithoutSpaces, "to")
	amount, _ := strconv.Atoi(strings.ReplaceAll(splittedMoveWithoutSpaces, "move", ""))
	from, _ := strconv.Atoi(operationArray[0])
	to, _ := strconv.Atoi(operationArray[1])
	return Move{
		amount: amount,
		from:   from,
		to:     to,
	}
}

// returns the stacks and the start index for the moves
func parseCrateInput(input []string) ([]stack.Stack, int) {
	inputStack := make([]string, 0)
	numberOfStacks := 0
	startMoveIndex := 0
	for i := 0; i < len(input); i++ {
		currentLine := input[i]
		if strings.Contains(currentLine, "1") {
			indices := strings.Split(currentLine, " ")
			lengthIndex := len(indices) - 1
			startMoveIndex = i + 2
			numberOfStacks, _ = strconv.Atoi(indices[lengthIndex])
			break
		} else {
			inputStack = append(inputStack, currentLine)
		}
	}

	stacks := make([]stack.Stack, numberOfStacks)

	for i := numberOfStacks - 2; i >= 0; i-- {
		for j := 0; j < len(stacks); j++ {
			index := (j * 4) + 1
			inputString := inputStack[i]
			asciiChar := inputString[index]
			if asciiChar < 91 && asciiChar > 64 {
				stacks[j].Push(string(rune(asciiChar)))
			}
		}
	}

	return stacks, startMoveIndex
}
