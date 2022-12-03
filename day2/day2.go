package day2

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	leftRules := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissor",
	}

	rightRules := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissor",
	}

	scoreByHandType := map[string]int{
		"Rock":    1,
		"Paper":   2,
		"Scissor": 3,
	}

	score := 0
	for _, line := range lines {
		var left, right string
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "%s %s", &left, &right)

		//fmt.Println("Left:", left, "right:", right)
		//fmt.Println("Left:", leftRules[left], "right:", rightRules[right])

		leftHand := leftRules[left]
		rightHand := rightRules[right]

		if !complex {
			score += scoreByHandType[rightHand]

			if leftHand == rightHand {
				score += 3
			} else {
				game := leftHand + "-" + rightHand
				switch game {
				case "Paper-Scissor", "Rock-Paper", "Scissor-Rock":
					score += 6
				default:
				}
			}
		} else {
			// level 2 !

			switch right {
			case "X": // need to lose
				switch leftHand {
				case "Rock":
					score += scoreByHandType["Scissor"]
				case "Paper":
					score += scoreByHandType["Rock"]
				case "Scissor":
					score += scoreByHandType["Paper"]
				}
			case "Y": // need to draw
				score += scoreByHandType[leftHand]
				score += 3
			case "Z":
				switch leftHand {
				case "Rock":
					score += scoreByHandType["Paper"]
				case "Paper":
					score += scoreByHandType["Scissor"]
				case "Scissor":
					score += scoreByHandType["Rock"]
				}
				score += 6
			}
		}
	}

	return score
}
