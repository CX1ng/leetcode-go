package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeakInMountainArray(t *testing.T) {
	as := assert.New(t)
	output := peakIndexInMountainArray([]int{0, 1, 0})
	as.Equal(1, output)
	output = peakIndexInMountainArray([]int{0, 2, 1, 0})
	as.Equal(1, output)
}
