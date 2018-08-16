package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixScore(t *testing.T) {
	as := assert.New(t)
	output := matrixScore([][]int{[]int{0, 0, 1, 1}, []int{1, 0, 1, 0}, []int{1, 1, 0, 0}})
	as.Equal(39, output)
}
