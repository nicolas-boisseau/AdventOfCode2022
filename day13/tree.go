package day13

import (
	"bytes"
	"fmt"
	//"github.com/golang-collections/collections/set"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	isLeaf bool
	value  int64
	childs []*Node
	parent *Node
}

func (n *Node) String() string {
	buff := bytes.NewBufferString("")
	if n.isLeaf {
		fmt.Fprint(buff, n.value)
	} else {
		fmt.Fprint(buff, "[")
		for i, child := range n.childs {
			fmt.Fprint(buff, child)
			if i < len(n.childs)-1 {
				fmt.Fprint(buff, ",")
			}
		}
		fmt.Fprint(buff, "]")
	}

	return buff.String()
}

func ReadNode(str string) *Node {
	root := Node{}

	offset := 0
	re, _ := regexp.Compile("[0-9]+")

	currentNode := &root
	for offset < len(str) || currentNode == nil {

		if string(str[offset]) == "[" {
			child := Node{
				parent: currentNode,
				isLeaf: false,
				childs: make([]*Node, 0),
			}
			currentNode.childs = append(currentNode.childs, &child)
			currentNode = &child
		} else if string(str[offset]) == "," {
			currentNode = currentNode.parent
			child := Node{
				parent: currentNode,
				isLeaf: false,
				childs: make([]*Node, 0),
			}
			currentNode.childs = append(currentNode.childs, &child)
			currentNode = &child
		} else if string(str[offset]) == "]" {
			//if currentNode.parent == &root && linq.From(currentNode.parent.childs).AllT(func(n *Node) bool { return !n.isValid }) {
			//	currentNode.parent.childs = []*Node{}
			//}
			currentNode = currentNode.parent
		} else {
			nextComma := strings.Index(str[offset:], ",")
			rightLimit := nextComma // default
			nextClosingSquareBrace := strings.Index(str[offset:], "]")
			if nextComma == -1 || (nextClosingSquareBrace != -1 && nextClosingSquareBrace < nextComma) {
				rightLimit = nextClosingSquareBrace
			}
			if rightLimit == -1 {
				rightLimit = len(str) - 1
			}
			if re.MatchString(str[offset:]) {
				val, err := strconv.ParseInt(str[offset:offset+rightLimit], 10, 64)
				if err != nil {
					log.Fatalln(err)
				}
				currentNode.isLeaf = true
				currentNode.value = val
				offset += rightLimit
				continue
			}
		}

		offset++
	}

	return &root
}

func IndexOf(list []*Node, toSearch *Node) int {
	for i, n := range list {
		if n == toSearch {
			return i
		}
	}
	return -1
}

func (n *Node) NextInListNotVisited(visited map[*Node]bool) (*Node, bool) {

	for _, n := range n.childs {
		if !visited[n] {
			return n, false
		}
	}

	//fmt.Println(n, "is running out of items !")

	return n.parent, true
}

func (n *Node) Next(visited map[*Node]bool) (*Node, bool) {

	visited[n] = true

	return n.NextInListNotVisited(visited)
}

func (n *Node) CompareReal(n2 *Node) int {
	visited := make(map[*Node]bool)
	return n.Compare2(n2, visited)
}

func (n *Node) Compare2(n2 *Node, visited map[*Node]bool) int {

	//fmt.Println("Comparing: ", n, "with", n2)

	left, leftOutOfItems := n.Next(visited)
	right, rightOutOfItems := n2.Next(visited)

	if leftOutOfItems && !rightOutOfItems {
		return -1
	} else if !leftOutOfItems && rightOutOfItems {
		return 1
	}

	//fmt.Println("Next() compare: ", left, "with", right)

	if left == nil {
		return -1
	} else if right == nil {
		return 1
	}

	if !left.isLeaf {
		if right.isLeaf {
			right.isLeaf = false
			right.childs = []*Node{&Node{isLeaf: true, value: right.value, childs: []*Node{}, parent: right}}
			return left.Compare2(right, visited)
		}

		if left == nil {
			return -1
		} else if right == nil {
			return 1
		}

		return left.Compare2(right, visited)

	} else if !right.isLeaf {
		left.isLeaf = false
		left.childs = []*Node{&Node{isLeaf: true, value: left.value, childs: []*Node{}, parent: left}}
		return left.Compare2(right, visited)
	}

	if left.value < right.value {
		return -1
	} else if left.value > right.value {
		return 1
	}

	// return -1 for <, 0 for == or 1 for >
	return left.Compare2(right, visited)
}
