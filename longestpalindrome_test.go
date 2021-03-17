package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	begin, maxLen := 0, 1

	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			if dp[i][j] && j-i+1 > maxLen {
				begin = i
				maxLen = j - i + 1
			}
		}
	}

	return s[begin : begin+maxLen]
}

var tests5 = []struct {
	s   string
	out string
}{
	{
		s:   "babad",
		out: "bab",
	},
	{
		s:   "cbbd",
		out: "bb",
	},
	{
		s:   "ac",
		out: "a",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, v := range tests5 {
		assert.Equal(t, v.out, longestPalindrome(v.s))
	}
}
