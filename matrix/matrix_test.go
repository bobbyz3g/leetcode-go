package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedMatrix(t *testing.T) {
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
	for _, v := range tests59 {
		assert.Equal(t, v.matrix, GenerateMatrix(v.n))
	}
}

func TestSumRegion(t *testing.T) {

	var tests304 = []struct {
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

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := NewNumMatrix(matrix)
	for _, v := range tests304 {
		assert.Equal(t, numMatrix.SumRegion(v.row1, v.col1, v.row2, v.col2), v.sum)
	}
}

func TestMatrixReshape(t *testing.T) {

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

	for _, v := range tests566 {
		assert.Equal(t, Reshape(v.matrix, v.r, v.c), v.output)
	}
}

func TestSearch(t *testing.T) {
	var tests74 = []struct {
		matrix [][]int
		target int
		res    bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			13,
			false,
		},
	}

	for _, v := range tests74 {
		assert.Equal(t, v.res, Search(v.matrix, v.target))
	}
}

func TestSpiralOrder(t *testing.T) {
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

	for _, v := range tests54 {
		assert.Equal(t, v.out, SpiralOrder(v.matrix))
	}
}

func TestKWeakRows(t *testing.T) {
	tests := []struct {
		mat [][]int
		k   int
		row []int
	}{
		{
			mat: [][]int{{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:   3,
			row: []int{2, 0, 3},
		},
		{
			mat: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{1, 1, 1},
				{1, 1, 1},
				{0, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
			k:   6,
			row: []int{4, 6, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.row, KWeakRows(tt.mat, tt.k))
	}
}
