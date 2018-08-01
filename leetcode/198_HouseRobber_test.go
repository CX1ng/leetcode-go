package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	as := assert.New(t)
	// input: 1,2,3,1
	// output: 4
	input := []int{1, 2, 3, 1}
	output := rob(input)
	as.Equal(4, output)
	// input: 2,7,9,3,1
	// ouput: 12
	input = []int{2, 7, 9, 3, 1}
	output = rob(input)
	as.Equal(12, output)
}
