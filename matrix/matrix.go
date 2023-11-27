/*
 * Developed by Kaiser925 on 2021/8/1.
 * Lasted modified 2021/8/1.
 * Copyright (c) 2021.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package matrix

import (
	"sort"

	"github.com/Kaiser925/leetcode-go/algebra"
)

// Reshape reshapes the matrix from m * n to r * c
func Reshape(nums [][]int, r int, c int) [][]int {
	nR, nC := len(nums), len(nums[0])
	if nR*nC != r*c {
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

// Search searches for a value in an m x n matrix.
func Search(matrix [][]int, target int) bool {
	// find the row of target maybe in.
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1

	if row == -1 {
		return false
	}

	// find the column.
	col := sort.SearchInts(matrix[row], target)

	return col < len(matrix[row]) && matrix[row][col] == target
}

// GenerateMatrix generates an n x n matrix filled with elements
// from 1 to n2 in spiral order.
func GenerateMatrix(n int) [][]int {
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

// SpiralOrder returns all elements of the matrix in spiral order.
func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom, left, right := 0, m-1, 0, n-1

	for {
		// move to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}

		if top++; top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}

		if right--; right < left {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}

		if bottom--; bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}

		if left++; left > right {
			break
		}
	}

	return result
}

// NumMatrix is used to calculate matrix region sum.
type NumMatrix struct {
	preSum [][]int
}

// NewNumMatrix builds NumMatrix from [][]int slice.
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

// SumRegion calculates the sum of matrix region.
func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.preSum[row2+1][col2+1] - n.preSum[row1][col2+1] - n.preSum[row2+1][col1] + n.preSum[row1][col1]
}

// KWeakRows the indices of the k weakest rows in the matrix ordered from weakest to strongest,
// mat is an m x n binary matrix.
// A row i is weaker than a row j if one of the following is true:
// - The number of 1 in row i is less than the number of 1 in row j.
// - Both rows have the same number of 1 and i < j.
func KWeakRows(mat [][]int, k int) []int {
	ret := make([]int, k)
	pairs := make([]algebra.IntPair, len(mat))
	for i, row := range mat {
		pow := sort.Search(len(row), func(j int) bool {
			return row[j] == 0
		})*100 + i
		pow = pow*100 + 1 // avoid error order of same pow of line.
		pairs[i] = algebra.IntPair{Left: pow, Right: i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].Left-pairs[i].Left > 0
	})
	for i, pair := range pairs[:k] {
		ret[i] = pair.Right
	}
	return ret
}

func SetZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	var rowZero, colZero bool

	for i := 0; i < len(matrix[0]) && !rowZero; i++ {
		rowZero = matrix[0][i] == 0
	}

	for i := 0; i < len(matrix) && !colZero; i++ {
		colZero = matrix[i][0] == 0
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowZero {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if colZero {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
