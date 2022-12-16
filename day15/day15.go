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
	if yToWatch >= s.y-d && yToWatch <= s.y+d {
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
		x1 := s.x - d + delta
		x2 := s.x + d - delta
		if x1 > x2 {
			x2, x1 = x1, x2
		}
		newRange := &Range{
			Start: Point{y: y, x: x1},
			End:   Point{y: y, x: x2},
		}

		if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x >= r.Start.x && newRange.End.x <= r.End.x }) {
			return ranges
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool {
			return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
		}) {
			r := linq.From(ranges).WhereT(func(r *Range) bool {
				return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
			}).First().(*Range)
			r.End.x = newRange.End.x
			return ranges
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool {
			return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
		}) {
			r := linq.From(ranges).WhereT(func(r *Range) bool {
				return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
			}).First().(*Range)
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

func Process(fileName string, yToWatch, yMax int, isSample, complex bool) int {

	if !complex {
		nb, _ := ProcessInternal(fileName, yToWatch, isSample, complex)
		return nb
	} else {
		//nb, _ := ProcessInternal(fileName, 11, isSample, complex)
		//return nb

		for y := yToWatch; y <= yMax; y++ {
			_, x := ProcessInternal(fileName, y, isSample, complex)
			//fmt.Println(y, nb, x)
			if x != -1 {
				return x*4000000 + y
			}
		}
	}

	return 0
}

func ProcessInternal(fileName string, yToWatch int, isSample, complex bool) (int, int) {
	lines := common.ReadLinesFromFile(fileName)

	points := make([]Point, 0)
	sensors := make([]SensorAndBeacon, 0)
	beaconsXPosOnLine := make([]int, 0)
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

		if b.y == yToWatch && !linq.From(beaconsXPosOnLine).AnyWithT(func(x int) bool { return x == b.x }) {
			beaconsXPosOnLine = append(beaconsXPosOnLine, b.x)
		}
	}

	maxX := linq.From(points).SelectT(func(p Point) int { return p.x }).Max().(int)
	maxY := linq.From(points).SelectT(func(p Point) int { return p.y }).Max().(int)
	minX := linq.From(points).SelectT(func(p Point) int { return p.x }).Min().(int)
	minY := linq.From(points).SelectT(func(p Point) int { return p.y }).Min().(int)

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
		if isSample {
			for _, r := range ranges {
				for x := r.Start.x; x <= r.End.x; x++ {
					index := fmt.Sprintf("%d,%d", r.Start.y, x)
					if env[index] != 2 {
						env[index] = 3
					}
				}
			}
		}

		//PrintEnv(env, minY, maxY, minX, maxX)
		//fmt.Println()
	}

	if isSample {
		PrintEnv(env, minY, maxY, minX, maxX)
	}

	if complex {
		limitMinX, limitMaxX := 0, 20
		if !isSample {
			limitMinX, limitMaxX = 0, 4000000
		}
		minX = limitMinX
		maxX = limitMaxX
		for _, r := range ranges {

			if r.Start.x < limitMinX {
				r.Start.x = limitMinX
			}
			if r.End.x > limitMaxX {
				r.End.x = limitMaxX
			}
		}
	}

	for linq.From(ranges).AnyWithT(func(newRange *Range) bool {
		return linq.From(ranges).AnyWithT(func(r *Range) bool {
			return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
		}) ||
			linq.From(ranges).AnyWithT(func(r *Range) bool {
				return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
			}) ||
			linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x })
	}) {
		newRange := linq.From(ranges).FirstWithT(func(newRange *Range) bool {
			return linq.From(ranges).AnyWithT(func(r *Range) bool {
				return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
			}) ||
				linq.From(ranges).AnyWithT(func(r *Range) bool {
					return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
				}) ||
				linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x })
		}).(*Range)

		if linq.From(ranges).AnyWithT(func(r *Range) bool {
			return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
		}) {
			r := linq.From(ranges).WhereT(func(r *Range) bool {
				return newRange.Start.x >= r.Start.x && newRange.Start.x <= r.End.x && newRange.End.x > r.End.x
			}).First().(*Range)
			r.End.x = newRange.End.x
			i := IndexOf(ranges, newRange)
			ranges = append(ranges[0:i], ranges[i+1:]...)
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool {
			return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
		}) {
			r := linq.From(ranges).WhereT(func(r *Range) bool {
				return newRange.Start.x < r.Start.x && newRange.End.x <= r.End.x && newRange.End.x >= r.Start.x
			}).First().(*Range)
			r.Start.x = newRange.Start.x
			i := IndexOf(ranges, newRange)
			ranges = append(ranges[0:i], ranges[i+1:]...)
		} else if linq.From(ranges).AnyWithT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x }) {
			r := linq.From(ranges).WhereT(func(r *Range) bool { return newRange.Start.x < r.Start.x && newRange.End.x > r.End.x }).First().(*Range)
			r.Start.x = newRange.Start.x
			r.End.x = newRange.End.x
			i := IndexOf(ranges, newRange)
			ranges = append(ranges[0:i], ranges[i+1:]...)
		}
	}

	nb := 0
	for _, r := range ranges {
		nb += int(math.Abs(float64(r.Start.x)-float64(r.End.x))) + 1
	}

	if !complex {
		for _, bX := range beaconsXPosOnLine {
			if linq.From(ranges).AnyWithT(func(r *Range) bool { return r.Start.x <= bX && r.End.x >= bX }) {
				nb--
			}
		}
	}

	//fmt.Println(minX)
	//fmt.Println(maxX)

	if int(math.Abs(float64(maxX)-float64(minX)))-nb == 0 {

		linq.From(ranges).OrderByT(func(r *Range) int { return r.Start.x }).ToSlice(&ranges)

		fmt.Println("YEAH !")

		return nb, ranges[0].End.x + 1
	}

	return nb, -1
	//return linq.From(env).CountWithT(func(kv linq.KeyValue) bool {
	//	pattern := fmt.Sprintf("%d,", yToWatch)
	//	return kv.Key.(string)[0:len(pattern)] == pattern && kv.Value != 2
	//})
}

func IndexOf(list []*Range, toSearch *Range) int {
	for i, n := range list {
		if n == toSearch {
			return i
		}
	}
	return -1
}

func PrintEnv(env map[string]int, minY, maxY, minX, maxX int) {
	return

	for y := minY; y < maxY; y++ {
		if y < 10 {
			fmt.Printf(" ")
		}
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
					if y == 11 && x == 14 {
						fmt.Print("@")
					} else {
						fmt.Print("#")
					}
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
