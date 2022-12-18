package day17

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
)

type Point struct {
	x int
	y int
}

//func (g *Grid) DropSandAt(startPos Point) bool {
//
//	isInRange := func(y, x int) bool {
//		return x >= 0 && y >= 0 && y < g.h && x < g.w
//	}
//
//	sand := startPos
//
//	sandAtRest := false
//	for !sandAtRest && (isInRange(rock.pos.y+1, rock.pos.x) || isInRange(rock.pos.y+1, rock.pos.x-1) || isInRange(rock.pos.y+1, rock.pos.x+1)) {
//		if isInRange(rock.pos.y+1, rock.pos.x) && g.content[rock.pos.y+1][rock.pos.x] < 99 {
//			rock.pos.y++
//		} else if isInRange(rock.pos.y+1, rock.pos.x-1) && g.content[rock.pos.y+1][rock.pos.x-1] < 99 {
//			rock.pos.y++
//			rock.pos.x--
//		} else if isInRange(rock.pos.y+1, rock.pos.x+1) && g.content[rock.pos.y+1][rock.pos.x+1] < 99 {
//			rock.pos.y++
//			rock.pos.x++
//		} else {
//			g.content[rock.pos.y][rock.pos.x] = 100
//			sandAtRest = true
//		}
//	}
//
//	return sandAtRest // false if lost in infinite cave......
//}

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
}

func (env *Env) DropRock(rock *Rock) {
	isInRange := func(r *Rock, dY, dX int) bool {
		newY := r.pos.y + dY
		newX := r.pos.x + dX
		inSpace := newX >= 1 && newY >= 1 && newY < env.maxY && newX < env.maxX-2

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

	env.rocks = append(env.rocks, rock)

	//fmt.Println("Starting new block fall")
	env.PrintEnv()

	falling := false
	rockAtRest := false
	for !rockAtRest && (!falling || isInRange(rock, -1, 0)) { // && (isInRange(rock, -1, 0)) || isInRange(rock, 0, -1) || isInRange(rock, 0, 1)

		if falling && isInRange(rock, -1, 0) {
			//fmt.Println("fall 1")
			rock.pos.y--
		} else if !falling && env.jets[env.jetIndex] == '<' {
			//rock.pos.y--
			//fmt.Println("try move left 1")
			if isInRange(rock, 0, -1) {
				//fmt.Println("left 1")
				rock.pos.x--
			} else {
				//fmt.Println("BLOCKED!")
			}
		} else if !falling && env.jets[env.jetIndex] == '>' {
			//rock.pos.y--
			//fmt.Println("try move right 1")
			if isInRange(rock, 0, 1) {
				//fmt.Println("right 1")
				rock.pos.x++
			} else {
				//fmt.Println("BLOCKED!")
			}
		} else {
			rockAtRest = true
		}

		if !falling {
			//fmt.Printf("incrementing jetIndex (before incrementing: %d)\n", env.jetIndex)
			env.jetIndex = (env.jetIndex + 1) % len(env.jets)
		}

		falling = !falling

		env.PrintEnv()

		//fmt.Println("next...")
	}

	return
	//return rockAtRest // false if lost in infinite cave......

}

func (env *Env) RockIsBlocked(r *Rock) bool {

	// check walls
	for _, p := range r.points {
		rX1 := r.pos.x + p.x
		rY1 := r.pos.y + p.y

		if rX1 > env.maxX-2 || rX1 < env.minX+1 || rY1 < 1 {
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

	if maxY == 0 {
		return 0 // start...
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

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	e := &Env{
		env:   make(map[string]int),
		minX:  0,
		maxX:  9,
		minY:  0,
		maxY:  1,
		jets:  lines[0],
		debug: false,
		rocks: make([]*Rock, 0),
	}

	funcs := make([]func() *Rock, 0)
	funcs = append(funcs, func() *Rock {
		return e.CreateLineRock()
	})
	funcs = append(funcs, func() *Rock {
		return e.CreateCrossRock()
	})
	funcs = append(funcs, func() *Rock {
		return e.CreateLRock()
	})
	funcs = append(funcs, func() *Rock {
		return e.CreateTowerRock()
	})
	funcs = append(funcs, func() *Rock {
		return e.CreateSquareRock()
	})

	towerHeight := 0
	for i := 1; i <= 2022; i++ {
		e.DropRock(funcs[(i-1)%len(funcs)]())

		towerHeight = e.MaxRocksY()
		if i > 2015 {
			fmt.Println(i, "=", towerHeight)
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
