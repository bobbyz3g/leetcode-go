package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 566
func matrixReshape(nums [][]int, r int, c int) ([][]int) {
	nR, nC := len(nums), len(nums[0])
	if nR * nC != r * c {
		return nums
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	for i := 0; i < nR*nC; i++ {
		result[i/c][i%c] = nums[i/nC][i%nC]
	}

	return result
}

var tests_566 = []struct{
	matrix [][]int
	r int
	c int
	output [][]int
} {
	{
		matrix: [][]int {
			[]int{1, 2},
			[]int{3, 4},
		},
		r: 1,
		c: 4,
		output: [][]int{
			[]int{1,2,3,4},
		},
	},
	{
		matrix: [][]int {
			[]int{1, 2},
			[]int{3, 4},
		},
		r: 2,
		c: 4,
		output: [][]int {
			[]int{1, 2},
			[]int{3, 4},
		},
	},
}

func TestMatrixReshape(t *testing.T) {
	for _, v := range tests_566 {
		assert.Equal(t, matrixReshape(v.matrix, v.r, v.c), v.output)
	}
}
