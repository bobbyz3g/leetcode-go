package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 448. Find All Numbers Disappeared in an Array
func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0, 1)
	l := len(nums)

	// local array hash.
	for _, v := range nums {
		index := (v - 1) % l
		nums[index] = nums[index] + l
	}

	for i, v := range nums {
		if v <= l {
			result = append(result, i+1)
		}
	}

	return result
}

var disappear_tests = []struct {
	a []int
	b []int
}{
	{
		a: []int{4, 3, 2, 7, 8, 2, 3, 1},
		b: []int{5, 6},
	},
}

func TestFindDisappearedNumbers(t *testing.T) {
	for _, v := range disappear_tests {
		assert.Equal(t, findDisappearedNumbers(v.a), v.b)
	}
}
