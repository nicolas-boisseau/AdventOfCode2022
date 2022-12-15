package day15

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", 10, -1, false)

	fmt.Println(result)
	assert.Equal(t, 26, result)
}

func Test_1(t *testing.T) {
	str := "10,5"
	fmt.Println(str[0:3])
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", 2000000, -1, false)

	fmt.Println(result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", 0, 20, true)

	fmt.Println(result)
	assert.Equal(t, 56000011, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", 0, 4000000, true)

	fmt.Println(result)
}
