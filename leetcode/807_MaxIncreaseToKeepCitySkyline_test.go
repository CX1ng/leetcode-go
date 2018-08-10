package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxIncreaseToKeepCitySkyline(t *testing.T) {
	as := assert.New(t)
	output := maxIncreaseKeepingSkyline([][]int{[]int{3, 0, 8, 4}, []int{2, 4, 5, 7}, []int{9, 2, 6, 3}, []int{0, 3, 1, 0}})
	as.Equal(35, output)
}

func TestGetIncrease(t *testing.T) {
	as := assert.New(t)
	output := getIncrease(1, 2, 3)
	as.Equal(-2, output)
	output = getIncrease(2, 1, 3)
	as.Equal(-2, output)
	output = getIncrease(3, 2, 1)
	as.Equal(1, output)
	output = getIncrease(2, 3, 1)
	as.Equal(1, output)
}
