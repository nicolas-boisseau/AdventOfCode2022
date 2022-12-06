package day5

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"regexp"
	"strconv"
	"strings"
)

func Process(fileName string, complex bool) string {
	lines := common.ReadLinesFromFile(fileName)

	// FINDING THE BOTTOM OF THE STACKS
	rexp, _ := regexp.Compile(`[0-9]+`)
	indexOfStackStart := 0
	for index, line := range lines {
		if rexp.MatchString(line) {
			indexOfStackStart = index
			break
		}
	}

	// READING STACKS CONTENT
	rexpLetter, _ := regexp.Compile("[A-Z]+")
	stacks := make(map[string]*stack.Stack)
	for i := 1; i < len(lines[indexOfStackStart]) && rexp.MatchString(string(lines[indexOfStackStart][i])); i += 4 {

		currentStackIndex := string(lines[indexOfStackStart][i])
		stacks[currentStackIndex] = stack.New()

		for j := indexOfStackStart - 1; j >= 0 && rexpLetter.MatchString(string(lines[j][i])); j-- {
			stacks[currentStackIndex].Push(string(lines[j][i]))
		}
	}

	// PRINT STACKS FOR DEBUG
	//for k, s := range stacks {
	//	fmt.Println("Stack", k, ":")
	//	for s.Len() > 0 {
	//		fmt.Println(s.Pop())
	//	}
	//}

	// READING INSTRUCTIONS (2 lines after stacks content)
	for _, line := range lines[indexOfStackStart+2:] {
		var count int
		var s1, s2 string
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "move %d from %s to %s", &count, &s1, &s2)

		// APPLY INSTRUCTION
		if !complex {
			for i := 0; i < count; i++ {
				stacks[s2].Push(stacks[s1].Pop())
			}
		} else {
			// LEVEL 2
			tmpStack := stack.New()
			for i := 0; i < count; i++ {
				tmpStack.Push(stacks[s1].Pop())
			}
			for tmpStack.Len() > 0 {
				stacks[s2].Push(tmpStack.Pop())
			}
		}
	}

	// PRINT RESULT (top of each stacks)
	result := ""
	for i := 1; i < len(stacks)+1; i++ {
		ind := strconv.FormatInt(int64(i), 10)
		result += stacks[ind].Peek().(string)
	}

	return result
}
