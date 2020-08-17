package main

//110. 平衡二叉树
//https://leetcode-cn.com/problems/balanced-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getDepth(root) >= 0
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := getDepth(node.Left)
	right := getDepth(node.Right)
	if left == -1 || right == -1 || left-right > 1 || left-right < -1 {
		return -1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
func main() {

}
