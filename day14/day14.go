package day14

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Grid struct {
	h       int
	w       int
	content [][]int
}
type Point struct {
	x int
	y int
}
type Segment struct {
	start Point
	end   Point
}

func newGrid(h int, w int) *Grid {
	g := Grid{h: h, w: w}
	g.content = make([][]int, h)
	for i := range g.content {
		g.content[i] = make([]int, w)
	}
	return &g
}

func (g *Grid) String() string {
	output := bytes.NewBufferString("")
	for i := range g.content {
		for j := range g.content[i] {
			if g.content[i][j] == 99 {
				fmt.Fprint(output, "#")
			} else {
				fmt.Fprint(output, ".")
			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grid) drawSegment(s Segment) {
	if s.start.x == s.end.x {
		inc := 1
		if s.start.y > s.end.y {
			inc = -1
		}
		for y := s.start.y; y != s.end.y+inc; y += inc {
			g.content[y][s.start.x] = 99
		}
	} else if s.start.y == s.end.y {
		inc := 1
		if s.start.x > s.end.x {
			inc = -1
		}
		for x := s.start.x; x != s.end.x+inc; x += inc {
			g.content[s.start.y][x] = 99
		}
	} else {
	}
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	segments := make([]Segment, 0)
	allPoints := make([]Point, 0)
	for _, line := range lines {
		//fmt.Println(line)
		splitted := strings.Split(line, " -> ")
		var previousPoint *Point
		previousPoint = nil
		for _, coord := range splitted {
			var x, y int
			reader := strings.NewReader(coord)
			fmt.Fscanf(reader, "%d,%d", &x, &y)
			newPoint := &Point{x: x, y: y}
			allPoints = append(allPoints, *newPoint)
			if previousPoint != nil {
				segments = append(segments, Segment{
					start: *previousPoint,
					end:   *newPoint,
				})
			}
			previousPoint = newPoint
		}
	}

	maxX := linq.From(allPoints).SelectT(func(p Point) int { return p.x }).Max().(int)
	maxY := linq.From(allPoints).SelectT(func(p Point) int { return p.y }).Max().(int)

	g := newGrid(maxY+1, maxX+1)
	for _, s := range segments {
		g.drawSegment(s)
	}

	fmt.Println(g)

	return 0
}
