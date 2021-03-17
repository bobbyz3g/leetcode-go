package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// 561
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}
	return result
}

var tests561 = []struct {
	input  []int
	output int
}{
	{
		input:  []int{1, 4, 3, 2},
		output: 4,
	},
	{
		input:  []int{6, 2, 6, 5, 1, 2},
		output: 9,
	},
}

func TestArrayPairSum(t *testing.T) {
	for _, v := range tests561 {
		assert.Equal(t, arrayPairSum(v.input), v.output)
	}
}
