package day3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_runes(t *testing.T) {
	fmt.Println('a' - 97 + 1)
	fmt.Println('A' - (65 - 27))

	rexp, _ := regexp.Compile(`[a-z]`)
	fmt.Println(rexp.MatchString("a"))
	fmt.Println(rexp.MatchString("A"))
}

func Test_RuneToInt(t *testing.T) {
	assert.Equal(t, 1, RunToInt('a'))
	assert.Equal(t, 4, RunToInt('d'))
	assert.Equal(t, 27, RunToInt('A'))
	assert.Equal(t, 52, RunToInt('Z'))
}

func Test_Process_Sample(t *testing.T) {
	result := Process("sample.txt", false)

	fmt.Println(result)
	assert.Equal(t, 157, result)
}

func Test_Process_Input(t *testing.T) {
	result := Process("input.txt", false)

	fmt.Println(result)
}

func Test_Process_Sample_Complex(t *testing.T) {
	result := Process("sample.txt", true)

	fmt.Println(result)
	assert.Equal(t, 70, result)
}

func Test_Process_Input_Complex(t *testing.T) {
	result := Process("input.txt", true)

	fmt.Println(result)
}
