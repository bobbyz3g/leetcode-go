package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests90 = []struct {
	nums []int
	res  [][]int
}{
	{
		[]int{1, 2, 2},
		[][]int{
			nil,
			{1},
			{2},
			{1, 2},
			{2, 2},
			{1, 2, 2},
		},
	},
	{
		[]int{0},
		[][]int{
			nil,
			{0},
		},
	},
}

func TestSubsetsWithDup(t *testing.T) {
	for _, v := range tests90 {
		assert.Equal(t, v.res, SubsetsWithDup(v.nums))
	}
}
