package day16

import (
	"bytes"
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"regexp"
)

type Graph struct {
	nodes []string
	edges []*Edge
	sum   int
}

type Edge struct {
	start string
	end   string
}

func newEdge(start, end string) *Edge {
	e := Edge{
		start: start,
		end:   end,
	}
	return &e
}

func newGraph() *Graph {
	g := Graph{
		nodes: make([]string, 0),
		edges: make([]*Edge, 0),
	}
	return &g
}

func (g *Graph) AddEdge(n1, n2 string) {
	if !From(g.nodes).Contains(n1) {
		g.nodes = append(g.nodes, n1)
	}
	if !From(g.nodes).Contains(n2) {
		g.nodes = append(g.nodes, n2)
	}

	g.edges = append(g.edges, newEdge(n1, n2))
	g.edges = append(g.edges, newEdge(n2, n1))
	//g.edges[n1] = n2
	//g.edges[n2] = n1
}

func (g *Graph) String() string {
	buffer := bytes.NewBufferString("")
	fmt.Fprintln(buffer, "NODES:", g.nodes)
	fmt.Fprintln(buffer, "EDGES:")
	//visited := make([]string, 0)
	for _, e := range g.edges {
		fmt.Fprintln(buffer, e.start, "->", e.end)
	}
	return buffer.String()
}

const startNode = "start"
const endNode = "end"

// Prints all paths from 's' to 'd'
func (g *Graph) printAllPaths(s, d string) {
	beingVisited := make(map[string]int, len(g.nodes))
	currentPath := make([]string, 0)

	currentPath = append(currentPath, s)
	g.dfs(s, d, beingVisited, currentPath)
}

func copyMap(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func (g *Graph) dfs(u, d string, beingVisited map[string]int, currentPath []string) {
	//fmt.Println("Visiting", u, "currentPath=", currentPath)
	isLowerCase, _ := regexp.MatchString("[a-z]", u)
	if isLowerCase {
		beingVisited[u] = beingVisited[u] + 1
	}
	//fmt.Println(beingVisited)

	if u == d {
		fmt.Println("Path found:", currentPath)
		g.sum++
		return
	}

	adjOfU := make([]string, 0)
	From(g.edges).
		WhereT(func(e *Edge) bool { return e.start == u }).
		SelectT(func(e *Edge) string { return e.end }).
		ToSlice(&adjOfU)
	//fmt.Println("Adj of", u, "=", adjOfU)

	for _, adj := range adjOfU {

		jokerAlreadyDone := From(beingVisited).AnyWithT(func(kv KeyValue) bool { return kv.Value.(int) > 1 })
		if beingVisited[adj] < 1 || (!jokerAlreadyDone && beingVisited[adj] < 2) && adj != "start" {
			currentPath = append(currentPath, adj)
			g.dfs(adj, d, copyMap(beingVisited), currentPath)
			//currentPath.remove(i);
			i := indexOf(currentPath, adj)
			//fmt.Println("Before removing ", adj, ":", currentPath)
			currentPath = append(currentPath[:i], currentPath[i+1:]...)
			//fmt.Println("After removed ", adj, ":", currentPath)

		}
	}

	if isLowerCase {
		beingVisited[u] = beingVisited[u] - 1
	}
}

func indexOf(array []string, str string) int {
	for i, v := range array {
		if v == str {
			return i
		}
	}
	return -1
}
