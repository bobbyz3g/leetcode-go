package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests300 = []struct {
	nums []int
	out  int
}{
	{
		nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
		out:  4,
	},
	{
		nums: []int{0, 1, 0, 3, 2, 3},
		out:  4,
	},
	{
		nums: []int{7, 7, 7, 7, 7, 7, 7},
		out:  1,
	},
	{
		nums: []int{},
		out:  0,
	},
}

func TestLengthOfLIS(t *testing.T) {
	for _, v := range tests300 {
		assert.Equal(t, v.out, LengthOfLIS(v.nums))
	}
}
