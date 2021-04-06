package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests80 = []struct {
	in        []int
	out       int
	newNumber []int
}{
	{
		[]int{1, 1, 1, 2, 2, 3},
		5,
		[]int{1, 1, 2, 2, 3},
	},
	{
		[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		7,
		[]int{0, 0, 1, 1, 2, 3, 3},
	},
	{
		[]int{1},
		1,
		[]int{1},
	},
	{
		[]int{},
		0,
		[]int{},
	},
	{
		[]int{1, 2, 3},
		3,
		[]int{1, 2, 3},
	},
}

func TestRemoveDuplicates2(t *testing.T) {
	for _, v := range tests80 {
		l := RemoveDuplicates2(v.in)
		assert.Equal(t, v.out, l)
		assert.Equal(t, v.newNumber, v.in[:l])
	}
}

var tests26 = []struct {
	in        []int
	out       int
	newNumber []int
}{
	{
		[]int{1, 1, 2},
		2,
		[]int{1, 2},
	},
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
		[]int{0, 1, 2, 3, 4},
	},
}

func TestRemoveDuplicates1(t *testing.T) {
	for _, v := range tests26 {
		l := RemoveDuplicates1(v.in)
		assert.Equal(t, v.out, l)
		assert.Equal(t, v.newNumber, v.in[:l])
	}
}
