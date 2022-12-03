package day1

import (
	. "github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
)

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	elfs := make(map[int]int, 0)

	currentElf := 1
	for _, line := range lines {

		calories, err := strconv.ParseInt(line, 10, 32)
		if err == nil {
			elfs[currentElf] += int(calories)
		} else {
			currentElf++
		}
	}

	if !complex {
		return From(elfs).SelectT(func(kv KeyValue) int { return kv.Value.(int) }).Max().(int)
	} else {
		sumOfThreeMostCalories := From(elfs).
			SelectT(func(kv KeyValue) int { return kv.Value.(int) }).
			OrderByDescendingT(func(cal int) int { return cal }).
			Take(3).
			SumInts()

		return int(sumOfThreeMostCalories)
	}
}
