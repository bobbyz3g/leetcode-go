package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests59 = []struct {
	n      int
	matrix [][]int
}{
	{
		3,
		[][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		},
	},
	{
		1,
		[][]int{
			{1},
		},
	},
	{
		0,
		[][]int{},
	},
}

func TestGeneratedMatrix(t *testing.T) {
	for _, v := range tests59 {
		assert.Equal(t, v.matrix, GenerateMatrix(v.n))
	}
}
