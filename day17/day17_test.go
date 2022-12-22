package day17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", false, false, 2022)

	fmt.Println(result)
	assert.Equal(t, 3068, result)
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", false, false, 2022)

	fmt.Println(result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", true, false, 1000000000000)

	fmt.Println(result)
	assert.Equal(t, 1514285714288, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", true, false, 1000000000000)

	fmt.Println(result)
}

func Test_Pattern_Compare(t *testing.T) {

	pattern1 := [][]bool{
		[]bool{true, false},
		[]bool{true, true},
	}
	pattern2 := [][]bool{
		[]bool{true, false},
		[]bool{true, true},
	}

	assert.True(t, IsSame(pattern1, pattern2))
}
