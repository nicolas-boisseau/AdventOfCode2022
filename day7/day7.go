package day7

import (
	"bytes"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Node struct {
	name     string
	isDir    bool
	size     int
	children []*Node
	parent   *Node
}

func newNode(name string, isDir bool, parent *Node) *Node {
	n := Node{name: name}
	n.size = 0
	n.isDir = isDir
	n.children = make([]*Node, 0)
	n.parent = parent
	return &n
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	rootNode := newNode("/", true, nil)
	currentNode := rootNode
	isListing := false
	for _, line := range lines {
		if line[0] == '$' {
			isListing = false

			if line[2:4] == "ls" {
				isListing = true
			} else {

				var dirName string
				reader := strings.NewReader(line)
				fmt.Fscanf(reader, "$ cd %s", &dirName)

				//fmt.Println("DIR=", dirName)
				if dirName == ".." && currentNode.parent != nil {
					currentNode = currentNode.parent
				} else if dirName == "/" {
					currentNode = rootNode
				} else {
					if linq.From(currentNode.children).AnyWithT(func(n *Node) bool { return n.name == dirName }) {
						//fmt.Println("dir exists, moving to dir")

						dirNode := linq.From(currentNode.children).FirstWithT(func(n *Node) bool { return n.name == dirName }).(*Node)

						currentNode = dirNode

					} else {
						// new node
						newNode := newNode(dirName, true, currentNode)
						currentNode.children = append(currentNode.children, newNode)

						currentNode = newNode
					}
				}
			}
		} else if isListing {

			if line[0:3] == "dir" {
				var dirName string
				reader := strings.NewReader(line)
				fmt.Fscanf(reader, "dir %s", &dirName)

				if linq.From(currentNode.children).AnyWithT(func(n *Node) bool { return n.name == dirName }) {
					newNode := newNode(dirName, true, currentNode)
					currentNode.children = append(currentNode.children, newNode)
				}
			} else {
				// This is a file !
				var fileName string
				var fileSize int
				reader := strings.NewReader(line)
				fmt.Fscanf(reader, "%d %s", &fileSize, &fileName)

				fileNode := newNode(fileName, false, currentNode)
				fileNode.size = fileSize
				currentNode.children = append(currentNode.children, fileNode)
			}
		}
	}

	fmt.Println("=================================")
	//fmt.Println(rootNode.String(0))
	//fmt.Println(rootNode.DirSize())

	dirSizes := make(map[string]int, 0)
	dirSizes = rootNode.DirSizes(dirSizes)
	fmt.Println(dirSizes)

	var result int
	if !complex {
		result = int(linq.
			From(dirSizes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(int) }).
			WhereT(func(value int) bool { return value <= 100000 }).
			SumInts())
	} else {

		availableSpace := 70000000 - dirSizes["/"]
		spaceToFreeUp := 30000000 - availableSpace

		result = linq.
			From(dirSizes).
			SelectT(func(kv linq.KeyValue) int { return kv.Value.(int) }).
			WhereT(func(size int) bool { return size >= spaceToFreeUp }).
			OrderByT(func(size int) int { return size }).
			First().(int)
	}

	return result
}

func (n *Node) String(level int) string {
	buff := bytes.NewBufferString("")
	indent := ""
	for i := 0; i < level; i++ {
		indent += "-"
	}
	fmt.Fprintf(buff, "%s%s", indent, n.name)
	if n.isDir {
		fmt.Fprintf(buff, " (%d)", n.DirSize())
	}
	fmt.Fprintln(buff)
	if len(n.children) > 0 {
		for _, child := range n.children {
			fmt.Fprint(buff, child.String(level+1))
		}
		fmt.Fprintln(buff)
	}
	return buff.String()
}

func (n *Node) DirSize() int {
	size := 0
	if len(n.children) > 0 {
		for _, child := range n.children {
			if !child.isDir {
				size += child.size
			} else {
				size += child.DirSize()
			}
		}
	}

	return size
}

func (n *Node) DirSizes(dirSizes map[string]int) map[string]int {
	dirSize := n.DirSize()
	dirPath := n.ComputePath()
	if !linq.From(dirSizes).AnyWithT(func(kv linq.KeyValue) bool { return kv.Key.(string) == dirPath }) {
		dirSizes[dirPath] = dirSize
	}

	for _, child := range n.children {
		if child.isDir {
			dirSizes = child.DirSizes(dirSizes)
		}
	}

	return dirSizes
}

func (n *Node) ComputePath() string {
	path := n.name

	currentNode := n.parent
	for currentNode != nil {
		path = currentNode.name + "__" + path
		currentNode = currentNode.parent
	}

	return path
}
