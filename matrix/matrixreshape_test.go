package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests566 = []struct {
	matrix [][]int
	r      int
	c      int
	output [][]int
}{
	{
		matrix: [][]int{
			{1, 2},
			{3, 4},
		},
		r: 1,
		c: 4,
		output: [][]int{
			{1, 2, 3, 4},
		},
	},
	{
		matrix: [][]int{
			{1, 2},
			{3, 4},
		},
		r: 2,
		c: 4,
		output: [][]int{
			{1, 2},
			{3, 4},
		},
	},
}

func TestMatrixReshape(t *testing.T) {
	for _, v := range tests566 {
		assert.Equal(t, Reshape(v.matrix, v.r, v.c), v.output)
	}
}
