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

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{name: "empty", str: "", want: 0},
		{name: "one space", str: " ", want: 0},
		{name: "two space", str: " ", want: 0},
		{name: "only space", str: "    ", want: 0},
		{name: "one rune", str: "a", want: 1},
		{name: "two rune", str: "ab", want: 2},
		{name: "two rune with space", str: "ab ab", want: 2},
		{name: "end with space", str: "ab ", want: 2},
		{name: "normal", str: "Hello World", want: 5},
		{name: "   fly me   to   the moon  ", str: "   fly me   to   the moon  ", want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, LengthOfLastWord(tt.str), "LengthOfLastWord(%v)", tt.str)
		})
	}
}

func TestAddBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "11, 1", args: args{a: "11", b: "1"}, want: "100"},
		{name: "1010, 1011", args: args{a: "1010", b: "1011"}, want: "10101"},
		{name: "0, 0", args: args{a: "0", b: "0"}, want: "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddBinary(tt.args.a, tt.args.b), "AddBinary(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
