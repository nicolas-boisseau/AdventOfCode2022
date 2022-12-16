package day16

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Valve struct {
	name          string
	rate          int
	NextValvesStr []string
	NextValves    []*Valve
	isOpened      bool
}

func (v *Valve) String() string {
	output := bytes.NewBufferString("")
	fmt.Fprintln(output, "Valve : ", v.name, "at rate", v.rate, " leading to :")
	for _, otherV := range v.NextValves {
		fmt.Fprint(output, otherV.name, ",")
	}
	fmt.Fprintln(output)
	return output.String()
}

func GetValve(valves []*Valve, name string) *Valve {
	return linq.From(valves).FirstWithT(func(vv *Valve) bool { return vv.name == name }).(*Valve)
}

func ExtractValves(lines []string) []*Valve {
	valves := make([]*Valve, 0)
	for _, line := range lines {

		var valve string
		var rate int
		reader := strings.NewReader(line)
		fmt.Fscanf(reader, "Valve %s has flow rate=%d; tunnels lead to valves", &valve, &rate)

		ind := strings.Index(line, "valves")
		leadToValvesStr := ""
		if ind != -1 {
			leadToValvesStr = line[ind+len("valves")+1:]
		} else {
			ind = strings.Index(line, "valve")
			leadToValvesStr = line[ind+len("valve")+1:]
		}
		splittedLeadToValves := strings.Split(leadToValvesStr, ", ")

		//fmt.Println(valve, rate, ":")
		//fmt.Println(splittedLeadToValves)

		valves = append(valves, &Valve{
			name:          valve,
			rate:          rate,
			NextValvesStr: splittedLeadToValves,
			NextValves:    make([]*Valve, 0),
		})
	}

	// evaluate valves
	for _, v := range valves {
		for _, otherStrValve := range v.NextValvesStr {
			otherV := GetValve(valves, otherStrValve)
			v.NextValves = append(v.NextValves, otherV)
		}
	}
	return valves
}

func FindBestNextValve(nextValves []*Valve, visited map[*Valve]bool) (*Valve, bool) {

	var bestNextValve *Valve
	if linq.From(nextValves).
		AnyWithT(func(valve *Valve) bool { return !valve.isOpened }) {
		bestNextValve = linq.From(nextValves).
			WhereT(func(valve *Valve) bool { return !valve.isOpened }).
			OrderByDescendingT(func(valve *Valve) int { return valve.rate }).First().(*Valve)
	}

	var nextBestValve *Valve
	var nextNextBestValve *Valve
	if bestNextValve == nil {
		// must find a nextValve which can access to a better pressure
		for _, nextV := range nextValves {
			if !visited[nextV] {
				visited[nextV] = true
				bestNextValve, _ = FindBestNextValve(nextV.NextValves, visited)

				if nextNextBestValve == nil || (bestNextValve != nil && bestNextValve.rate > nextNextBestValve.rate) {
					nextNextBestValve = bestNextValve
					nextBestValve = nextV
				}
			}
		}

		return nextBestValve, true
	}

	return bestNextValve, true
}

func FindBestPath(valves []*Valve, v1, v2 string) int {

	// set nodes to the config
	aConfig := Config{
		GridWidth:     len(valves),
		GridHeight:    len(valves),
		InvalidNodes:  []Node{},
		WeightedNodes: []Node{},
	}

	// create the algo with defined config
	algo, err := New(aConfig)
	if err != nil {
		fmt.Println("invalid astar config", err)
		return -1
	}

	// run it
	foundPath, err := algo.FindPath(v1, v2)
	if err != nil || len(foundPath) == 0 {
		fmt.Println("No path found ...")
		return -1
	}

	return len(foundPath)
}

func CurrentPressure(valves []*Valve) int {
	return int(linq.From(valves).
		WhereT(func(valve *Valve) bool { return valve.isOpened }).
		SelectT(func(valve *Valve) int { return valve.rate }).SumInts())
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	valves := ExtractValves(lines)

	currentValve := GetValve(valves, "AA")
	currentValve.isOpened = true

	totalPressure := 0
	for minute := 1; minute <= 30; minute++ {

		fmt.Println("== Minute", minute, "==")

		currentPressure := CurrentPressure(valves)
		totalPressure += currentPressure

		if !currentValve.isOpened {
			fmt.Println("Open valve", currentValve.name)
			currentValve.isOpened = true
		} else {

			visited := make(map[*Valve]bool)
			bestNextValve, needMove := FindBestNextValve(currentValve.NextValves, visited)
			if needMove {
				fmt.Println("Move to valve", bestNextValve.name)
				currentValve = bestNextValve
			}
		}

		fmt.Println("Current valve:", currentValve.name, ", release", currentPressure, "pressure")

		fmt.Println()
	}

	return len(lines)
}
