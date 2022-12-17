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

	falling := false
	rockAtRest := false
	for !rockAtRest && (!falling || isInRange(rock, -1, 0)) { // && (isInRange(rock, -1, 0)) || isInRange(rock, 0, -1) || isInRange(rock, 0, 1)

		if falling && isInRange(rock, -1, 0) {
			fmt.Println("fall 1")
			rock.pos.y--
		} else if !falling && env.jets[env.jetIndex] == '<' {
			//rock.pos.y--
			fmt.Println("try move left 1")
			if isInRange(rock, 0, -1) {
				fmt.Println("left 1")
				rock.pos.x--
			} else {
				fmt.Println("BLOCKED!")
			}
		} else if !falling && env.jets[env.jetIndex] == '>' {
			//rock.pos.y--
			fmt.Println("try move right 1")
			if isInRange(rock, 0, 1) {
				fmt.Println("right 1")
				rock.pos.x++
			} else {
				fmt.Println("BLOCKED!")
			}
		} else {
			rockAtRest = true
		}

		if !falling {
			fmt.Printf("incrementing jetIndex (before incrementing: %d)\n", env.jetIndex)
			env.jetIndex++
		}

		falling = !falling

		env.PrintEnv()

		fmt.Println("next...")
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

func (env *Env) CreateRock() *Rock {
	r := &Rock{
		points: []*Point{
			&Point{0, 0},
			&Point{1, 0},
			&Point{2, 0},
			&Point{3, 0},
		},
		pos: &Point{3, env.maxY - 1},
	}
	return r
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	e := &Env{
		env:  make(map[string]int),
		minX: 0,
		maxX: 10,
		minY: 0,
		maxY: 5,
		jets: lines[0],
	}

	e.DropRock(e.CreateRock())

	e.PrintEnv()

	return len(lines)
}

func (e *Env) PrintEnv() {
	for y := e.maxY - 1; y >= e.minY; y-- {

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
