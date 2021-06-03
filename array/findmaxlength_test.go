package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests525 = []struct {
	nums   []int
	output int
}{
	{
		[]int{0, 1},
		2,
	},
	{
		[]int{0, 1, 0},
		2,
	},
	{
		[]int{0},
		0,
	},
}

func TestFindMaxLength(t *testing.T) {
	for _, tt := range tests525 {
		assert.Equal(t, tt.output, FindMaxLength(tt.nums))
	}
}
