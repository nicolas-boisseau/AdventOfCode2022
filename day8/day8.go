package day8

import (
	"bytes"
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
)

type Grid struct {
	h       int
	w       int
	content [][]int
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
			fmt.Fprint(output, g.content[i][j])
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func (g *Grid) CheckIfTreeVisibleFromDirectionX(x, y, directionX int) (bool, int) {
	i := x
	j := y
	referenceTree := g.content[i][j]
	i += directionX
	viewingDistance := 0
	for i >= 0 && i < g.w {
		viewingDistance++
		if g.content[i][j] >= referenceTree {
			return false, viewingDistance
		}

		i += directionX
	}

	return true, viewingDistance
}

func (g *Grid) CheckIfTreeVisibleFromDirectionY(x, y, directionY int) (bool, int) {
	i := x
	j := y
	referenceTree := g.content[i][j]
	j += directionY
	viewingDistance := 0
	for j >= 0 && j < g.w {
		viewingDistance++
		if g.content[i][j] >= referenceTree {
			return false, viewingDistance
		}

		j += directionY
	}

	return true, viewingDistance
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	g := newGrid(len(lines), len(lines[0]))
	for i, line := range lines {
		//fmt.Println(line)

		for j, c := range line {
			n, _ := strconv.Atoi(string(c))
			g.content[i][j] = n
		}
	}

	visibleTrees := 0
	betterView := 0
	for i, _ := range g.content {
		for j, _ := range g.content[i] {
			if i == 0 || i == len(g.content)-1 || j == 0 || j == len(g.content[i])-1 {
				visibleTrees++
			} else {
				isVisibleRight, viewRight := g.CheckIfTreeVisibleFromDirectionX(i, j, 1)
				isVisibleLeft, viewLeft := g.CheckIfTreeVisibleFromDirectionX(i, j, -1)
				isVisibleBottom, viewBottom := g.CheckIfTreeVisibleFromDirectionY(i, j, 1)
				isVisibleTop, viewTop := g.CheckIfTreeVisibleFromDirectionY(i, j, -1)

				if isVisibleRight || isVisibleLeft || isVisibleTop || isVisibleBottom {
					visibleTrees++
				}

				//fmt.Println("content[", i, ",", j, "]=", g.content[i][j], ", view=", viewRight*viewLeft*viewTop*viewBottom)
				view := viewRight * viewLeft * viewTop * viewBottom
				if view > betterView {
					betterView = view
				}
			}
		}
	}

	if !complex {
		return visibleTrees
	} else {
		return betterView
	}
}
