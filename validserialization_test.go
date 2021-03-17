package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func isValidSerialization(preorder string) bool {
	diff := 1
	for _, c := range strings.Split(preorder, ",") {
		diff--
		if diff < 0 {
			return false
		}
		if c != "#" {
			diff += 2
		}
	}
	return diff == 0
}

var tests331 = []struct {
	preorder string
	ok       bool
}{
	{
		"9,3,4,#,#,1,#,#,2,#,6,#,#",
		true,
	},
	{
		"1,#",
		false,
	},
	{
		"9,#,#,1",
		false,
	},
	{
		"#",
		true,
	},
}

func TestIsValidSerialization(t *testing.T) {
	for _, v := range tests331 {
		assert.Equal(t, v.ok, isValidSerialization(v.preorder))
	}
}
