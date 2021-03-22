package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func hammingWeight(num uint32) int {
	var nums uint32
	for num > 0 {
		nums += (num & 1)
		num = num >> 1
	}
	return int(nums)
}

var tests191 = []struct {
	in  uint32
	out int
}{
	{
		0b00000000000000000000000000001011,
		3,
	},
	{
		0b00000000000000000000000010000000,
		1,
	},
	{
		0b11111111111111111111111111111101,
		31,
	},
}

func TestHammingWeight(t *testing.T) {
	for _, v := range tests191 {
		assert.Equal(t, v.out, hammingWeight(v.in))
	}
}
