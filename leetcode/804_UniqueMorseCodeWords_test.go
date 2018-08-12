package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	as := assert.New(t)
	output := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	as.Equal(2, output)
}
