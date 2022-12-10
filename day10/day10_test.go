package day10

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", false)

	fmt.Println(result)
	assert.Equal(t, -1, result)
}

func Test_Process_Sample2(t *testing.T) {
	result := Process("sample2.txt", false)

	fmt.Println(result)
	assert.Equal(t, 13140, result)
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", false)

	fmt.Println(result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", true)

	fmt.Println(result)
	assert.Equal(t, 1, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", true)

	fmt.Println(result)
}
