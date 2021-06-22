package algebra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests264 = []struct {
	n   int
	num int
}{
	{10, 12},
	{1, 1},
}

func TestNthUglyNumber(t *testing.T) {
	for _, v := range tests264 {
		assert.Equal(t, v.num, NthUglyNumber(v.n))
	}
}

var tests263 = []struct {
	n  int
	ok bool
}{
	{6, true},
	{8, true},
	{14, false},
	{1, true},
	{0, false},
}

func TestIsUgly(t *testing.T) {
	for _, v := range tests263 {
		assert.Equal(t, v.ok, IsUgly(v.n))
	}
}
