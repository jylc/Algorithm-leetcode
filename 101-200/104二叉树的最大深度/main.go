package main

//104. 二叉树的最大深度
//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	x, y := 1, 1
	x = maxDepth(root.Left) + 1
	y = maxDepth(root.Right) + 1
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {

}
