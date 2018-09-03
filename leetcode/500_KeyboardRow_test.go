package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindWord(t *testing.T) {
	as := assert.New(t)
	output := findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	as.Equal([]string{"Alaska", "Dad"}, output)
}
