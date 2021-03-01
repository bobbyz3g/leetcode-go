package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)

	for i := range nums {
		preSum[i+1] = preSum[i] + nums[i]
	}

	return NumArray{
		preSum,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.preSum[j+1] - this.preSum[i]
}

var tests_303 = []struct {
	input  []int
	output int
}{
	{
		input:  []int{0, 2},
		output: 1,
	},
	{
		input:  []int{2, 5},
		output: -1,
	},
	{
		input:  []int{0, 5},
		output: -3,
	},
}

func TestSumRange(t *testing.T) {
	array := Constructor([]int{-2, 0, 3, -5, 2, -1})
	for _, v := range tests_303 {
		assert.Equal(t, array.SumRange(v.input[0], v.input[1]), v.output)
	}
}
