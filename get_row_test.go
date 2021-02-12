package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 119 Pascal's Traingle II
var traingle_tests = []struct {
	a int
	b []int
}{
	{
		a: 1,
		b: []int{1, 1},
	},
	{
		a: 3,
		b: []int{1, 3, 3, 1},
	},
}

func getRow(rowIndex int) []int {
	l := rowIndex + 1
	result := make([]int, l)
	result[0] = 1

	for i := 0; i < rowIndex; i++ {
		pre := result[0]
		for j := 1; j < l; j++ {
			cur := result[j]
			result[j] = cur + pre
			pre = cur
		}
	}
	return result
}

func TestGetRow(t *testing.T) {
	for _, v := range traingle_tests {
		assert.Equal(t, getRow(v.a), v.b)
	}
}
