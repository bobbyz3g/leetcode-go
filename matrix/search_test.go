package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests74 = []struct {
	matrix [][]int
	target int
	res    bool
}{
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		3,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		13,
		false,
	},
}

func TestSearch(t *testing.T) {
	for _, v := range tests74 {
		assert.Equal(t, v.res, Search(v.matrix, v.target))
	}
}
