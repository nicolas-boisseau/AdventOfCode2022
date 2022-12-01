package common

import (
	"bufio"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLinesFromFile(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	lines := make([]string, 0)
	for scanner.Scan() {
		//		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ConvertAsArrayOfInt(arrayOfString []string) []int64 {
	arrayOfInts := make([]int64, 0)
	for _, str := range arrayOfString {
		n, _ := strconv.ParseInt(str, 10, 0)
		arrayOfInts = append(arrayOfInts, n)
	}

	return arrayOfInts
}
