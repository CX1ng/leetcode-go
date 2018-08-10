package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	as := assert.New(t)
	output := numJewelsInStones("aA", "aAAbbbb")
	as.Equal(3, output)
	output = numJewelsInStones("z", "ZZ")
	as.Equal(0, output)
}
