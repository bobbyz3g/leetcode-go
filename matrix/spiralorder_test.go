package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests54 = []struct {
	matrix [][]int
	out    []int
}{
	{
		[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
}

func TestSpiralOrder(t *testing.T) {
	for _, v := range tests54 {
		assert.Equal(t, v.out, SpiralOrder(v.matrix))
	}
}
