package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	as := assert.New(t)
	output := reverseString("hello")
	as.Equal("olleh", output)
	output = reverseString("A man, a plan, a canal: Panama")
	as.Equal("amanaP :lanac a ,nalp a ,nam A", output)
}
