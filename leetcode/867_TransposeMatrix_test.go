package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	as := assert.New(t)
	output := transpose([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	as.Len(output, 3)
	as.Equal([]int{1, 4, 7}, output[0])
	as.Equal([]int{2, 5, 8}, output[1])
	as.Equal([]int{3, 6, 9}, output[2])
	output = transpose([][]int{[]int{1, 2, 3}, []int{4, 5, 6}})
	as.Len(output, 3)
	as.Equal([]int{1, 4}, output[0])
	as.Equal([]int{2, 5}, output[1])
	as.Equal([]int{3, 6}, output[2])
}
