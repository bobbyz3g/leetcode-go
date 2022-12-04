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

// DeleteDuplicates deletes all duplicates such that each element appears only once.
// Return the linked list sorted as well
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// DeleteDuplicates2 deletes all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
// Return the linked list sorted as wel
func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}

	cur := dummy
	next := cur.Next

	for next != nil && next.Next != nil {
		if next.Val == next.Next.Val {
			val := next.Val
			for next != nil && next.Val == val {
				next = next.Next
				cur.Next = next
			}
		} else {
			cur = next
			next = cur.Next
		}
	}

	return dummy.Next
}

// ReverseList reverses the given list.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ReverseBetween reverse the nodes of the list from position left to position right,
// and return the reversed list
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

// HasCycle return true  if the linked list has a cycle in it.
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && slow != nil {
		slow, fast = slow.Next, fast.Next
		if fast == nil || slow == nil {
			break
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
