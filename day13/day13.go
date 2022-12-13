package day13

import (
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"sort"
)

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	indice := 1
	sum := 0
	packets := make([]*Node, 0)
	for i, _ := range lines {
		if i%3 == 0 {
			n := ReadNode(lines[i])
			n2 := ReadNode(lines[i+1])

			result := n.CompareReal(n2)
			//fmt.Println(n, " VS ", n2, "=", result)
			if result == -1 {
				//fmt.Println(indice)
				sum += indice
			}

			packets = append(packets, n, n2)

			indice++
		}
	}

	d1 := ReadNode("[[2]]")
	d2 := ReadNode("[[6]]")
	packets = append(packets, d1, d2)

	//fmt.Println(len(packets))

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].CompareReal(packets[j]) < 0
	})

	if !complex {
		return sum
	} else {
		//fmt.Println("Packets:")
		//for _, p := range packets {
		//	//fmt.Print(len(p.String()), ": ")
		//	fmt.Println(p)
		//}

		return (IndexOf(packets, d1) + 1) * (IndexOf(packets, d2) + 1)
	}
}
