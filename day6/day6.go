package day6

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

func isAllDifferent2(currentWindow string) bool {

	return linq.From(currentWindow).
		GroupByT(func(r rune) rune { return r }, func(r rune) rune { return r }).
		Count() == len(currentWindow)
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
