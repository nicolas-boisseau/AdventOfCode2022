package day16

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

//type Valve struct {
//	name          string
//	rate          int
//	NextValvesStr []string
//	NextValves    []*Node
//	isOpened      bool
//}

func (v *Node) String() string {
	output := bytes.NewBufferString("")
	fmt.Fprintln(output, "Valve : ", v.name, "at rate", v.rate, " leading to :")
	for _, otherV := range v.NextValves {
		fmt.Fprint(output, otherV.name, ",")
	}
	fmt.Fprintln(output)
	return output.String()
}

func GetValve(valves []*Node, name string) *Node {
	return linq.From(valves).FirstWithT(func(vv *Node) bool { return vv.name == name }).(*Node)
}

func ExtractValves(lines []string) []*Node {
	valves := make([]*Node, 0)
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

		valves = append(valves, &Node{
			name:          valve,
			rate:          rate,
			NextValvesStr: splittedLeadToValves,
			NextValves:    make([]*Node, 0),
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

func FindBestPath(valves []*Node, v1Name, v2Name string) []*Node {

	for _, v := range valves {
		v.parent = nil
	}

	v1 := linq.From(valves).FirstWithT(func(node *Node) bool { return node.name == v1Name }).(*Node)
	v2 := linq.From(valves).FirstWithT(func(node *Node) bool { return node.name == v2Name }).(*Node)

	// set nodes to the config
	aConfig := Config{
		GridWidth:     len(valves),
		GridHeight:    len(valves),
		InvalidNodes:  []*Node{},
		WeightedNodes: []*Node{},
	}

	// create the algo with defined config
	algo, err := New(aConfig)
	if err != nil {
		fmt.Println("invalid astar config", err)
		return []*Node{}
	}

	// run it
	foundPath, err := algo.FindPath(v1, v2)
	if err != nil || len(foundPath) == 0 {
		fmt.Println("No path found ...")
		return []*Node{}
	}

	//fmt.Println("FOUND PATH :")
	orderedPath := make([]*Node, 0)
	for i := len(foundPath) - 1; i >= 0; i-- {
		//fmt.Println(foundPath[i].name, " -> ")
		orderedPath = append(orderedPath, foundPath[i])
	}

	return orderedPath
}

func FindBestNextMove(allValves []*Node, node *Node) *Node {

	pathTo := make(map[string][]*Node)
	for _, v := range allValves {
		if v == node {
			continue
		}
		path := FindBestPath(allValves, node.name, v.name)
		if indexOf(node.NextValvesStr, path[0].name) != -1 {
			pathTo[v.name] = path
		}
	}

	var nextBests []*Node
	linq.From(allValves).
		WhereT(func(n *Node) bool {
			return !n.isOpened &&
				n.name != node.name
		}).
		OrderByDescendingT(func(n *Node) int {
			totalPotential := n.rate - len(pathTo[n.name])

			for _, intermediate := range pathTo[n.name] {
				if !intermediate.isOpened && intermediate != n {
					totalPotential += intermediate.rate
				}
			}

			return totalPotential
		}).
		ToSlice(&nextBests)

	if len(nextBests) > 0 {
		return pathTo[nextBests[0].name][0]
	} else {
		return nil
	}
}

func CurrentPressure(valves []*Node) int {
	return int(linq.From(valves).
		WhereT(func(valve *Node) bool { return valve.isOpened }).
		SelectT(func(valve *Node) int { return valve.rate }).SumInts())
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	valves := ExtractValves(lines)

	pathTo := make(map[string][]*Node)
	for _, v := range valves {
		for _, v2 := range valves {
			if v == v2 {
				continue
			}
			pathTo[v2.name] = FindBestPath(valves, v.name, v2.name)
		}

		v.pathTo = pathTo
	}

	//FindBestPath(valves, "DD", "BB")

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
			bestNextValve := FindBestNextMove(valves, currentValve)
			if bestNextValve != nil {
				fmt.Println("Move to valve", bestNextValve.name)
				currentValve = bestNextValve
			}
		}

		fmt.Println("Current valve:", currentValve.name, ", release", currentPressure, "pressure")

		fmt.Println()
	}

	return totalPressure
}
