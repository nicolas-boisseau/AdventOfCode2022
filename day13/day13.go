package day13

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	n := ReadNode(lines[0])
	fmt.Println(n)

	return len(lines)
}
