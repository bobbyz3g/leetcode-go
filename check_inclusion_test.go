package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
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

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	var target, cur [26]int
	for _, v := range s1 {
		target[v-'a']++
	}

	for i := 0; i < l1; i++ {
		cur[s2[i]-'a']++
	}

	if cur == target {
		return true
	}

	for i := l1; i < l2; i++ {
		cur[s2[i-l1]-'a']--
		cur[s2[i]-'a']++
		if cur == target {
			return true
		}
	}

	return false
}

func TestCheckInclusion(t *testing.T) {
	for _, v := range tests {
		assert.Equal(t, checkInclusion(v.a, v.b), v.result)
	}
}
