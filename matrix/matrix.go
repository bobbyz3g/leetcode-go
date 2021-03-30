package matrix

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
