package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	as := assert.New(t)
	output := hammingDistance(1, 4)
	as.Equal(2, output)
}
