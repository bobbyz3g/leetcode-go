package algebra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumArray_SumRange(t *testing.T) {
	var tests303 = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{0, 2},
			output: 1,
		},
		{
			input:  []int{2, 5},
			output: -1,
		},
		{
			input:  []int{0, 5},
			output: -3,
		},
	}

	array := Constructor([]int{-2, 0, 3, -5, 2, -1})
	for _, v := range tests303 {
		assert.Equal(t, array.SumRange(v.input[0], v.input[1]), v.output)
	}
}
