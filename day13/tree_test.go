package day13

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Node_String(t *testing.T) {

	childs := []*Node{
		&Node{isLeaf: true, value: 4, childs: []*Node{}},
		&Node{isLeaf: true, value: 6, childs: []*Node{}},
	}

	n := Node{
		isLeaf: false,
		childs: childs,
	}

	assert.Equal(t, "[4,6]", n.String())
}

func Test_Node_String_Complex(t *testing.T) {

	n := Node{
		isLeaf: false,
		childs: []*Node{
			&Node{
				isLeaf: true,
				value:  4,
			},
			&Node{
				isLeaf: false,
				childs: []*Node{
					&Node{
						isLeaf: false,
						childs: []*Node{
							&Node{
								isLeaf: true,
								value:  45,
							},
							&Node{
								isLeaf: true,
								value:  17,
							},
						},
					},
					&Node{
						isLeaf: true,
						value:  22,
					},
				},
			},
		},
	}

	assert.Equal(t, "[4,[[45,17],22]]", n.String())
}

func Test_ReadNode(t *testing.T) {

	sample := "[4,22]"

	n := ReadNode(sample)

	assert.Equal(t, sample, n.String())
}

func Test_ReadNode2(t *testing.T) {

	sample := "[4,[12,42]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

func Test_ReadNode3(t *testing.T) {

	sample := "[[34,52],[12,42]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

//func Test_ReadNode5(t *testing.T) {
//
//	sample := "[]"
//
//	n := ReadNode(sample)
//
//	fmt.Println(n)
//
//	assert.Equal(t, sample, n.String())
//}

func Test_ReadNode4(t *testing.T) {

	sample := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"

	n := ReadNode(sample)

	fmt.Println(n)

	assert.Equal(t, sample, n.String())
}

func Test_Add_Nodes(t *testing.T) {
	n1Str := "[1,2]"
	n2Str := "[[3,4],5]"
	n1 := ReadNode(n1Str)
	n2 := ReadNode(n2Str)

	n3 := Add(n1, n2)

	fmt.Println(n3)
	assert.Equal(t, "[[1,2],[[3,4],5]]", n3.String())
}

func Test_Next(t *testing.T) {
	n1 := ReadNode("[4,3,2,1,0]")

	visited := make(map[*Node]bool)

	curNode, _ := n1.Next(visited)
	assert.Equal(t, "4", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "3", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "2", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "1", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "0", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Nil(t, curNode)
}

func Test_Next2(t *testing.T) {
	curNode := ReadNode("[[1],[2,3,4]]")

	visited := make(map[*Node]bool)

	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "[1]", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "1", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "[2,3,4]", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "2", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "3", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Equal(t, "4", curNode.String())
	curNode, _ = curNode.Next(visited)
	assert.Nil(t, curNode)
	//assert.Equal(t, int64(1), curNode.Next(visited).value)
	//assert.Equal(t, int64(0), curNode.Next(visited).Next(visited).value)
	//assert.Equal(t, nil, curNode.Next(visited).Next(visited).Next(visited))
}

func Test_Compare(t *testing.T) {
	n1 := ReadNode("[1,1,3,1,1]")
	n2 := ReadNode("[1,1,5,1,1]")

	assert.Equal(t, -1, n1.CompareReal(n2))
}

func Test_Compare2(t *testing.T) {
	n1 := ReadNode("[[1],[2,3,4]]")
	n2 := ReadNode("[[1],4]")

	assert.Equal(t, -1, n1.CompareReal(n2))
}

func Test_Compare3(t *testing.T) {
	n1 := ReadNode("[9]")
	n2 := ReadNode("[[8,7,6]]")

	assert.Equal(t, 1, n1.CompareReal(n2))
}

func Test_Compare4(t *testing.T) {
	n1 := ReadNode("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	n2 := ReadNode("[1,[2,[3,[4,[5,6,0]]]],8,9]")

	assert.Equal(t, 1, n1.CompareReal(n2))
}

func Test_Compare5(t *testing.T) {
	n1 := ReadNode("[]")
	n2 := ReadNode("[[1],[2,3,4]]")

	assert.Equal(t, -1, n1.CompareReal(n2))
}

func Test_Compare6(t *testing.T) {
	n1 := ReadNode("[7,7,7,7]")
	n2 := ReadNode("[7,7,7]")

	assert.Equal(t, 1, n1.CompareReal(n2))
}

func Test_ReadNodeWithEmpty(t *testing.T) {
	n1 := ReadNode("[[[]]]")
	n2 := ReadNode("[[]]")

	fmt.Println(n1.String())
	fmt.Println(n2.String())
}

func Test_Compare7(t *testing.T) {
	n1 := ReadNode("[[[4,6,9,3,1],4,[3],[2,4,[6,10]],6],[]]")
	n2 := ReadNode("[[4,8],[[],6],[3,4]]")

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, 1, n1.CompareReal(n2))
}

func Test_Compare8(t *testing.T) {
	n1 := ReadNode("[[[2,7,[10,6,0],[10]],[[1,6,9],9],[[2],[],1,[8]]],[[]]]")
	n2 := ReadNode("[[],[5,9,3,[9,[10,8,10,1],6],[3]],[[10,3,[6,5,1,5]]],[5,[]]]")

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, 1, n1.CompareReal(n2))
}

func Test_CompareEmptyNodes(t *testing.T) {
	n1 := ReadNode("[[]]")
	n2 := ReadNode("[]")

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, 1, n1.CompareReal(n2))

	fmt.Println(n1.String())
	fmt.Println(n2.String())
}
