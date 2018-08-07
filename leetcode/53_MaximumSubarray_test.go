package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarray(t *testing.T){
	as := assert.New(t)
	output := maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	as.Equal(6, output)
}
