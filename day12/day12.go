package day12

import (
	"bytes"
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"regexp"
)

type Grid struct {
	h       int
	w       int
	content [][]int
	visited [][]bool
}

type Point struct {
	x int
	y int
}

func newGrid(h int, w int) *Grid {
	g := Grid{h: h, w: w}
	g.content = make([][]int, h)
	g.visited = make([][]bool, h)
	for i := range g.content {
		g.content[i] = make([]int, w)
		g.visited[i] = make([]bool, w)
	}
	return &g
}

func (g *Grid) String() string {
	output := bytes.NewBufferString("")
	for i := range g.content {
		for j := range g.content[i] {
			if g.visited[i][j] {
				fmt.Fprint(output, "#")
			} else {
				//fmt.Fprint(output, string(rune(g.content[i][j]+96)))
				fmt.Fprint(output, ".")

			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grid) Content(p *Point) int {
	return g.content[p.x][p.y]
}

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
	for y, line := range lines {
		//fmt.Println(line)

		for x, c := range line {
			n := 0
			if c == 'S' {
				start = Node{X: x, Y: y}
				weightedNodes = append(weightedNodes, Node{X: x, Y: y, Weighting: RuneToInt('a')})
			} else if c == 'E' {
				end = Node{X: x, Y: y}
				weightedNodes = append(weightedNodes, Node{X: x, Y: y, Weighting: RuneToInt('z')})
			} else {
				n = RuneToInt(c)
				weightedNodes = append(weightedNodes, Node{X: x, Y: y, Weighting: n})
			}
			g.content[y][x] = n
		}
	}

	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(g)

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

	// run it
	foundPath, err := algo.FindPath(start, end)
	if err != nil || len(foundPath) == 0 {
		fmt.Println("No path found ...")
		return -1
	}

	// the foundPath has now the way to the target

	// IMPORTANT:
	// the path is in the opposite way so the endpoint node is on index 0
	// you can avoid it by switching the startNode<>endNode parameter
	for _, node := range foundPath {
		fmt.Println(node)
		g.visited[node.Y][node.X] = true
	}

	fmt.Println(g)

	return len(foundPath)
}
