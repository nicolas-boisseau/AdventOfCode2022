package day9

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", 2, 10, 10)

	fmt.Println(result)
	assert.Equal(t, 13, result)
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", 2, 1000, 1000)

	fmt.Println(result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", 10, 10, 10)

	fmt.Println(result)
	assert.Equal(t, 1, result)
}

func Test_Process_Sample_Complex2(t *testing.T) {
	result := Process("sample2.txt", 10, 100, 100)

	fmt.Println(result)
	assert.Equal(t, 36, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", 10, 1000, 1000)

	fmt.Println(result)
}
