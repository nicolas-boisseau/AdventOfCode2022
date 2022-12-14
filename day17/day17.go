package day17

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

type Point struct {
	x int
	y int
}

type Rock struct {
	points []*Point
	pos    *Point
}

type Env struct {
	minX, maxX, minY, maxY int
	env                    map[string]int
	rocks                  []*Rock
	jets                   string
	jetIndex               int
	debug                  bool
	floorY                 int
}

func (env *Env) DropRock(rock *Rock) {
	isInRange := func(r *Rock, dY, dX int) bool {
		newY := r.pos.y + dY
		newX := r.pos.x + dX
		inSpace := newX > env.minX && newY > env.floorY && newX < env.maxX-1

		if !inSpace {
			return false
		}

		r.pos.y += dY
		r.pos.x += dX
		rockIsBlocked := env.RockIsBlocked(r)
		r.pos.y -= dY
		r.pos.x -= dX

		return !rockIsBlocked
	}

	fall := func() bool {
		if isInRange(rock, -1, 0) {
			//fmt.Println("fall 1")
			rock.pos.y--
			return false
		}

		return true
	}

	move := func() {
		jet := env.jets[env.jetIndex]
		env.jetIndex = (env.jetIndex + 1) % len(env.jets)

		if jet == '<' {
			//fmt.Println("try move left 1")
			if isInRange(rock, 0, -1) {
				//fmt.Println("left 1")
				rock.pos.x--
			} else {
				//fmt.Println("BLOCKED!")
			}
		} else if jet == '>' {
			//fmt.Println("try move right 1")
			if isInRange(rock, 0, 1) {
				//fmt.Println("right 1")
				rock.pos.x++
			} else {
				//fmt.Println("BLOCKED!")
			}
		}
	}

	env.rocks = append(env.rocks, rock)

	//fmt.Println("Starting new block fall")
	env.PrintEnv()

	rockAtRest := false
	for !rockAtRest { // && (isInRange(rock, -1, 0)) || isInRange(rock, 0, -1) || isInRange(rock, 0, 1)

		move()
		rockAtRest = fall()

		env.PrintEnv()

		//fmt.Println("next...")
	}
}

func (env *Env) RockIsBlocked(r *Rock) bool {

	// check walls
	for _, p := range r.points {
		rX1 := r.pos.x + p.x
		rY1 := r.pos.y + p.y

		if (rX1 >= env.maxX-1 || rX1 <= env.minX) || (rY1 <= env.floorY) {
			return true
		}
	}

	for _, otherRock := range env.rocks {
		if otherRock != r {

			for _, p1 := range r.points {
				rY1 := r.pos.y + p1.y
				rX1 := r.pos.x + p1.x
				for _, p2 := range otherRock.points {
					rY2 := otherRock.pos.y + p2.y
					rX2 := otherRock.pos.x + p2.x

					if rY1 == rY2 && rX1 == rX2 {
						// collision !
						return true
					}
				}
			}
		}
	}

	return false
}

func (env *Env) MaxRocksY() int {
	maxY := 0
	for _, r := range env.rocks {
		for _, p := range r.points {
			rY1 := r.pos.y + p.y

			if rY1 > maxY {
				maxY = rY1
			}
		}
	}

	return maxY
}

func (env *Env) CreateLineRock() *Rock {
	newMaxY := env.MaxRocksY() + 4
	env.maxY = newMaxY + 1

	r := &Rock{
		points: []*Point{
			&Point{0, 0},
			&Point{1, 0},
			&Point{2, 0},
			&Point{3, 0},
		},
		pos: &Point{3, newMaxY},
	}
	return r
}

func (env *Env) CreateCrossRock() *Rock {
	newMaxY := env.MaxRocksY() + 4
	env.maxY = newMaxY + 3

	r := &Rock{
		points: []*Point{
			&Point{0, 1},
			&Point{1, 0},
			&Point{1, 1},
			&Point{1, 2},
			&Point{2, 1},
		},
		pos: &Point{3, newMaxY},
	}
	return r
}

func (env *Env) CreateLRock() *Rock {
	newMaxY := env.MaxRocksY() + 4
	env.maxY = newMaxY + 3

	r := &Rock{
		points: []*Point{
			&Point{0, 0},
			&Point{1, 0},
			&Point{2, 0},
			&Point{2, 1},
			&Point{2, 2},
		},
		pos: &Point{3, newMaxY},
	}
	return r
}

func (env *Env) CreateTowerRock() *Rock {
	newMaxY := env.MaxRocksY() + 4
	env.maxY = newMaxY + 4

	r := &Rock{
		points: []*Point{
			&Point{0, 0},
			&Point{0, 1},
			&Point{0, 2},
			&Point{0, 3},
		},
		pos: &Point{3, newMaxY},
	}
	return r
}

func (env *Env) CreateSquareRock() *Rock {
	newMaxY := env.MaxRocksY() + 4
	env.maxY = newMaxY + 2

	r := &Rock{
		points: []*Point{
			&Point{0, 0},
			&Point{0, 1},
			&Point{1, 0},
			&Point{1, 1},
		},
		pos: &Point{3, newMaxY},
	}
	return r
}

func (e *Env) PatternAtPos(y, x, h, w int) [][]bool {
	pattern := make([][]bool, h)
	for yy := h - 1; yy >= 0; yy-- {
		pattern[yy] = make([]bool, w)
		for xx := x; xx < w; xx++ {
			if e.IsAnyRock(y+yy, xx) {
				pattern[yy][xx] = true
			}
		}
	}
	return pattern
}

func IsSame(pattern1, pattern2 [][]bool) bool {
	for y := 0; y < len(pattern1); y++ {
		for x := 0; x < len(pattern1[y]); x++ {
			if pattern1[y][x] != pattern2[y][x] {
				return false
			}
		}
	}
	return true
}

func Process(fileName string, complex bool, debug bool, targetLevel int) int {
	lines := common.ReadLinesFromFile(fileName)

	e := &Env{
		env:    make(map[string]int, 1000),
		minX:   0,
		maxX:   9,
		minY:   0,
		maxY:   0,
		floorY: 0,
		jets:   lines[0],
		debug:  debug,
		rocks:  make([]*Rock, 0),
	}

	rockFactory := make([]func() *Rock, 0)
	rockFactory = append(rockFactory, func() *Rock { return e.CreateLineRock() })
	rockFactory = append(rockFactory, func() *Rock { return e.CreateCrossRock() })
	rockFactory = append(rockFactory, func() *Rock { return e.CreateLRock() })
	rockFactory = append(rockFactory, func() *Rock { return e.CreateTowerRock() })
	rockFactory = append(rockFactory, func() *Rock { return e.CreateSquareRock() })

	towerHeight := 0
	patternTowerHeight := 0
	var patternToSearch [][]bool
	incrementByLevel := make(map[int]int)
	for i := 1; i <= targetLevel; i++ {
		e.DropRock(rockFactory[(i-1)%len(rockFactory)]())

		previousTowerHeight := towerHeight
		towerHeight = e.MaxRocksY()
		if i == targetLevel {
			fmt.Println(i, "=", towerHeight)
		}

		if i == 1000 {
			fmt.Println("searching pattern")
			// recherche d'un pattern qui se r??p??te
			patternToSearch = e.PatternAtPos(towerHeight-15, 1, 15, e.maxX)
			patternTowerHeight = towerHeight
		}

		if i >= 1001 {

			incrementByLevel[i-1000] = towerHeight - previousTowerHeight

			patternToSearch2 := e.PatternAtPos(towerHeight-15, 1, 15, e.maxX)
			if IsSame(patternToSearch, patternToSearch2) {
				fmt.Println("PATTERN FOUND !!")
				i_diff := i - 1000
				fmt.Println("i =", i, ", DIFF = ", i_diff)
				fmt.Println("Tower height:", towerHeight, ", DIFF = ", towerHeight-patternTowerHeight)

				fmt.Println("Increment by level :", incrementByLevel)
				totalIncrement := linq.From(incrementByLevel).SelectT(func(kv linq.KeyValue) int { return kv.Value.(int) }).SumInts()
				fmt.Println("TOTAL increment = ", totalIncrement)

				j := i
				for j+i_diff < targetLevel {
					j += i_diff
					towerHeight += int(totalIncrement)
				}
				// reste juste quelques niveaux!
				lastLevel := 1
				for j < targetLevel {
					j++
					towerHeight += incrementByLevel[lastLevel]
					lastLevel++
				}

				// ??a y est on y est !
				fmt.Println(j)
				fmt.Println(towerHeight)

				return towerHeight
			}
		}
	}

	e.PrintEnv()

	return towerHeight
}

func (e *Env) PrintEnv() {
	if !e.debug {
		return
	}

	for y := e.maxY - 1; y >= e.minY; y-- {
		if y < 10 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d : ", y)

		for x := e.minX; x < e.maxX; x++ {

			if e.IsAnyRock(y, x) {
				fmt.Print("@")
				continue
			}

			if x == 0 || x == e.maxX-1 {
				fmt.Print("|")
				continue
			} else if y == 0 {
				fmt.Print("_")
				continue
			}

			index := fmt.Sprintf("%d,%d", y, x)
			if e, exists := e.env[index]; exists {
				if e == 1 {
					fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (env *Env) IsAnyRock(y int, x int) bool {
	for _, r := range env.rocks {
		for _, p := range r.points {
			if r.pos.x+p.x == x && r.pos.y+p.y == y {
				return true
			}
		}
	}
	return false
}
