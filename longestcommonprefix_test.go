package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	var result string
	if len(strs) == 0 {
		return result
	}

	result = strs[0]
	for i := range strs[0] {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return result
}

var tests_14 = []struct {
	strs   []string
	output string
}{
	{
		strs:   []string{"flower", "flow", "flight"},
		output: "fl",
	},
	{
		strs:   []string{"dog", "racecar", "car"},
		output: "",
	},
	{
		strs:   []string{"a"},
		output: "a",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range tests_14 {
		assert.Equal(t, longestCommonPrefix(v.strs), v.output)
	}
}
