package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	as := assert.New(t)
	output := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	as.Equal(2, output)
}
