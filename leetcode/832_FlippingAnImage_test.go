package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlipAnInvertImage(t *testing.T) {
	as := assert.New(t)
	output := flipAndInvertImage([][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}})
	as.Equal(3, len(output))
	as.Equal([]int{1, 0, 0}, output[0])
	as.Equal([]int{0, 1, 0}, output[1])
	as.Equal([]int{1, 1, 1}, output[2])
	output = flipAndInvertImage([][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}})
	as.Equal(4, len(output))
	as.Equal([]int{1, 1, 0, 0}, output[0])
	as.Equal([]int{0, 1, 1, 0}, output[1])
	as.Equal([]int{0, 0, 0, 1}, output[2])
	as.Equal([]int{1, 0, 1, 0}, output[3])
}
