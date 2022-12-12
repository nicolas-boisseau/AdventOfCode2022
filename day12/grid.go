package day12

import (
	"bytes"
	"fmt"
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

func (g *Grid) SumVisited() int {
	result := 0
	for y := range g.visited {
		for x := range g.visited[y] {
			if g.visited[y][x] {
				result++
			}
		}
	}
	return result
}
