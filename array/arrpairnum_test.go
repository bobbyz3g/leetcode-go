package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests561 = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 4, 3, 2},
		output: 4,
	},
	{
		input:  []int{6, 2, 6, 5, 1, 2},
		output: 9,
	},
}

func TestArrayPairSum(t *testing.T) {
	for _, v := range tests561 {
		assert.Equal(t, PairSum(v.input), v.output)
	}
}
