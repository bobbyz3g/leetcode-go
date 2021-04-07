package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests81 = []struct {
	nums   []int
	target int
	ok     bool
}{
	{
		[]int{2, 5, 6, 0, 0, 1, 2},
		0,
		true,
	},
	{
		[]int{2, 5, 6, 0, 0, 1, 2},
		3,
		false,
	},
}

func TestExistInRotatedSorted(t *testing.T) {
	for _, v := range tests81 {
		assert.Equal(t, v.ok, ExistInRotatedSorted(v.nums, v.target))
	}
}
