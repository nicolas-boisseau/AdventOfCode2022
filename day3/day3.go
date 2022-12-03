package day3

import (
	"errors"
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"regexp"
	"strings"
)

func RunToInt(r rune) int {
	rexp, _ := regexp.Compile(`[a-z]`)
	if rexp.MatchString(string(r)) {
		return int(r) - 97 + 1
	} else {
		return int(r) - (65 - 27)
	}
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	score := 0
	for i, line := range lines {

		if !complex {

			firstRucksack := line[0 : len(line)/2]
			secondRucksack := line[len(line)/2:]

			commonRune, err := findCommonRune(firstRucksack, secondRucksack)

			if err == nil {
				score += RunToInt(commonRune)
			} else {
				fmt.Println("ERROR")
			}
		} else {
			// level 2
			if i%3 == 0 {
				firstRucksackOfTheGroup := lines[i]
				secondRucksackOfTheGroup := lines[i+1]
				thirdRucksackOfTheGroup := lines[i+2]

				commonRune, err := findCommonRune(firstRucksackOfTheGroup, secondRucksackOfTheGroup, thirdRucksackOfTheGroup)
				if err == nil {
					score += RunToInt(commonRune)
				} else {
					fmt.Println("ERROR")
				}
			}
		}
	}

	return score
}

func findCommonRune(firstRucksack string, othersRucksacks ...string) (rune, error) {
	for _, r := range firstRucksack {
		isFoundInOther := true
		for _, otherRucksack := range othersRucksacks {
			if !strings.ContainsRune(otherRucksack, r) {
				isFoundInOther = false
			}
		}
		if isFoundInOther {
			return r, nil
		}
	}
	return '0', errors.New("Couldn't found a common rune")
}
