package day9

import (
	"bytes"
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Grid struct {
	h       int
	w       int
	content [][]int
	knots   []*Point
}

type Point struct {
	x int
	y int
}

func newGrid(h int, w int) *Grid {
	g := Grid{h: h, w: w}
	g.content = make([][]int, h)
	for i := range g.content {
		g.content[i] = make([]int, w)
	}
	g.knots = make([]*Point, 0)
	return &g
}

func (g *Grid) String() string {
	output := bytes.NewBufferString("")
	for i := range g.content {
		for j := range g.content[i] {
			knotShown := false
			for k := 0; k < len(g.knots); k++ {
				if g.knots[k].y == i && g.knots[k].x == j {
					if k == 0 {
						fmt.Fprint(output, "H")
					} else {
						fmt.Fprint(output, k)
					}
					knotShown = true
					break
				}
			}
			if knotShown {
				continue
			} else if g.content[i][j] > 0 {
				fmt.Fprint(output, "#")
			} else {
				fmt.Fprint(output, ".")
			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grid) AjustKnotPosition(knotIndex int) {
	precedingKnot := g.knots[knotIndex-1]
	currentKnot := g.knots[knotIndex]

	headOnTheLeft := precedingKnot.x-currentKnot.x < 0
	headOnTheRight := precedingKnot.x-currentKnot.x > 0
	headOnUpside := precedingKnot.y-currentKnot.y < 0
	headOnDownside := precedingKnot.y-currentKnot.y > 0
	headFarOnTheLeft := precedingKnot.x-currentKnot.x < -1
	headFarOnTheRight := precedingKnot.x-currentKnot.x > 1
	headFarOnUpside := precedingKnot.y-currentKnot.y < -1
	headFarOnDownside := precedingKnot.y-currentKnot.y > 1

	if headFarOnTheLeft {
		currentKnot.x--
		if headOnDownside {
			currentKnot.y++
		} else if headOnUpside {
			currentKnot.y--
		}
	} else if headFarOnTheRight {
		currentKnot.x++
		if headOnDownside {
			currentKnot.y++
		} else if headOnUpside {
			currentKnot.y--
		}
	} else if headFarOnUpside {
		currentKnot.y--
		if headOnTheRight {
			currentKnot.x++
		} else if headOnTheLeft {
			currentKnot.x--
		}
	} else if headFarOnDownside {
		currentKnot.y++
		if headOnTheRight {
			currentKnot.x++
		} else if headOnTheLeft {
			currentKnot.x--
		}
	}
}

func (g *Grid) Sum() int {
	result := 0
	for i := range g.content {
		for j := range g.content[i] {
			result += g.content[i][j]
		}
	}
	return result
}

func Process(fileName string, numberOfKnots, h, w int) int {
	lines := common.ReadLinesFromFile(fileName)

	g := newGrid(h, w)
	for i := 0; i < numberOfKnots; i++ {
		g.knots = append(g.knots, &Point{x: g.w / 2, y: g.h / 2})
	}

	head := g.knots[0]
	tail := g.knots[numberOfKnots-1]
	g.content[tail.y][tail.x] = 1

	fmt.Println(g)

	for _, line := range lines {
		var direction string
		var move int
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "%s %d", &direction, &move)

		for i := 0; i < move; i++ {

			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y--
			case "D":
				head.y++
			}

			for i := 1; i < numberOfKnots; i++ {
				g.AjustKnotPosition(i)
			}

			// mark tail visited
			g.content[tail.y][tail.x] = 1

			// only print grid when size is small
			if h <= 100 {
				fmt.Println(g)
			}
		}

	}

	return g.Sum()
}
