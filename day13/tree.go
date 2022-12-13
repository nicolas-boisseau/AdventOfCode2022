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
	//left   *Node
	//right  *Node
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

func Add(n1, n2 *Node) *Node {
	output := &Node{
		parent: nil,
		//left:   n1,
		//right:  n2,
		childs: []*Node{n1, n2},
	}
	n1.parent = output
	n2.parent = output

	//reduced := output.Explode(1)
	//splitted := output.Split()
	//for reduced || splitted {
	//	reduced = output.Explode(1)
	//	splitted = output.Split()
	//}

	return output
}

//

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

//func (n *Node) FindLeftLeafNode() *Node {
//	if n.parent == nil {
//		return nil
//	}
//
//	currentNode := n.parent
//	//fmt.Println("Starting with : ", currentNode.String())
//	// on remonte jusqu'à trouver un noeuds avec une branche gauche
//	previousNode := n
//	for currentNode != nil {
//		if currentNode.left != previousNode {
//			break
//		}
//		previousNode = currentNode
//		currentNode = currentNode.parent
//		if currentNode != nil {
//			//fmt.Println("Moving to parent : ", currentNode.String())
//		}
//	}
//
//	if currentNode == nil {
//		//fmt.Println("No left...")
//		return nil
//	}
//
//	// on prend la direction gauche
//	currentNode = currentNode.left
//	//fmt.Println("Take to left child : ", currentNode.String())
//
//	// si trouvé, alors on descends jusqu'à la feuille la plus à droite
//	if currentNode != nil {
//		for !currentNode.isLeaf {
//			currentNode = currentNode.right
//			//fmt.Println("Descending to child right : ", currentNode.String())
//		}
//		return currentNode
//	}
//
//	return nil
//}
//
//func (n *Node) FindRightLeafNode() *Node {
//	if n.parent == nil {
//		return nil
//	}
//
//	currentNode := n.parent
//	//fmt.Println("Starting with : ", currentNode.String())
//	// on remonte jusqu'à trouver un noeuds avec une branche droite
//	previousNode := n
//	for currentNode != nil {
//		if currentNode.right != previousNode {
//			break
//		}
//		previousNode = currentNode
//		currentNode = currentNode.parent
//		if currentNode != nil {
//			//fmt.Println("Moving to parent : ", currentNode.String())
//		}
//	}
//
//	if currentNode == nil {
//		///fmt.Println("No right...")
//		return nil
//	}
//
//	// on prend la direction droite
//	currentNode = currentNode.right
//
//	// si trouvé, alors on descends jusqu'à la feuille la plus à gauche
//	if currentNode != nil {
//		for !currentNode.isLeaf {
//			currentNode = currentNode.left
//		}
//		return currentNode
//	}
//
//	return nil
//}
//
//func (n *Node) Explode(level int) bool {
//
//	somethingHappen := false
//
//	if level > 4 && n.left != nil && n.right != nil && n.left.isLeaf && n.right.isLeaf {
//		//fmt.Println("Exploding:", n.left.value, n.right.value)
//		left := n.FindLeftLeafNode()
//		if left != nil {
//			left.value += n.left.value
//		}
//		right := n.FindRightLeafNode()
//		if right != nil {
//			right.value += n.right.value
//		}
//		n.isLeaf = true
//		n.value = 0
//		n.left = nil
//		n.right = nil
//		somethingHappen = true
//	}
//	if n.left != nil {
//		somethingHappenInChild := n.left.Explode(level + 1)
//		somethingHappen = somethingHappen || somethingHappenInChild
//	}
//	if n.right != nil {
//		somethingHappenInChild := n.right.Explode(level + 1)
//		somethingHappen = somethingHappen || somethingHappenInChild
//	}
//
//	return somethingHappen
//}
//
//func (n *Node) Split() bool {
//	nodesToSplit := n.SelectNodeToSplit()
//
//	//for _, nodeToSplit := range nodesToSplit {
//	if len(nodesToSplit) > 0 {
//		// Split only first node
//		nodeToSplit := nodesToSplit[0]
//		//fmt.Println("Spliting:", nodeToSplit.value)
//		nodeToSplit.isLeaf = false
//		nodeToSplit.left = &Node{
//			parent: nodeToSplit,
//			isLeaf: true,
//			value:  int64(math.Floor(float64(nodeToSplit.value) / 2.0)),
//		}
//		nodeToSplit.right = &Node{
//			parent: nodeToSplit,
//			isLeaf: true,
//			value:  int64(math.Ceil(float64(nodeToSplit.value) / 2.0)),
//		}
//	}
//
//	return len(nodesToSplit) > 0
//}
//
//func (n *Node) SelectNodeToSplit() []*Node {
//
//	selectedNodes := make([]*Node, 0)
//
//	if n.isLeaf && n.value >= 10 {
//		//fmt.Println("Spliting:", n.value)
//		//n.isLeaf = false
//		//n.left = &Node{
//		//	parent: n,
//		//	isLeaf: true,
//		//	value:  int64(math.Floor(float64(n.value) / 2.0)),
//		//}
//		//n.right = &Node{
//		//	parent: n,
//		//	isLeaf: true,
//		//	value:  int64(math.Ceil(float64(n.value) / 2.0)),
//		//}
//
//		selectedNodes = append(selectedNodes, n)
//	}
//	if n.left != nil {
//		childNodes := n.left.SelectNodeToSplit()
//		selectedNodes = append(selectedNodes, childNodes...)
//	}
//	if n.right != nil {
//		childNodes := n.right.SelectNodeToSplit()
//		selectedNodes = append(selectedNodes, childNodes...)
//
//	}
//
//	return selectedNodes
//}
//
//func (n *Node) Magnitude() int64 {
//
//	var magnitude int64 = 0
//
//	if n.isLeaf {
//		if n.parent.left == n {
//			magnitude += n.value
//		} else {
//			magnitude += n.value
//		}
//	} else {
//		magnitude += 3*n.left.Magnitude() + 2*n.right.Magnitude()
//	}
//
//	return magnitude
//}
