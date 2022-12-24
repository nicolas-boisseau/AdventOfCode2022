package day20

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
)

type Num struct {
	n        int
	pos      int
	next     *Num
	previous *Num
}

func FixPos(pos int, maxPos int) int {
	if pos > maxPos {
		return (pos + 1) % maxPos
	} else if pos < 0 {
		return FixPos(maxPos+pos, maxPos)
	}

	return pos
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	nums := make([]*Num, 0)

	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}

		nums = append(nums, &Num{
			n:   n,
			pos: i,
		})
	}

	// affect next & previous pointers
	for i := range nums {
		if i+1 < len(nums) {
			nums[i].next = nums[i+1]
		} else {
			nums[i].next = nums[0]
		}

		if i-1 >= 0 {
			nums[i].previous = nums[i-1]
		} else {
			nums[i].previous = nums[len(nums)-1]
		}
	}

	if fileName == "sample.txt" {
		fmt.Println("Before starting mixing!")
		PrintNums(nums)
	}

	// Start !
	for _, n := range nums {
		diff := n.n
		if n.n < 0 {
			diff -= 1
		}
		newPos := FixPos(n.pos+diff, len(nums))
		if newPos > n.pos {
			// tout ce qui était avant a - 1
			for i := n.pos; i <= newPos; i++ {
				if i == n.pos {
					continue
				}
				u := NumAtPos(nums, i)
				u.pos--
			}
		} else if newPos < n.pos {
			// tout ce qui était après a + 1
			for i := n.pos; i >= newPos; i-- {
				if i == n.pos {
					continue
				}
				u := NumAtPos(nums, i)
				u.pos++
			}
		}

		n.pos = newPos

		if fileName == "sample.txt" {
			fmt.Println("After mixing ", n.n)
			PrintNums(nums)
		}
	}

	zeroPos := NumPos(nums, 0)

	return NumAtPos(nums, FixPos(zeroPos+1000-1, len(nums))).n +
		NumAtPos(nums, FixPos(zeroPos+2000-1, len(nums))).n +
		NumAtPos(nums, FixPos(zeroPos+3000-1, len(nums))).n
}

func NumAtPos(nums []*Num, pos int) *Num {
	return linq.From(nums).WhereT(func(u *Num) bool { return u.pos == pos }).First().(*Num)
}

func NumPos(nums []*Num, n int) int {
	return linq.From(nums).WhereT(func(u *Num) bool { return u.n == n }).First().(*Num).pos
}

func PrintNums(nums []*Num) {
	for i := 0; i < len(nums); i++ {
		fmt.Print(NumAtPos(nums, i).n, ",")
	}
	fmt.Println()
}
