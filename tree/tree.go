/*
 * Developed by Kaiser925 on 2022/11/21.
 * Lasted modified 2022/11/21.
 * Copyright (c) 2022.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// FromSortedArray returns a binary search tree from given sorted array.
func FromSortedArray(nums []int) *Node {
	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &Node{Val: nums[0]}
	}

	mid := 0
	if n%2 == 0 {
		mid = n/2 - 1
	} else {
		mid = n / 2
	}

	root := &Node{Val: nums[mid]}

	if mid > 0 {
		root.Left = FromSortedArray(nums[:mid])
	}

	root.Right = FromSortedArray(nums[mid+1:])
	return root
}

func IsSameTree(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func IsSymmetric(root *Node) bool {
	return isSymmetric(root, root)
}

func isSymmetric(p, q *Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSymmetric(p.Left, q.Right) && isSymmetric(p.Right, q.Left)
}

func InorderTraversal(root *Node) []int {
	var res []int
	var inorder func(*Node)
	inorder = func(n *Node) {
		if n == nil {
			return
		}
		inorder(n.Left)
		res = append(res, n.Val)
		inorder(n.Right)
	}
	inorder(root)
	return res
}

func MaxDepth(root *Node) int {
	depth := 0
	return getMaxDepth(root, depth)
}

func getMaxDepth(root *Node, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	if root.Left == nil && root.Right == nil {
		return depth
	}
	l := getMaxDepth(root.Left, depth)
	r := getMaxDepth(root.Right, depth)
	if l > r {
		return l
	}
	return r
}
