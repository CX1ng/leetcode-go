package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	as := assert.New(t)
	output := minCostClimbingStairs([]int{10, 15, 20})
	as.Equal(15, output)
	output = minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	as.Equal(6, output)
}
