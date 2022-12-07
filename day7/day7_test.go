package day7

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", false)

	fmt.Println(result)
	assert.Equal(t, 95437, result)
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", false)

	fmt.Println(result)
	assert.Equal(t, 919137, result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", true)

	fmt.Println(result)
	assert.Equal(t, 24933642, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", true)

	fmt.Println(result)
}
