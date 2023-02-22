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

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var la, lb int
	for n := headA; n != nil; n = n.Next {
		la++
	}
	for n := headB; n != nil; n = n.Next {
		lb++
	}

	if la > lb {
		for i := 0; i < la-lb; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < lb-la; i++ {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}
	return nil
}

// RemoveElements removes elements which value is val from head.
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// MergeInBetween removes nodes that index between a, b in list1,
// and put list2 in their place.
func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{Next: list1}
	cur := dummy
	var beforeA, afterB *ListNode
	for i := 0; cur != nil; i++ {
		if i == a {
			beforeA = cur
		}
		if i == b+1 {
			afterB = cur.Next
		}
		cur = cur.Next
	}
	beforeA.Next = list2
	bTail := list2
	for bTail.Next != nil {
		bTail = bTail.Next
	}
	bTail.Next = afterB
	return dummy.Next
}

// Partition partitions list such that all nodes less than
// x come before nodes greater than or equal to x.
func Partition(head *ListNode, x int) *ListNode {
	partOne, partTwo := &ListNode{}, &ListNode{}
	oneTail, twoTail := partOne, partTwo
	for n := head; n != nil; {
		next := n.Next
		n.Next = nil
		if n.Val >= x {
			twoTail.Next = n
			twoTail = twoTail.Next
		} else {
			oneTail.Next = n
			oneTail = oneTail.Next
		}
		n = next
	}
	oneTail.Next = partTwo.Next
	return partOne.Next
}

// SwapPairs swaps every two adjacent nodes and return its head
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	tail := dummy
	for tail.Next != nil && tail.Next.Next != nil {
		n1 := tail.Next
		n2 := tail.Next.Next
		tail.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		tail = n1
	}
	return dummy.Next
}
