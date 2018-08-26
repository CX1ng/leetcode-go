package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUncommonFromSentences(t *testing.T) {
	as := assert.New(t)
	output := uncommonFromSentences("this apple is sweet", "this apple is sour")
	as.Len(output, 2)
	as.Contains(output, "sour")
	as.Contains(output, "sweet")
	output = uncommonFromSentences("apple apple", "banana")
	as.Equal([]string{"banana"}, output)
}
