package day7

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strings"
)

type Node struct {
	name     string
	isLeaf   bool
	value    int64
	children []*Node
	parent   *Node
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	rootNode := &Node{value: 0}
	currentNode := rootNode
	for _, line := range lines {
		if line[0] == '$' {

			if line[2:4] == "ls" {
				fmt.Println("LS")
			} else {
				fmt.Println("CD")

				var directory string
				reader := strings.NewReader(line)
				fmt.Fscanf(reader, "$ cd %s", &directory)

				fmt.Println("DIR=", directory)
				if directory == ".." && currentNode.parent != nil {
					currentNode = currentNode.parent
				}
				if linq.From(currentNode.children).AnyWithT(func(n *Node) bool { return n.name == directory }) {
					fmt.Println("dir exists, moving to dir")
				} else {
					// new node
					newNode := &Node{name: directory}
				}
			}
		}
	}

	return len(lines)
}
