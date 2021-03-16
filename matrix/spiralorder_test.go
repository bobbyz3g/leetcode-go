package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom, left, right := 0, m-1, 0, n-1

	for {
		// move to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}

		if top += 1; top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}

		if right -= 1; right < left {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}

		if bottom -= 1; bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}

		if left += 1; left > right {
			break
		}
	}

	return result
}

var tests_54 = []struct {
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
	for _, v := range tests_54 {
		assert.Equal(t, v.out, spiralOrder(v.matrix))
	}
}
