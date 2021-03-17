package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	top, bottom, left, right := 0, n-1, 0, n-1

	for v := 1; v <= n*n; {
		for i := left; i <= right; i++ {
			result[top][i] = v
			v++
		}
		top++

		for i := top; i <= bottom; i++ {
			result[i][right] = v
			v++
		}
		right--

		for i := right; i >= left; i-- {
			result[bottom][i] = v
			v++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			result[i][left] = v
			v++
		}
		left++
	}

	return result
}

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
		assert.Equal(t, v.matrix, generateMatrix(v.n))
	}
}
