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
	for y := range g.content {
		for x := range g.content[y] {
			if g.content[y][x] == 99 {
				fmt.Fprint(output, "#")
			} else if g.content[y][x] == 100 {
				fmt.Fprint(output, "o")
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

func (g *Grid) DropSandAt(startPos Point) bool {

	isInRange := func(y, x int) bool {
		return x >= 0 && y >= 0 && y < g.h && x < g.w
	}

	sand := startPos

	sandAtRest := false
	for !sandAtRest && (isInRange(sand.y+1, sand.x) || isInRange(sand.y+1, sand.x-1) || isInRange(sand.y+1, sand.x+1)) {
		if isInRange(sand.y+1, sand.x) && g.content[sand.y+1][sand.x] < 99 {
			sand.y++
		} else if isInRange(sand.y+1, sand.x-1) && g.content[sand.y+1][sand.x-1] < 99 {
			sand.y++
			sand.x--
		} else if isInRange(sand.y+1, sand.x+1) && g.content[sand.y+1][sand.x+1] < 99 {
			sand.y++
			sand.x++
		} else {
			g.content[sand.y][sand.x] = 100
			sandAtRest = true
		}
	}

	return sandAtRest // false if lost in infinite cave......
}

func (g *Grid) SumOfSand() int {
	result := 0
	for y := range g.content {
		for x := range g.content[y] {
			if g.content[y][x] == 100 {
				result++
			}
		}
	}
	return result
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

	allPoints = append(allPoints, Point{x: 500, y: 0})

	maxX := linq.From(allPoints).SelectT(func(p Point) int { return p.x }).Max().(int)
	maxY := linq.From(allPoints).SelectT(func(p Point) int { return p.y }).Max().(int)

	if complex {
		maxY += 2

		segments = append(segments, Segment{
			start: Point{y: maxY, x: 0},
			end:   Point{y: maxY, x: maxX + 499},
		})
	}

	g := newGrid(maxY+1, maxX+500)
	for _, s := range segments {
		g.drawSegment(s)
	}

	notLost := g.DropSandAt(Point{x: 500, y: 0})
	for notLost && g.content[0][500] != 100 {
		//fmt.Println(g)
		notLost = g.DropSandAt(Point{x: 500, y: 0})
	}

	//fmt.Println(g)

	return g.SumOfSand()
}
