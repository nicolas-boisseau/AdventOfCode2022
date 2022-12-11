package day11

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
	"strings"
)

type Monkey struct {
	items             []int
	operation         func(int) int
	test              func(int) bool
	diviser           int
	nextMonkeyIfTrue  int
	nextMonkeyIfFalse int
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	monkeys := make([]*Monkey, 0)
	for i, line := range lines {
		if len(line) >= 6 && line[0:6] == "Monkey" {

			// read items
			startingItemsRaw := lines[i+1][18:]
			startingItemsStrList := strings.Split(startingItemsRaw, ", ")
			var startingItems []int
			linq.From(startingItemsStrList).SelectT(func(str string) int {
				val, _ := strconv.Atoi(str)
				return val
			}).ToSlice(&startingItems)

			// Read operation
			var op string
			var op_value string
			reader := strings.NewReader(lines[i+2])
			fmt.Fscanf(reader, "  Operation: new = old %s %s", &op, &op_value)

			// Read test
			var diviser int
			reader = strings.NewReader(lines[i+3])
			fmt.Fscanf(reader, "  Test: divisible by %d", &diviser)

			// Read next monkeys
			var nextMonkeyIfTrue int
			reader = strings.NewReader(lines[i+4])
			fmt.Fscanf(reader, "    If true: throw to monkey %d", &nextMonkeyIfTrue)
			var nextMonkeyIfFalse int
			reader = strings.NewReader(lines[i+5])
			fmt.Fscanf(reader, "    If false: throw to monkey %d", &nextMonkeyIfFalse)

			// build monkeys
			monkeys = append(monkeys, &Monkey{
				items: startingItems,
				operation: func(old int) int {
					op_value_num := old
					if op_value != "old" {
						op_value_num, _ = strconv.Atoi(op_value)
					}
					switch op {
					case "+":
						return old + op_value_num
					case "*":
						return old * op_value_num
					}
					return old
				},
				test: func(v int) bool {
					return v%diviser == 0
				},
				diviser:           diviser,
				nextMonkeyIfTrue:  nextMonkeyIfTrue,
				nextMonkeyIfFalse: nextMonkeyIfFalse,
			})
		}
	}

	inspectionsByMonkey := make(map[int]int)
	numberOfRounds := 20
	if complex {
		numberOfRounds = 10000
	}

	// limit to only the divisers that we need to watch
	modulo := linq.From(monkeys).
		SelectT(func(monkey *Monkey) int { return monkey.diviser }).
		AggregateT(func(v, v2 int) int { return v * v2 }).(int)

	for round := 1; round <= numberOfRounds; round++ {

		for currentMonkeyIndex, currentMonkey := range monkeys {

			for len(currentMonkey.items) > 0 {

				currentItem := currentMonkey.items[0]
				currentItem = currentMonkey.operation(currentItem)
				if !complex {
					currentItem /= 3
				} else {
					currentItem %= modulo
				}
				if currentMonkey.test(currentItem) {
					monkeys[currentMonkey.nextMonkeyIfTrue].items =
						append(monkeys[currentMonkey.nextMonkeyIfTrue].items, currentItem)
				} else {
					monkeys[currentMonkey.nextMonkeyIfFalse].items =
						append(monkeys[currentMonkey.nextMonkeyIfFalse].items, currentItem)
				}

				inspectionsByMonkey[currentMonkeyIndex]++

				currentMonkey.items = currentMonkey.items[1:]
			}
		}

		if round == 1 || round == 20 || round == 1000 || round == 2000 {
			//fmt.Println("Monkeys stats after round", round)
			//for _, currentMonkey := range monkeys {
			//	fmt.Println(currentMonkey)
			//}
			//fmt.Println(inspectionsByMonkey)
		}
	}

	fmt.Println(inspectionsByMonkey)

	topInspections :=
		linq.From(inspectionsByMonkey).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(int) }).
			OrderByDescendingT(func(v int) int { return v }).
			Take(2).
			AggregateT(func(v1, v2 int) int { return v1 * v2 }).(int)

	return topInspections
}
