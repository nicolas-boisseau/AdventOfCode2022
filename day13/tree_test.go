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
	s1 := "[[],[],[5,4,0,10,7],[2,[[],2,0,[7,4,7,7,10],[]],[]]]"
	s2 := "[[],[],[[5,[2,2,8,5,7],3,9,[4,6,0,2,0]],1,7,0],[],[[[7,6,5],9,[2,2,10,5,6]],4,[0,[],[9,4,1,8]],8,7]]"
	n1 := ReadNode(s1)
	n2 := ReadNode(s2)

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, -1, n1.CompareReal(n2))
}

func Test_Compare9(t *testing.T) {
	s1 := "[[],[10,[5],[],[6,[2,4,4,8,9],[0,0,8,9],[8],4]],[5,[],7,3,[[5],[8,9,1],7,[9]]],[6,5]]"
	s2 := "[[],[[10],9,5,9,1],[7,4,[]],[6,[10,6,[2],[5,10,4,9]]]]"
	n1 := ReadNode(s1)
	n2 := ReadNode(s2)

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, -1, n1.CompareReal(n2))
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

func Test_Compare_SuspiciousNodes(t *testing.T) {

	n1 := ReadNode("[[],[],[5,4,0,10,7],[2,[[],2,0,[7,4,7,7,10],[]],[]]]")
	n2 := ReadNode("[[],[],[7,[0,[0,6,6,9,4],[],[8,3]],3,7,10]]")

	fmt.Println(n1.String())
	fmt.Println(n2.String())

	assert.Equal(t, -1, n1.CompareReal(n2))
}
