package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", false)

	fmt.Println(result)
	assert.Equal(t, 24000, result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", true)

	fmt.Println(result)
	assert.Equal(t, 45000, result)
}

func Test_Process_Input1(t *testing.T) {
	result := Process("input1.txt", false)

	fmt.Println(result)
}

func Test_Process_Input1_Complex(t *testing.T) {
	result := Process("input1.txt", true)

	fmt.Println(result)
}
