package nestedinteger

type NestedInteger interface {
	// IsInteger returns true if this NestedInteger holds a single integer,
	// rather than a nested list.
	IsInteger() bool

	// GetInteger return the single integer that this NestedInteger holds, if it holds a single integer
	// The result is undefined if this NestedInteger holds a nested list
	// So before calling this method, you should have a check
	GetInteger() int

	// SetInteger sets this NestedInteger to hold a single integer.
	SetInteger(value int)

	// Add sets this NestedInteger to hold a nested list and adds a nested integer to it.
	Add(elem NestedInteger)

	// GetList returns the nested list that this NestedInteger holds, if it holds a nested list
	// The list length is zero if this NestedInteger holds a single integer
	// You can access NestedInteger's List element directly if you want to modify it
	GetList() []*NestedInteger
}

type nestedInteger struct {
	isInt    bool
	val      int
	integers []*NestedInteger
}

func newNestedInteger() NestedInteger {
	return &nestedInteger{
		false,
		0,
		[]*NestedInteger{},
	}
}

func (n nestedInteger) IsInteger() bool {
	return n.isInt
}

func (n nestedInteger) GetInteger() int {
	return n.val
}

func (n *nestedInteger) SetInteger(value int) {
	n.isInt = true
	n.val = value
}

func (n *nestedInteger) Add(elem NestedInteger) {
	n.isInt = false
	n.integers = append(n.integers, &elem)
}

func (n nestedInteger) GetList() []*NestedInteger {
	return n.integers
}
