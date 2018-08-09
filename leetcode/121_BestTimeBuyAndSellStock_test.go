package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	as := assert.New(t)
	output := maxProfit([]int{7, 1, 5, 3, 6, 4})
	as.Equal(5, output)
	output = maxProfit([]int{7, 6, 4, 3, 1})
	as.Equal(0, output)
}
