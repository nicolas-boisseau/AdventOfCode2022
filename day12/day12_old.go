package day12

//
//import (
//	"bytes"
//	"fmt"
//	"github.com/ahmetalpbalkan/go-linq"
//	"github.com/golang-collections/collections/set"
//	"github.com/nicolas-boisseau/AdventOfCode2022/common"
//	"math"
//	"regexp"
//	"sort"
//)
//
//type Grid struct {
//	h       int
//	w       int
//	content [][]int
//	visited [][]bool
//}
//
//type Point struct {
//	x int
//	y int
//}
//
//func newGrid(h int, w int) *Grid {
//	g := Grid{h: h, w: w}
//	g.content = make([][]int, h)
//	g.visited = make([][]bool, h)
//	for i := range g.content {
//		g.content[i] = make([]int, w)
//		g.visited[i] = make([]bool, w)
//	}
//	return &g
//}
//
//type TreeNode struct {
//	name     string
//	children []*TreeNode
//	parent   *TreeNode
//}
//
//func newTreeNode(name string, isDir bool, parent *TreeNode) *TreeNode {
//	n := TreeNode{name: name}
//	n.children = make([]*TreeNode, 0)
//	n.parent = parent
//	return &n
//}
//
//func (g *Grid) String() string {
//	output := bytes.NewBufferString("")
//	for i := range g.content {
//		for j := range g.content[i] {
//			if g.visited[i][j] {
//				fmt.Fprint(output, "#")
//			} else {
//				fmt.Fprint(output, string(rune(g.content[i][j]+96)))
//			}
//		}
//
//		fmt.Fprintln(output)
//
//	}
//	return output.String()
//}
//
//func RuneToInt(r rune) int {
//	rexp, _ := regexp.Compile(`[a-z]`)
//	if rexp.MatchString(string(r)) {
//		return int(r) - 97 + 1
//	} else {
//		return int(r) - (65 - 27)
//	}
//}
//
//func Process(fileName string, complex bool) int {
//	lines := common.ReadLinesFromFile(fileName)
//
//	g := newGrid(len(lines), len(lines[0]))
//	var start, end *Point
//	for i, line := range lines {
//		//fmt.Println(line)
//
//		for j, c := range line {
//			n := 0
//			if c == 'S' {
//				start = &Point{x: i, y: j}
//			} else if c == 'E' {
//				end = &Point{x: i, y: j}
//			} else {
//				n = RuneToInt(c)
//			}
//			g.content[i][j] = n
//		}
//	}
//
//	fmt.Println(start)
//	fmt.Println(end)
//
//	fmt.Println(g)
//
//	return FindCheapestPath(g, start, end)
//}
//
//func FindCheapestPath(g *Grid, start *Point, end *Point) int {
//	graph := g.BuildGraph()
//	fmt.Println("Graph generated OK")
//	//fmt.Println(graph)
//
//	shortestPathCost := graph.FindCheapestPath(g, start.x, start.y, end.x, end.y)
//
//	return shortestPathCost
//}
//
//type Graph2 struct {
//	nodes [][]*Node
//}
//
//type Node struct {
//	x, y             int
//	cost             int
//	score, heuristic float64
//	steps            int
//	neighbors        []*Node
//}
//
//func (n *Node) String() string {
//	buff := bytes.NewBufferString("")
//	fmt.Fprintln(buff, "Pos:", n.x, ",", n.y, ", Cost:", n.cost)
//	fmt.Fprint(buff, "Voisins: ")
//	for _, v := range n.neighbors {
//		fmt.Fprint(buff, "{", v.x, v.y, "}")
//	}
//	fmt.Fprintln(buff)
//	return buff.String()
//}
//
//func (g *Graph2) String() string {
//	buff := bytes.NewBufferString("")
//	for i := 0; i < len(g.nodes); i++ {
//		for j := 0; j < len(g.nodes[i]); j++ {
//			fmt.Fprint(buff, g.nodes[i][j].cost, " ")
//		}
//		fmt.Fprintln(buff)
//	}
//
//	for i := 0; i < len(g.nodes); i++ {
//		for j := 0; j < len(g.nodes[i]); j++ {
//			for _, v := range g.nodes[i][j].neighbors {
//				fmt.Fprintln(buff, g.nodes[i][j].cost, "->", v.cost)
//			}
//		}
//	}
//
//	return buff.String()
//}
//
//func (g *Grid) BuildGraph() *Graph2 {
//	// Init nodes
//	output := &Graph2{
//		nodes: make([][]*Node, g.h),
//	}
//	for i, _ := range g.content {
//		output.nodes[i] = make([]*Node, g.w)
//		for j, _ := range g.content[i] {
//			output.nodes[i][j] = &Node{
//				x:         i,
//				y:         j,
//				cost:      g.content[i][j],
//				neighbors: make([]*Node, 0),
//			}
//		}
//	}
//
//	isInRange := func(x, y int) bool {
//		return x >= 0 && y >= 0 && x < g.h && y < g.w
//	}
//
//	addEdgeIfInRange := func(fromX, fromY int, toX, toY int, graph *Graph2) {
//		if isInRange(toX, toY) && graph.nodes[toX][toY].cost <= graph.nodes[fromX][fromY].cost+1 {
//			if !linq.From(graph.nodes[fromX][fromY].neighbors).AnyWithT(func(n *Node) bool { return n == graph.nodes[toX][toY] }) {
//				graph.nodes[fromX][fromY].neighbors = append(graph.nodes[fromX][fromY].neighbors, graph.nodes[toX][toY])
//			}
//			if !linq.From(graph.nodes[toX][toY].neighbors).AnyWithT(func(n *Node) bool { return n == graph.nodes[fromX][fromY] }) {
//				graph.nodes[toX][toY].neighbors = append(graph.nodes[toX][toY].neighbors, graph.nodes[fromX][fromY])
//			}
//		}
//	}
//
//	// Compute all neighbors
//	for i, _ := range g.content {
//		for j, _ := range g.content[i] {
//
//			addEdgeIfInRange(i, j, i+1, j, output)
//			addEdgeIfInRange(i, j, i, j+1, output)
//			addEdgeIfInRange(i, j, i-1, j, output)
//			addEdgeIfInRange(i, j, i, j-1, output)
//
//		}
//	}
//
//	return output
//}
//
////
////func CompareNodes(n1, n2 Node) int {
////	if n1.heuristic < n2.heuristic {
////		return 1
////	} else if n1.heuristic == n2.heuristic {
////		return 0
////	} else {
////		return -1
////	}
////}
//
//type byHeuristic []*Node
//
//func (h byHeuristic) Len() int {
//	return len(h)
//}
//func (h byHeuristic) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//func (h byHeuristic) Less(i, j int) bool {
//	return h[i].heuristic < h[j].heuristic
//}
//
//func (g *Graph2) GetAllNodes() []*Node {
//	nodes := make([]*Node, 0)
//	for i := 0; i < len(g.nodes); i++ {
//		for j := 0; j < len(g.nodes[i]); j++ {
//			nodes = append(nodes, g.nodes[i][j])
//		}
//	}
//	return nodes
//}
//
//func distance(x1, y1, x2, y2 int) float64 {
//	//return math.Sqrt(math.Pow(float64(x2) - float64(x1), 2.0) + math.Pow(float64(y2) - float64(y1), 2.0))
//	return (math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1)))
//}
//
//func (g *Graph2) FindCheapestPath(grid *Grid, fromX, fromY, toX, toY int) int {
//
//	currentNode := g.nodes[fromX][fromY]
//
//	steps := 0
//	for !(currentNode.x == toX && currentNode.y == toY) {
//
//		grid.visited[currentNode.x][currentNode.y] = true
//		fmt.Println(grid)
//
//		//var cheapestMoves
//		//linq.From(currentNode.neighbors).
//		//	WhereT(func(n *Node) bool { return !grid.visited[n.x][n.y] }).
//		//	OrderByDescendingT(func(n *Node) float64 { return distance(n.x, n.y, toX, toY) }).
//		//	ThenByT(func(n *Node) float64 { return float64(n.cost) })
//		//	SelectT()
//
//		stepsByMove
//		for _, n := range currentNode.neighbors {
//
//		}
//
//		steps++
//	}
//
//	return steps
//}
//
//// https://yourbasic.org/golang/implement-fifo-queue/
//// https://pkg.go.dev/container/list
//// https://pkg.go.dev/container/heap@go1.17.5
//// https://gobyexample.com/sorting-by-functions
//func (g *Graph2) FindShortestPath(fromX, fromY, toX, toY int) int {
//
//	closedList := set.New()
//	openList := make([]*Node, 0)
//
//	//for i := 0; i < len(g.nodes); i++ {
//	//	for j := 0; j < len(g.nodes[i]); j++ {
//	//		g.nodes[i][j].heuristic = distance(i, j, toX, toY)
//	//		g.nodes[i][j].score = float64(g.nodes[i][j].cost)
//	//	}
//	//}
//	//g.nodes[0][0].score = 0
//
//	openList = append(openList, g.nodes[fromX][fromY])
//	sort.Sort(byHeuristic(openList))
//
//	for len(openList) > 0 {
//		// defiler
//		u := openList[0]
//		openList = openList[1:]
//
//		if u.x == toX && u.y == toY {
//			//fmt.Println("fin")
//			fmt.Println(u.score)
//			fmt.Println(u.steps)
//			fmt.Println(u.heuristic)
//			//fmt.Println(u.cost)
//			//fmt.Println(u)
//			//fmt.Println(closedList.Len())
//
//			g_final := newGrid(len(g.nodes), len(g.nodes[0]))
//			for i := 0; i < len(g.nodes); i++ {
//				for j := 0; j < len(g.nodes[i]); j++ {
//					g_final.content[i][j] = g.nodes[i][j].cost
//				}
//			}
//			closedList.Do(func(nn interface{}) {
//				node := nn.(*Node)
//				g_final.visited[node.x][node.y] = true
//			})
//			fmt.Println(g_final)
//
//			return int(u.heuristic)
//		}
//
//		//var filteredNeighbors []*Node
//		//linq.From(u.neighbors).WhereT(func (n *Node) bool { return n.cost <= u.cost+1 }).ToSlice(&filteredNeighbors)
//		for _, v := range u.neighbors {
//			if !closedList.Has(v) && !linq.From(openList).AnyWithT(func(nn *Node) bool { return nn == v }) {
//				//if !closedList.Has(v) {
//				cost := float64(v.cost)
//				if u.score <= v.score {
//					cost = float64(0)
//				}
//				v.score = u.score + cost
//				v.steps = u.steps + 1
//				v.heuristic = float64(v.score) //float64(v.score) + distance(v.x, v.y, toX, toY)
//				openList = append(openList, v)
//				sort.Sort(byHeuristic(openList))
//			}
//		}
//
//		closedList.Insert(u)
//	}
//
//	return -1
//}
