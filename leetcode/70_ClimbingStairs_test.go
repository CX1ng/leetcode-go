package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	as := assert.New(t)
	// Input: 2
	// Output: 2
	output := climbStairs(2)
	as.Equal(2, output)
	// Input: 3
	// Output: 3
	output = climbStairs(3)
	as.Equal(3, output)
}
