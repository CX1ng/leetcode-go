package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindComplement(t *testing.T) {
	as := assert.New(t)
	output := findComplement(5)
	as.Equal(2, output)
	output = findComplement(1)
	as.Equal(0, output)
}
