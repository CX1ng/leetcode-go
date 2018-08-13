package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	as := assert.New(t)
	output := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	as.Equal(6, output.Val)
	as.Equal(3, output.Left.Val)
	as.Equal(2, output.Left.Right.Val)
	as.Equal(1, output.Left.Right.Right.Val)
	as.Equal(5, output.Right.Val)
	as.Equal(0, output.Right.Left.Val)
}
