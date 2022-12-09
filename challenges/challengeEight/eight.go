package challengeEight

type Tree struct {
	visible bool
	height  int
}

func NewTree(height int) *Tree {
	return &Tree{
		visible: false,
		height:  height,
	}
}

func challengeEight(input []string) int {
	treeMap := createMap(input)
	visibleTrees := countInnerVisibleTrees(treeMap)
	return visibleTrees
}

func challengeEightPartTwo(input []string) int {
	treeMap := createMap(input)
	countInnerVisibleTrees(treeMap)
	highestScenicScore := calculateHighestScenicScore(treeMap)
	return highestScenicScore
}

func calculateHighestScenicScore(treeMap [][]*Tree) int {
	highestSceneryScore := 0
	for row := 0; row < len(treeMap); row++ {
		for col := 0; col < len(treeMap[row]); col++ {
			score := calculateSceneryScore(treeMap, row, col)
			if score > highestSceneryScore {
				highestSceneryScore = score
			}
		}
	}
	return highestSceneryScore
}

func calculateSceneryScore(treeMap [][]*Tree, row int, col int) int {
	score := 1
	current := treeMap[row][col]
	distance := 0
	for i := col - 1; i >= 0; i-- {
		distance++
		if treeMap[row][i].height >= current.height {
			break
		}
	}

	score *= distance

	distance = 0
	for i := col + 1; i < len(treeMap[row]); i++ {
		distance++
		if treeMap[row][i].height >= current.height {
			break
		}
	}

	score *= distance

	distance = 0
	for i := row - 1; i >= 0; i-- {
		distance++
		if treeMap[i][col].height >= current.height {
			break
		}
	}

	score *= distance

	distance = 0
	for i := row + 1; i < len(treeMap); i++ {
		distance++
		if treeMap[i][col].height >= current.height {
			break
		}
	}

	score *= distance

	return score
}

func countInnerVisibleTrees(treeMap [][]*Tree) int {
	visibleTrees := 0
	for row := 0; row < len(treeMap); row++ {
		for col := 0; col < len(treeMap[row]); col++ {
			treeMap[row][col].visible = true
			if row == 0 || row == len(treeMap)-1 {
				visibleTrees++
			} else if col == 0 || col == len(treeMap[row])-1 {
				visibleTrees++
			} else if isVisible(treeMap, row, col) {
				visibleTrees++
			} else {
				treeMap[row][col].visible = false
			}
		}
	}
	return visibleTrees
}

func isVisible(treeMap [][]*Tree, row int, col int) bool {
	current := treeMap[row][col]
	leftVisible := true
	rightVisible := true
	topVisible := true
	bottomVisible := true
	for i := col - 1; i >= 0; i-- {
		if treeMap[row][i].height >= current.height {
			leftVisible = false
		}
	}
	for i := col + 1; i < len(treeMap[row]); i++ {
		if treeMap[row][i].height >= current.height {
			rightVisible = false
		}
	}
	for i := row - 1; i >= 0; i-- {
		if treeMap[i][col].height >= current.height {
			topVisible = false
		}
	}
	for i := row + 1; i < len(treeMap); i++ {
		if treeMap[i][col].height >= current.height {
			bottomVisible = false
		}
	}
	return leftVisible || rightVisible || topVisible || bottomVisible
}

func createMap(input []string) [][]*Tree {
	asciiOffset := 48
	height := len(input)
	width := len(input[0])
	treeMap := make([][]*Tree, height)
	for r := range input {
		row := make([]*Tree, width)
		for c := range input[r] {
			row[c] = NewTree(int(input[r][c]) - asciiOffset)
		}
		treeMap[r] = row
	}
	return treeMap
}
