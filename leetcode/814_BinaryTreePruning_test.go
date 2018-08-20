package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPruneTree(t *testing.T) {
	as := assert.New(t)
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 1}}
	output := pruneTree(root)
	as.Equal(1, output.Val)
	as.Equal((*TreeNode)(nil), output.Left)
	as.Equal(0, output.Right.Val)
	as.Equal((*TreeNode)(nil), output.Right.Left)
	as.Equal(1, output.Right.Right.Val)
	as.Equal((*TreeNode)(nil), output.Right.Right.Left)
	as.Equal((*TreeNode)(nil), output.Right.Right.Right)
}
