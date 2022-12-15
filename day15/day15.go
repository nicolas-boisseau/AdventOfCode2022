package day15

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"math"
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
type SensorAndBeacon struct {
	S Point
	B Point
}

type Range struct {
	Start Point
	End   Point
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
			if g.content[y][x] == 1 {
				fmt.Fprint(output, "S")
			} else if g.content[y][x] == 2 {
				fmt.Fprint(output, "B")
			} else if g.content[y][x] == 3 {
				fmt.Fprint(output, "#")
			} else {
				fmt.Fprint(output, ".")
			}
		}

		fmt.Fprintln(output)

	}
	return output.String()
}

func fillSensorView(env map[string]int, s Point, distance float64, yToWatch int, ranges []*Range) []*Range {
	d := int(distance)
	//for y := s.y - d; y <= s.y+d && y <= yToWatch; y++ {
	if yToWatch >= s.y-d || yToWatch <= s.y+d {
		y := yToWatch

		//for x := -s.x - d; x <= s.x+d; x++ {
		//	curPoint := Point{y: y, x: x}
		//	if s.distance(curPoint) <= float64(d) && y == yToWatch {
		//		index := fmt.Sprintf("%d,%d", y, x)
		//		//fmt.Println(index)
		//		if _, exists := env[index]; !exists {
		//			env[index] = 3
		//		}
		//	}
		//}
		delta := int(math.Abs(float64(s.y) - float64(y)))
		newRange := &Range{
			Start: Point{y: y, x: s.x - d + delta},
			End:   Point{y: y, x: s.x + d - delta},
		}

		if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x >= r.Start.x && newRange.End.x <= r.End.x }) {
			return ranges
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x >= r.Start.x && newRange.End.x > r.End.x }) {
			r := linq.From(ranges).WhereT(func(r *Range) bool { return newRange.Start.x >= r.Start.x && newRange.End.x > r.End.x }).First().(*Range)
			r.End.x = newRange.End.x
			return ranges
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x }) {
			r := linq.From(ranges).WhereT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x }).First().(*Range)
			r.Start.x = newRange.Start.x
			return ranges
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x }) {
			r := linq.From(ranges).WhereT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x }).First().(*Range)
			r.Start.x = newRange.Start.x
			r.End.x = newRange.End.x
			return ranges
		} else {
			ranges = append(ranges, newRange)
		}
	}
	//}
	//return Range{
	//	Start: Point{y: s.y, x: s.x},
	//	End:   Point{y: s.y, x: s.x},
	//}
	return ranges
}

func (p Point) distance(p2 Point) float64 {
	//return math.Sqrt(math.Pow(float64(x2) - float64(x1), 2.0) + math.Pow(float64(y2) - float64(y1), 2.0))
	return (math.Abs(float64(p2.x)-float64(p.x)) + math.Abs(float64(p2.y)-float64(p.y)))
}

func Process(fileName string, yToWatch, yMax int, complex bool) int {

	if !complex {
		nb, _ := ProcessInternal(fileName, yToWatch, false)
		return nb
	} else {
		for y := yToWatch; y <= yMax; y++ {
			nb, spaceLeft := ProcessInternal(fileName, y, true)
			fmt.Println(y, nb, spaceLeft)
		}
	}

	return 0
}

func ProcessInternal(fileName string, yToWatch int, complex bool) (int, int) {
	lines := common.ReadLinesFromFile(fileName)

	points := make([]Point, 0)
	sensors := make([]SensorAndBeacon, 0)
	for _, line := range lines {
		var sX, sY, bX, bY int
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)

		s := Point{x: sX, y: sY}
		b := Point{x: bX, y: bY}
		sb := SensorAndBeacon{S: s, B: b}

		sensors = append(sensors, sb)
		points = append(points, s)
		points = append(points, b)
	}

	maxX := linq.From(points).SelectT(func(p Point) int { return p.x }).Max().(int)
	//maxY := linq.From(points).SelectT(func(p Point) int { return p.y }).Max().(int)
	minX := linq.From(points).SelectT(func(p Point) int { return p.x }).Min().(int)
	//minY := linq.From(points).SelectT(func(p Point) int { return p.y }).Min().(int)

	//env := set.New()
	env := make(map[string]int)
	//for y := minY; y < maxY; y++ {
	//	for x := minX; x < maxX; x++ {
	//		index := fmt.Sprintf("%d,%d", y, x)
	//		env[index] = 0
	//	}
	//}

	for _, s := range sensors {

		sIndex := fmt.Sprintf("%d,%d", s.S.y, s.S.x)
		bIndex := fmt.Sprintf("%d,%d", s.B.y, s.B.x)

		env[sIndex] = 1
		env[bIndex] = 2

		//fmt.Println(g)
	}

	//PrintEnv(env, minY, maxY, minX, maxX)

	ranges := make([]*Range, 0)
	for _, s := range sensors {

		ranges = fillSensorView(env, s.S, s.S.distance(s.B), yToWatch, ranges)

		//ranges = append(ranges, r)

		// for DEBUG ONLY
		//for _, r := range ranges {
		//	for x := r.Start.x; x <= r.End.x; x++ {
		//		index := fmt.Sprintf("%d,%d", r.Start.y, x)
		//		env[index] = 3
		//	}
		//}

		//PrintEnv(env, minY, maxY, minX, maxX)
		fmt.Println()
	}

	//PrintEnv(env, minX, maxX, minY, maxY)

	nb := 0
	for _, r := range ranges {
		nb += int(math.Abs(float64(r.Start.x) - float64(r.End.x)))
	}

	return nb, int(math.Abs(float64(maxX)-float64(minX))) - nb
	//return linq.From(env).CountWithT(func(kv linq.KeyValue) bool {
	//	pattern := fmt.Sprintf("%d,", yToWatch)
	//	return kv.Key.(string)[0:len(pattern)] == pattern && kv.Value != 2
	//})
}

func PrintEnv(env map[string]int, minY, maxY, minX, maxX int) {
	for y := minY; y < maxY; y++ {
		fmt.Printf("%d : ", y)

		for x := minX; x < maxX; x++ {

			index := fmt.Sprintf("%d,%d", y, x)
			if e, exists := env[index]; exists {
				switch e {
				case 1:
					fmt.Print("S")
				case 2:
					fmt.Print("B")
				case 3:
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
