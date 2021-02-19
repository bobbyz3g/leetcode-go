package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests_1004 = []struct {
	a      []int
	k      int
	output int
}{
	{
		a:      []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		k:      2,
		output: 6,
	},
	{
		a:      []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		k:      3,
		output: 10,
	},
}

func longestOnes(A []int, K int) int {
	var result, left, leftZeroNum, rightZeroNum int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right, v := range A {
		rightZeroNum = 1 - v + rightZeroNum
		for leftZeroNum < rightZeroNum-K {
			leftZeroNum = 1 - A[left] + leftZeroNum
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func TestLongestOnes(t *testing.T) {
	for _, v := range tests_1004 {
		assert.Equal(t, longestOnes(v.a, v.k), v.output)
	}
}
