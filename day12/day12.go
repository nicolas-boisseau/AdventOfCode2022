package day12

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"regexp"
)

func RuneToInt(r rune) int {
	rexp, _ := regexp.Compile(`[a-z]`)
	if rexp.MatchString(string(r)) {
		return int(r) - 97 + 1
	} else {
		return int(r) - (65 - 27)
	}
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	g := newGrid(len(lines), len(lines[0]))
	var start, end Node
	obstacleNodes := []Node{}
	weightedNodes := make([]Node, 0)
	possiblesStarts := make([]Node, 0)
	for y, line := range lines {
		//fmt.Println(line)

		for x, c := range line {
			n := 0
			if c == 'S' {
				start = Node{X: x, Y: y}
				n = RuneToInt('a')
				possiblesStarts = append(possiblesStarts, start)
			} else if c == 'E' {
				end = Node{X: x, Y: y}
				n = RuneToInt('z')
			} else {
				n = RuneToInt(c)

				if c == 'a' {
					possiblesStarts = append(possiblesStarts, Node{X: x, Y: y})
				}
			}
			g.content[y][x] = n
		}
	}

	//fmt.Println(start)
	//fmt.Println(end)
	//fmt.Println(g)

	// set nodes to the config
	aConfig := Config{
		GridWidth:     g.w,
		GridHeight:    g.h,
		InvalidNodes:  obstacleNodes,
		WeightedNodes: weightedNodes,
		grid:          g,
	}

	// create the algo with defined config
	algo, err := New(aConfig)
	if err != nil {
		fmt.Println("invalid astar config", err)
		return -1
	}

	if !complex {
		// run it
		foundPath, err := algo.FindPath(start, end)
		if err != nil || len(foundPath) == 0 {
			fmt.Println("No path found ...")
			return -1
		}

		// start is always visited but end should not appear as visited
		g.visited[start.Y][start.X] = true
		g.visited[end.Y][end.X] = false
		fmt.Println(g)
		fmt.Println(g.SumVisited())

		return len(foundPath)
	} else {

		best := 99999
		for _, s := range possiblesStarts {
			foundPath, err := algo.FindPath(s, end)
			if err == nil && len(foundPath) < best {
				best = len(foundPath)
			}
		}

		return best
	}
}
