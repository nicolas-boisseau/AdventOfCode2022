package day6

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

func isAllDifferent2(currentWindow string) bool {

	for i, r1 := range currentWindow {
		for j, r2 := range currentWindow {
			if i == j {
				continue
			}

			if r1 == r2 {
				return false
			}
		}
	}

	return true
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	windowSize := 4
	if complex {
		windowSize = 14
	}

	signal := 0
	for i := windowSize; i < len(lines[0]); i++ {
		currentWindow := lines[0][i-windowSize : i]

		if isAllDifferent2(currentWindow) {
			fmt.Println("SIGNAL DETECTED! ", i)
			signal = i
			break
		}

	}

	return signal
}
