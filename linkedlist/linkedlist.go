package linkedlist

// ListNode represents singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList build a new linked list from array.
func NewList(vals []int) *ListNode {
	head := &ListNode{}
	pre := head
	for _, val := range vals {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		pre.Next = node
		pre = node
	}
	return head.Next
}

// Equal returns the list l is equal to other.
func (l *ListNode) Equal(other *ListNode) bool {
	if l.Len() != other.Len() {
		return false
	}
	i, j := l, other
	for i != nil {
		if i.Val != j.Val {
			return false
		}
		i, j = i.Next, j.Next
	}
	return true
}

// Len returns the length of l.
func (l *ListNode) Len() int {
	var lLen int
	cur := l
	for cur != nil {
		lLen++
		cur = cur.Next
	}
	return lLen
}
