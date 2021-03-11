package linkedlist

// ListNode represents singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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
