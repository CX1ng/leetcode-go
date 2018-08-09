package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	as := assert.New(t)
	min := Min(1, 2)
	as.Equal(1, min)
	min = Min(-1, -2)
	as.Equal(-2, min)
	min = Min(0, 0)
	as.Equal(0, min)
}

func TestMax(t *testing.T) {
	as := assert.New(t)
	max := Max(1, 2)
	as.Equal(2, max)
	max = Max(-1, -2)
	as.Equal(-1, max)
	max = Max(0, 0)
	as.Equal(0, max)
}
