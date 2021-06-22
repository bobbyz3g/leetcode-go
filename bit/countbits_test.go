package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests338 = []struct {
	num    int
	output []int
}{
	{
		num:    2,
		output: []int{0, 1, 1},
	},
	{
		num:    5,
		output: []int{0, 1, 1, 2, 1, 2},
	},
}

func TestCountBits(t *testing.T) {
	for _, v := range tests338 {
		assert.Equal(t, CountBits(v.num), v.output)
	}
}
