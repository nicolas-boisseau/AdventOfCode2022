package day18

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Point struct {
	x, y, z int
}

type Cube struct {
	pos          Point
	exposedSides int
}

func newCube(pos Point) *Cube {
	return &Cube{pos, 6}
}

type Env struct {
	cubes map[string]*Cube
	minX  int
	maxX  int
	minY  int
	maxY  int
	minZ  int
	maxZ  int
}

func (env Env) CollideIfCubeExistsAt(pos Point) int {
	index := fmt.Sprintf("%d,%d,%d", pos.x, pos.y, pos.z)
	if e, exists := env.cubes[index]; exists {
		e.exposedSides--
		return -1
	}
	return 0
}

func (env Env) CubeExistsAt(pos Point) bool {
	index := fmt.Sprintf("%d,%d,%d", pos.x, pos.y, pos.z)
	if _, exists := env.cubes[index]; exists {
		return true
	}
	return false
}

func (env Env) InsertCube(cube *Cube) {

	cube.exposedSides =
		cube.exposedSides +
			env.CollideIfCubeExistsAt(Point{cube.pos.x + 1, cube.pos.y, cube.pos.z}) +
			env.CollideIfCubeExistsAt(Point{cube.pos.x - 1, cube.pos.y, cube.pos.z}) +
			env.CollideIfCubeExistsAt(Point{cube.pos.x, cube.pos.y + 1, cube.pos.z}) +
			env.CollideIfCubeExistsAt(Point{cube.pos.x, cube.pos.y - 1, cube.pos.z}) +
			env.CollideIfCubeExistsAt(Point{cube.pos.x, cube.pos.y, cube.pos.z + 1}) +
			env.CollideIfCubeExistsAt(Point{cube.pos.x, cube.pos.y, cube.pos.z - 1})

	index := fmt.Sprintf("%d,%d,%d", cube.pos.x, cube.pos.y, cube.pos.z)

	env.cubes[index] = cube
}

func (p Point) Index() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func (env Env) CanEscapeFrom(p Point, visited map[string]bool) bool {
	index := p.Index()
	visited[index] = true

	if !env.CubeExistsAt(p) &&
		(p.x <= env.minX || p.x >= env.maxX || p.y <= env.minY || p.y >= env.maxY || p.z <= env.minZ || p.z >= env.maxZ) {
		return true
	}

	upX := Point{p.x + 1, p.y, p.z}
	downX := Point{p.x - 1, p.y, p.z}
	upY := Point{p.x, p.y + 1, p.z}
	downY := Point{p.x, p.y - 1, p.z}
	upZ := Point{p.x, p.y, p.z + 1}
	downZ := Point{p.x, p.y, p.z - 1}

	if !visited[upX.Index()] && (!env.CubeExistsAt(upX) && env.CanEscapeFrom(upX, visited)) {
		return true
	}

	if !visited[downX.Index()] && (!env.CubeExistsAt(downX) && env.CanEscapeFrom(downX, visited)) {
		return true
	}

	if !visited[upY.Index()] && (!env.CubeExistsAt(upY) && env.CanEscapeFrom(upY, visited)) {
		return true
	}

	if !visited[downY.Index()] && (!env.CubeExistsAt(downY) && env.CanEscapeFrom(downY, visited)) {
		return true
	}

	if !visited[upZ.Index()] && (!env.CubeExistsAt(upZ) && env.CanEscapeFrom(upZ, visited)) {
		return true
	}

	if !visited[downZ.Index()] && (!env.CubeExistsAt(downZ) && env.CanEscapeFrom(downZ, visited)) {
		return true
	}

	return false
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	env := Env{
		cubes: make(map[string]*Cube, len(lines)),
	}

	for _, line := range lines {
		var x, y, z int
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "%d,%d,%d", &x, &y, &z)

		pos := Point{x, y, z}

		env.InsertCube(newCube(pos))
	}

	if !complex {

		numberOfSidesExposed := int(linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).exposedSides }).
			SumInts())

		return numberOfSidesExposed
	} else {
		//trappedAirAreas := 0
		env.minX = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.x }).
			Min().(int)
		env.maxX = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.x }).
			Max().(int)
		env.minY = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.y }).
			Min().(int)
		env.maxY = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.y }).
			Max().(int)
		env.minZ = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.z }).
			Min().(int)
		env.maxZ = linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).pos.z }).
			Max().(int)

		for x := env.minX; x <= env.maxX; x++ {
			for y := env.minY; y <= env.maxY; y++ {
				for z := env.minZ; z <= env.maxZ; z++ {

					refPoint := Point{x, y, z}
					visited := make(map[string]bool)

					if !env.CubeExistsAt(refPoint) && !env.CanEscapeFrom(refPoint, visited) {

						env.CollideIfCubeExistsAt(Point{x + 1, y, z})
						env.CollideIfCubeExistsAt(Point{x - 1, y, z})
						env.CollideIfCubeExistsAt(Point{x, y + 1, z})
						env.CollideIfCubeExistsAt(Point{x, y - 1, z})
						env.CollideIfCubeExistsAt(Point{x, y, z + 1})
						env.CollideIfCubeExistsAt(Point{x, y, z - 1})
					}
				}
			}
		}

		numberOfSidesExposed := int(linq.
			From(env.cubes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(*Cube).exposedSides }).
			SumInts())

		return numberOfSidesExposed
	}
}
