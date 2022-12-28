package day20

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
)

type Num struct {
	n        int
	next     *Num
	previous *Num
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	nums := make([]*Num, 0)

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}

		nums = append(nums, &Num{
			n: n,
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
	var zeroNum *Num
	for _, num := range nums {

		if num.n > 0 {
			target := num.next
			for i := 1; i < num.n; i++ {
				target = target.next
			}
			// Replace current pos
			old_next := num.next
			old_prev := num.previous
			old_prev.next = old_next
			old_next.previous = old_prev
			// Insert at new pos
			old_target_next := target.next
			target.next = num
			old_target_next.previous = num
			num.previous = target
			num.next = old_target_next
		} else if num.n < 0 {
			target := num.previous
			for i := -1; i > num.n; i-- {
				target = target.previous
			}
			// Replace current pos
			old_next := num.next
			old_prev := num.previous
			old_prev.next = old_next
			old_next.previous = old_prev
			// Insert at new pos
			old_target_prev := target.previous
			target.previous = num
			old_target_prev.next = num
			num.previous = old_target_prev
			num.next = target
		} else {
			zeroNum = num
		}

		if fileName == "sample.txt" {
			fmt.Println("After mixing ", num.n)
			PrintNums(nums)
		}
	}

	var zeroPlus1000, zeroPlus2000, zeroPlus3000 int
	current := zeroNum
	for i := 1; i <= 3000; i++ {
		current = current.next
		if i == 1000 {
			zeroPlus1000 = current.n
		} else if i == 2000 {
			zeroPlus2000 = current.n
		} else if i == 3000 {
			zeroPlus3000 = current.n
		}
	}

	return zeroPlus1000 + zeroPlus2000 + zeroPlus3000
}

func PrintNums(nums []*Num) {
	first := nums[0]
	current := first
	for {
		fmt.Print(current.n, ",")
		current = current.next

		if current == first {
			break
		}
	}
	fmt.Println()
}
