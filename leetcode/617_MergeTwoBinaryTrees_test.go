package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTrees(t *testing.T) {
	as := assert.New(t)
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}
	output := mergeTrees(t1, t2)
	as.Equal(3, output.Val)
	as.Equal(4, output.Left.Val)
	as.Equal(5, output.Left.Left.Val)
	as.Equal(4, output.Left.Right.Val)
	as.Equal(5, output.Right.Val)
	as.Equal(7, output.Right.Right.Val)
}
