package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

var tests_338 = []struct {
	num    int
	output []int
}{
	{
		num:    2,
		output: []int{0, 1, 1},
	},
	{
		num:    5,
		output: []int{0, 1, 1, 2, 1, 2},
	},
}

func TestCountBits(t *testing.T) {
	for _, v := range tests_338 {
		assert.Equal(t, countBits(v.num), v.output)
	}
}
