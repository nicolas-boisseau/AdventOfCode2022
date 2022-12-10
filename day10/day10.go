package day10

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
	"strings"
)

type Instruction struct {
	Cycles    int
	Operation func()
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	registryX := 1
	cycle := 1
	instructions := make([]*Instruction, 0)
	history := make(map[int]int)

	processOneCycle := func() {
		currentInstr := instructions[0]
		currentInstr.Cycles--
		if currentInstr.Cycles == 0 {
			currentInstr.Operation()
			instructions = instructions[1:]
		}

		// FOR LEVEL 2 (draw CRT lines)
		pos := cycle % 40
		spritePosition := registryX
		if pos == spritePosition-1 || pos == spritePosition || pos == spritePosition+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if pos == 0 {
			fmt.Println()
		}
		// END LEVEL 2

		cycle++
	}

	for _, line := range lines {
		words := strings.Split(line, " ")
		switch words[0] {
		case "addx":
			value, _ := strconv.Atoi(words[1])
			instructions = append(instructions, &Instruction{
				Cycles:    2,
				Operation: func() { registryX += value },
			})
		case "noop":
			instructions = append(instructions, &Instruction{
				Cycles:    1,
				Operation: func() {},
			})
		}

		//fmt.Printf("cycle n°%d = %d\n", cycle, registryX)
		history[cycle] = registryX

		// processing
		processOneCycle()
	}

	//fmt.Println("END OF INSTRUCTIONS")

	for len(instructions) > 0 {

		//fmt.Printf("cycle n°%d = %d\n", cycle, registryX)
		history[cycle] = registryX

		//removeProcessedOp()
		processOneCycle()

	}

	return history[20]*20 +
		history[60]*60 +
		history[100]*100 +
		history[140]*140 +
		history[180]*180 +
		history[220]*220

}
