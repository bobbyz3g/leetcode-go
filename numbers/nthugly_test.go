package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests264 = []struct {
	n   int
	num int
}{
	{
		10,
		12,
	},
	{
		1,
		1,
	},
}

func TestNthUglyNumber(t *testing.T) {
	for _, v := range tests264 {
		assert.Equal(t, v.num, NthUglyNumber(v.n))
	}
}
