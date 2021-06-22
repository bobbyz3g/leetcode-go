package nestedint

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NestedIterator struct {
	values []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var values []int
	var f func(nl []*NestedInteger)
	f = func(nl []*NestedInteger) {
		for _, v := range nl {
			if (*v).IsInteger() {
				values = append(values, (*v).GetInteger())
			} else {
				f((*v).GetList())
			}
		}
	}
	f(nestedList)
	return &NestedIterator{values}
}

// Next returns the next value.
func (n *NestedIterator) Next() int {
	val := n.values[0]
	n.values = n.values[1:]
	return val
}

// HasNext returns true if there is a next value.
func (n *NestedIterator) HasNext() bool {
	return len(n.values) > 0
}

func TestNestedInteger(t *testing.T) {
	tests := make([]NestedInteger, 0, 3)
	elem0 := NewNestedInteger()
	elem0.Add(&nestedInteger{true, 1, nil})
	elem0.Add(&nestedInteger{true, 1, nil})
	tests = append(tests, elem0)

	elem1 := NewNestedInteger()
	elem1.SetInteger(2)
	tests = append(tests, elem1)
	tests = append(tests, elem0)

	assert.Equal(t, false, tests[0].IsInteger())
	assert.Equal(t, true, tests[1].IsInteger())
	assert.Equal(t, false, tests[2].IsInteger())
}

func TestNestedIterator(t *testing.T) {
	tests := make([]*NestedInteger, 0, 3)
	elem0 := NewNestedInteger()
	elem0.Add(&nestedInteger{true, 1, nil})
	elem0.Add(&nestedInteger{true, 1, nil})
	tests = append(tests, &elem0)

	elem1 := NewNestedInteger()
	elem1.SetInteger(2)
	tests = append(tests, &elem1)
	tests = append(tests, &elem0)

	iter := Constructor(tests)
	expected := []int{1, 1, 2, 1, 1}
	for i := 0; iter.HasNext(); i++ {
		assert.Equal(t, expected[i], iter.Next())
	}
}
