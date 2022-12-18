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

// IsBalanced return true if given tree is balanced.
func IsBalanced(root *Node) bool {
	var height func(*Node) int
	height = func(root *Node) int {
		if root == nil {
			return 0
		}
		left := height(root.Left)
		right := height(root.Right)
		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	return height(root) >= 0
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

// InorderTraversal returns inorder traversal of its nodes' values.
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

// PreorderTraversal returns preorder traversal of its nodes' values.
func PreorderTraversal(root *Node) []int {
	var res []int
	var preorder func(*Node)
	preorder = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		preorder(n.Left)
		preorder(n.Right)
	}
	preorder(root)
	return res
}

// PostorderTraversal returns postorder traversal of
// its nodes' values.
func PostorderTraversal(root *Node) []int {
	var res []int
	var stack []*Node
	var prevWalked *Node
	for root != nil || len(stack) > 0 {
		for root != nil { // push all left sub-nodes of current node into stack
			stack = append(stack, root)
			root = root.Left
		}
		// pop node from stack
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == prevWalked {
			// root is a leaf node or the right sub-node has been walked
			res = append(res, root.Val)
			prevWalked = root
			root = nil
		} else {
			// there is a right sub-node which has not been walked
			stack = append(stack, root)
			root = root.Right
		}
	}
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

// MinDepth returns the min depth of tree.
// The minimum depth is the number of nodes
// along the shortest path from the root node
// down to the nearest leaf node.
func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	return 1 + min(MinDepth(root.Left), MinDepth(root.Right))
}

// HasPathSum returns true if the tree has a
// root-to-leaf path such that adding up all
// the values along the path equals targetSum.
func HasPathSum(root *Node, targetSum int) bool {
	if root == nil {
		return false
	}

	// root is leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

// Invert inverts tree, and returns its root.
func Invert(root *Node) *Node {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	Invert(root.Left)
	Invert(root.Right)
	return root
}
