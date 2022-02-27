package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	var tests567 = []struct {
		a      string
		b      string
		result bool
	}{
		{
			a:      "ab",
			b:      "eidbaooo",
			result: true,
		},
		{
			a:      "abc",
			b:      "b",
			result: false,
		},
	}
	for _, v := range tests567 {
		assert.Equal(t, CheckInclusion(v.a, v.b), v.result)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	var tests14 = []struct {
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

	for _, v := range tests14 {
		assert.Equal(t, LongestCommonPrefix(v.strs), v.output)
	}
}

func TestLongestPalindrome(t *testing.T) {

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
	for _, v := range tests5 {
		assert.Equal(t, v.out, LongestPalindrome(v.s))
	}
}

func TestPartition(t *testing.T) {

	var tests131 = []struct {
		s   string
		out [][]string
	}{
		{
			s:   "aab",
			out: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			s:   "a",
			out: [][]string{{"a"}},
		},
	}

	for _, v := range tests131 {
		assert.Equal(t, v.out, Partition(v.s))
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		s   string
		out int
	}{
		{
			s:   "abcabcbb",
			out: 3,
		},
		{
			s:   "bbbbb",
			out: 1,
		},
		{
			s:   "pwwkew",
			out: 3,
		},
		{
			s:   "a",
			out: 1,
		},
	}

	for _, v := range tests {
		assert.Equal(t, v.out, LengthOfLongestSubstring(v.s))
	}
}
