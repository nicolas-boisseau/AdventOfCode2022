package day13

import (
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	return len(lines)
}
