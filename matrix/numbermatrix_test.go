package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NumMatrix struct {
	preSum [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	l := len(matrix)
	if l == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	preSum := make([][]int, l+1)
	preSum[0] = make([]int, n+1)
	for i, row := range matrix {
		preSum[i+1] = make([]int, n+1)
		for j, v := range row {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + v
		}
	}

	return NumMatrix{
		preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]

}

var tests_304 = []struct {
	row1 int
	col1 int
	row2 int
	col2 int
	sum  int
}{
	{
		row1: 2,
		col1: 1,
		row2: 4,
		col2: 3,
		sum:  8,
	},
	{

		row1: 1,
		col1: 1,
		row2: 2,
		col2: 2,
		sum:  11,
	},
	{
		row1: 1,
		col1: 2,
		row2: 2,
		col2: 4,
		sum:  12,
	},
}

func TestSumRegion(t *testing.T) {

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := NewNumMatrix(matrix)
	for _, v := range tests_304 {
		assert.Equal(t, numMatrix.SumRegion(v.row1, v.col1, v.row2, v.col2), v.sum)
	}
}
