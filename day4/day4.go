package day4

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

func isOverlapped(start1, end1, start2, end2 int) bool {
	return (start1 >= start2 && end1 <= end2) || (start2 >= start1 && end2 <= end1)
}

func isOverlappedComplex(start1, end1, start2, end2 int) bool {
	return ((start1 >= start2 && start1 <= end2) || (end1 <= end2 && end1 >= start2)) ||
		((start2 >= start1 && start2 <= end1) || (end2 <= end1 && end2 >= start1))
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	score := 0
	for _, line := range lines {

		var start_1, end_1, start_2, end_2 int
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "%d-%d,%d-%d", &start_1, &end_1, &start_2, &end_2)

		if !complex && isOverlapped(start_1, end_1, start_2, end_2) {
			score++
		} else {
			if isOverlappedComplex(start_1, end_1, start_2, end_2) {
				score++
			}
		}

	}

	return score
}
